package winners

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *WinnersStore) Get(ctx context.Context, qe store.QueryExecutor, tenderID int) ([]models.Winners, error) {
	builder := squirrel.Select(
		"w.id",
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
		"w.tender_id",
		"w.accepted",
	).
		From("winners AS w").
		Where(squirrel.Eq{"w.tender_id": tenderID}).
		Join("organizations AS o ON o.id = w.organization_id").
		PlaceholderFormat(squirrel.Dollar)

	var winners []models.Winners
	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	for rows.Next() {
		var (
			winner       models.Winners
			organization models.Organization
			avatarURL    sql.NullString
		)

		err := rows.Scan(
			&winner.ID,
			&organization.ID,
			&organization.BrandName,
			&organization.FullName,
			&organization.ShortName,
			&organization.INN,
			&organization.OKPO,
			&organization.OGRN,
			&organization.KPP,
			&organization.TaxCode,
			&organization.Address,
			&avatarURL,
			&organization.Emails,
			&organization.Phones,
			&organization.Messengers,
			&organization.VerificationStatus,
			&organization.IsContractor,
			&organization.IsBanned,
			&organization.CustomerInfo,
			&organization.ContractorInfo,
			&organization.CreatedAt,
			&organization.UpdatedAt,
			&winner.TenderID,
			&winner.Accepted,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		organization.AvatarURL = avatarURL.String
		winner.Organization = organization

		winners = append(winners, winner)
	}

	return winners, nil
}

func (s *WinnersStore) GetOrganizationIDByWinnerID(ctx context.Context, qe store.QueryExecutor, winnerID int) (int, error) {
	builder := squirrel.Select("organization_id").
		From("winners").
		Where(squirrel.Eq{"id": winnerID}).
		PlaceholderFormat(squirrel.Dollar)

	var organizationID int

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&organizationID)
	if err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return organizationID, nil
}

func (s *WinnersStore) GetTenderIDByWinnerID(ctx context.Context, qe store.QueryExecutor, winnerID int) (int, error) {
	builder := squirrel.Select("tender_id").
		From("winners").
		Where(squirrel.Eq{"id": winnerID}).
		PlaceholderFormat(squirrel.Dollar)

	var tenderID int

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&tenderID)
	if err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return tenderID, nil
}
