package user

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (s *UserStore) Update(ctx context.Context, qe store.QueryExecutor, params store.UserUpdateParams) error {
	builder := squirrel.Update("users ").
		Set("updated_at", squirrel.Expr("CURRENT_TIMESTAMP")).
		Where(squirrel.Eq{"id": params.UserID}).
		PlaceholderFormat(squirrel.Dollar)

	if params.Phone.Set {
		builder = builder.Set("phone", params.Phone.Value)
	}

	if params.FirstName.Set {
		builder = builder.Set("first_name", params.FirstName.Value)
	}

	if params.LastName.Set {
		builder = builder.Set("last_name", params.LastName.Value)
	}

	if params.MiddleName.Set {
		builder = builder.Set("middle_name", params.MiddleName.Value)
	}

	if params.AvatarURL.Set {
		builder = builder.Set("avatar_url", params.AvatarURL.Value)
	}

	result, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("exec row: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return errstore.ErrUserNotFound
	}

	return nil
}
