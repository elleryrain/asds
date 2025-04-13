package tender

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *TenderStore) Create(ctx context.Context, qe store.QueryExecutor, params store.TenderCreateParams) (int, error) {
	var status models.TenderStatus
	if params.IsDraft {
		status = models.DraftStatus
	} else {
		status = models.OnModerationStatus
	}

	builder := squirrel.
		Insert("tenders").
		Columns(
			"organization_id",
			"city_id",
			"services_ids",
			"objects_ids",
			"name",
			"price",
			"is_contract_price",
			"is_nds_price",
			"floor_space",
			"description",
			"wishes",
			"specification",
			"attachments",
			"status",
			"verification_status",
			"is_draft",
			"reception_start",
			"reception_end",
			"work_start",
			"work_end",
		).
		Values(
			params.OrganizationID,
			params.CityID,
			pq.Array(params.ServiceIDs),
			pq.Array(params.ObjectIDs),
			params.Name,
			params.Price,
			params.IsContractPrice,
			params.IsNDSPrice,
			params.FloorSpace,
			params.Description,
			params.Wishes,
			params.Specification,
			pq.Array(params.Attachments),
			status,
			models.VerificationStatusInReview,
			params.IsDraft,
			params.ReceptionStart,
			params.ReceptionEnd,
			params.WorkStart,
			params.WorkEnd,
		).
		Suffix(`
			RETURNING id
		`).
		PlaceholderFormat(squirrel.Dollar)

	var id int

	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&id); err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return id, nil
}
