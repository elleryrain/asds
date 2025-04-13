package winners

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *WinnersStore) Create(ctx context.Context, qe store.QueryExecutor, params store.WinnersCreateParams) (models.Winners, error) {
	builder := squirrel.
		Insert("winners").
		Columns(
			"organization_id",
			"tender_id",
		).
		Values(
			params.OrganizationID,
			params.TenderID,
		).
		Suffix(`
			RETURNING
				id,
				tender_id,
				accepted
		`).
		PlaceholderFormat(squirrel.Dollar)

	var winner models.Winners

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&winner.ID,
		&winner.TenderID,
		&winner.Accepted,
	)
	if err != nil {
		return models.Winners{}, fmt.Errorf("query row: %w", err)
	}

	orgBuilder := squirrel.
		Select(
			"id",
			"brand_name",
			"full_name",
			"short_name",
			"inn",
			"okpo",
			"ogrn",
			"kpp",
			"tax_code",
			"address",
			"avatar_url",
			"emails",
			"phones",
			"messengers",
			"verification_status",
			"is_contractor",
			"is_banned",
			"customer_info",
			"contractor_info",
			"created_at",
			"updated_at",
		).
		From("organizations").
		Where(squirrel.Eq{"id": params.OrganizationID}).
		PlaceholderFormat(squirrel.Dollar)

	var (
		organization models.Organization
		avatarURL    sql.NullString
	)

	err = orgBuilder.RunWith(qe).QueryRowContext(ctx).Scan(
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
	)
	if err != nil {
		if err == sql.ErrNoRows {
			return models.Winners{}, fmt.Errorf("organization not found for ID: %d", params.OrganizationID)
		}
		return models.Winners{}, fmt.Errorf("query organization: %w", err)
	}

	organization.AvatarURL = avatarURL.String
	winner.Organization = organization

	return winner, nil
}
