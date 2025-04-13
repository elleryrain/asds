package organization

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *OrganizationStore) AddUser(ctx context.Context, qe store.QueryExecutor, params store.OrganizationAddUserParams) error {
	builder := squirrel.
		Insert("organization_users").
		Columns(
			"organization_id",
			"user_id",
			"is_owner",
		).
		Values(
			params.OrganizationID,
			params.UserID,
			params.IsOwner,
		).
		PlaceholderFormat(squirrel.Dollar)

	_, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("exec: %w", err)
	}

	return nil
}
