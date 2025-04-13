package organization

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (s *OrganizationStore) UpdateVerificationStatus(ctx context.Context, qe store.QueryExecutor, params store.OrganizationUpdateVerifStatusParams) error {
	builder := squirrel.Update("organizations").
		Set("verification_status", params.VerificationStatus).
		Set("updated_at", squirrel.Expr("CURRENT_TIMESTAMP")).
		Where(squirrel.Eq{"id": params.OrganizationID}).
		PlaceholderFormat(squirrel.Dollar)

	result, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("exec row: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")
	}

	return nil
}

func (s *OrganizationStore) Update(ctx context.Context, qe store.QueryExecutor, params store.OrganizationUpdateParams) error {
	builder := squirrel.Update("organizations").
		Set("updated_at", squirrel.Expr("CURRENT_TIMESTAMP")).
		Where(squirrel.Eq{"id": params.OrganizationID}).
		PlaceholderFormat(squirrel.Dollar)

	if params.Brand.Set {
		builder = builder.Set("brand_name", params.Brand.Value)
	}

	if params.AvatarURL.Set {
		builder = builder.Set("avatar_url", params.AvatarURL.Value)
	}

	if params.Emails.Set {
		builder = builder.Set("emails", params.Emails.Value)
	}

	if params.Phones.Set {
		builder = builder.Set("phones", params.Phones.Value)
	}

	if params.Messengers.Set {
		builder = builder.Set("messengers", params.Messengers.Value)
	}

	if params.CustomerInfo.Set {
		builder = builder.Set("customer_info", params.CustomerInfo.Value)
	}

	if params.ContractorInfo.Set {
		builder = builder.Set("contractor_info", params.ContractorInfo.Value)
	}

	result, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("exec row: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return errstore.ErrOrganizationNotFound
	}

	return nil
}
