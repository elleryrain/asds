package tender

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/deduplicate"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (s *TenderStore) GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Tender, error) {
	tenders, err := s.List(ctx, qe, store.TenderListParams{
		TenderIDs: models.Optional[[]int]{Set: true, Value: []int{id}},
	})
	if err != nil {
		return models.Tender{}, fmt.Errorf("list tenders: %w", err)
	}

	if len(tenders) == 0 {
		return models.Tender{}, errstore.ErrTenderNotFound
	}

	return tenders[0], nil
}

func (s *TenderStore) List(ctx context.Context, qe store.QueryExecutor, params store.TenderListParams) ([]models.Tender, error) {
	builder := squirrel.
		Select(
			"t.id",
			"t.city_id",
			"t.services_ids",
			"t.objects_ids",
			"t.name",
			"t.price",
			"t.is_contract_price",
			"t.is_nds_price",
			"t.floor_space",
			"t.description",
			"t.wishes",
			"t.specification",
			"t.attachments",
			"t.status",
			"t.verification_status",
			"t.is_draft",
			"t.reception_start",
			"t.reception_end",
			"t.work_start",
			"t.work_end",
			"t.created_at",
			"t.updated_at",
			"c.name",
			"c.id",
			"r.name",
			"r.id",
			"o.id",
			"o.brand_name",
			"o.full_name",
			"o.short_name",
			"o.inn",
			"o.okpo",
			"o.ogrn",
			"o.kpp",
			"o.tax_code",
			"o.address",
			"o.avatar_url",
			"o.emails",
			"o.phones",
			"o.messengers",
			"o.verification_status",
			"o.is_contractor",
			"o.is_banned",
			"o.customer_info",
			"o.contractor_info",
			"o.created_at",
			"o.updated_at",
		).
		From("tenders AS t").
		Join("cities AS c ON c.id = t.city_id").
		Join("regions AS r ON r.id = c.region_id").
		Join("organizations AS o ON o.id = t.organization_id").
		PlaceholderFormat(squirrel.Dollar)

	if params.OrganizationID.Set {
		builder = builder.Where(squirrel.Eq{"t.organization_id": params.OrganizationID.Value})
	}

	if params.TenderIDs.Set {
		builder = builder.Where(squirrel.Eq{"t.id": params.TenderIDs.Value})
	}

	if !params.WithDrafts {
		builder = builder.Where(squirrel.Eq{"t.is_draft": false})
	}

	if params.VerifiedOnly {
		builder = builder.Where(squirrel.Eq{"t.verification_status": models.VerificationStatusApproved})
	}

	if params.Limit.Set {
		builder = builder.Limit(params.Limit.Value)
	}

	if params.Offset.Set {
		builder = builder.Offset(params.Offset.Value)
	}

	var tenders []models.Tender
	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	var (
		tenderToServices = make(map[int][]int)
		tenderToObjects  = make(map[int][]int)
		serviceIDs       = make([]int, 0)
		objectIDs        = make([]int, 0)
	)

	for rows.Next() {
		var (
			tender           models.Tender
			tenderServiceIDs pq.Int64Array
			tenderObjectIDs  pq.Int64Array
			description      sql.NullString
			wishes           sql.NullString
			AvatarURL        sql.NullString
		)

		err = rows.Scan(
			&tender.ID,
			&tender.City.ID,
			&tenderServiceIDs,
			&tenderObjectIDs,
			&tender.Name,
			&tender.Price,
			&tender.IsContractPrice,
			&tender.IsNDSPrice,
			&tender.FloorSpace,
			&description,
			&wishes,
			&tender.Specification,
			pq.Array(&tender.Attachments),
			&tender.Status,
			&tender.VerificationStatus,
			&tender.IsDraft,
			&tender.ReceptionStart,
			&tender.ReceptionEnd,
			&tender.WorkStart,
			&tender.WorkEnd,
			&tender.CreatedAt,
			&tender.UpdatedAt,
			&tender.City.Name,
			&tender.City.ID,
			&tender.City.Region.Name,
			&tender.City.Region.ID,
			&tender.Organization.ID,
			&tender.Organization.BrandName,
			&tender.Organization.FullName,
			&tender.Organization.ShortName,
			&tender.Organization.INN,
			&tender.Organization.OKPO,
			&tender.Organization.OGRN,
			&tender.Organization.KPP,
			&tender.Organization.TaxCode,
			&tender.Organization.Address,
			&AvatarURL,
			&tender.Organization.Emails,
			&tender.Organization.Phones,
			&tender.Organization.Messengers,
			&tender.Organization.VerificationStatus,
			&tender.Organization.IsContractor,
			&tender.Organization.IsBanned,
			&tender.Organization.CustomerInfo,
			&tender.Organization.ContractorInfo,
			&tender.Organization.CreatedAt,
			&tender.Organization.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		tender.Organization.AvatarURL = AvatarURL.String
		tender.Description = description.String
		tender.Wishes = wishes.String

		tenderServiceIDsConverted := convert.Slice[[]int64, []int](tenderServiceIDs, func(i int64) int { return int(i) })
		tenderObjectIDsConverted := convert.Slice[[]int64, []int](tenderObjectIDs, func(i int64) int { return int(i) })

		tenders = append(tenders, tender)
		tenderToServices[tender.ID] = tenderServiceIDsConverted
		tenderToObjects[tender.ID] = tenderObjectIDsConverted
		serviceIDs = append(serviceIDs, tenderServiceIDsConverted...)
		objectIDs = append(objectIDs, tenderObjectIDsConverted...)
	}

	services, err := s.catalogStore.GetServicesByIDs(ctx, qe, deduplicate.Deduplicate(serviceIDs))
	if err != nil {
		return nil, fmt.Errorf("get services: %w", err)
	}

	objects, err := s.catalogStore.GetObjectsByIDs(ctx, qe, deduplicate.Deduplicate(objectIDs))
	if err != nil {
		return nil, fmt.Errorf("get objects: %w", err)
	}

	for i, tender := range tenders {
		tenderServiceIDs := tenderToServices[tender.ID]
		tenderServices := make([]models.Service, 0, len(serviceIDs))

		for _, id := range tenderServiceIDs {
			tenderServices = append(tenderServices, services[id])
		}

		tenderObjectIDs := tenderToObjects[tender.ID]
		tenderObjects := make([]models.Object, 0, len(objectIDs))

		for _, id := range tenderObjectIDs {
			tenderObjects = append(tenderObjects, objects[id])
		}

		tenders[i].Services = tenderServices
		tenders[i].Objects = tenderObjects
	}

	return tenders, nil
}

// Получает только "id", "name", "reception_start", "organization_id "
func (s *TenderStore) GetTenderNotifyInfoByObjectID(ctx context.Context, qe store.QueryExecutor, params store.TenderNotifyInfoParams) (models.Tender, error) {
	builder := squirrel.Select(
		"t.id",
		"t.name",
		"t.reception_start",
		"t.organization_id ").
		PlaceholderFormat(squirrel.Dollar)

	if params.TenderID.Set {
		builder = builder.From("tenders AS t").Where(squirrel.Eq{"t.id": params.TenderID.Value})
	}

	if params.AdditionID.Set {
		builder = builder.
			From("additions AS a").
			Join("tenders AS t ON a.tender_id = t.id").
			Where(squirrel.Eq{"a.id": params.AdditionID.Value})
	}

	if params.QuestionAnswerID.Set {
		builder = builder.
			From("question_answer AS qa").
			Join("tenders AS t ON qa.tender_id = t.id").
			Where(squirrel.Eq{"qa.id": params.QuestionAnswerID.Value})
	}

	var tender models.Tender
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&tender.ID,
		&tender.Name,
		&tender.ReceptionStart,
		&tender.Organization.ID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Tender{}, errstore.ErrTenderNotFound
		}
		return models.Tender{}, fmt.Errorf("db: %w", err)
	}

	return tender, nil
}
