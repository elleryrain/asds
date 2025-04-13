package organization

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *OrganizationStore) Create(ctx context.Context, qe store.QueryExecutor, params store.OrganizationCreateParams) (models.Organization, error) {
	builder := squirrel.
		Insert("organizations").
		Columns(
			"brand_name",
			"full_name",
			"short_name",
			"inn",
			"okpo",
			"ogrn",
			"kpp",
			"tax_code",
			"address",
			"emails",
			"phones",
			"messengers",
			"is_contractor",
			"customer_info",
			"contractor_info",
		).
		Values(
			params.BrandName,
			params.FullName,
			params.ShortName,
			params.INN,
			params.OKPO,
			params.OGRN,
			params.KPP,
			params.TaxCode,
			params.Address,
			models.ContactInfos{},
			models.ContactInfos{},
			models.ContactInfos{},
			params.IsContractor,
			models.CustomerInfo{},
			models.CustomerInfo{},
		).
		Suffix(`
			RETURNING
				id,
				brand_name,
				full_name,
				short_name,
				is_contractor,
				inn,
				okpo,
				ogrn,
				kpp,
				tax_code,
				address,
				avatar_url,
				emails,
				phones,
				messengers,
				created_at,
				updated_at
	`).
		PlaceholderFormat(squirrel.Dollar)

	var (
		createdOrganization models.Organization
		avatarURL           sql.NullString
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdOrganization.ID,
		&createdOrganization.BrandName,
		&createdOrganization.FullName,
		&createdOrganization.ShortName,
		&createdOrganization.IsContractor,
		&createdOrganization.INN,
		&createdOrganization.OKPO,
		&createdOrganization.OGRN,
		&createdOrganization.KPP,
		&createdOrganization.TaxCode,
		&createdOrganization.Address,
		&avatarURL,
		&createdOrganization.Emails,
		&createdOrganization.Phones,
		&createdOrganization.Messengers,
		&createdOrganization.CreatedAt,
		&createdOrganization.UpdatedAt,
	)
	if err != nil {
		return models.Organization{}, fmt.Errorf("query row: %w", err)
	}

	createdOrganization.AvatarURL = avatarURL.String

	return createdOrganization, nil
}
