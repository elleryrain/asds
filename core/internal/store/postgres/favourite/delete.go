package favourite

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *FavouriteStore) Delete(ctx context.Context, qe store.QueryExecutor, favouriteID int) error {
	builder := squirrel.Delete("favourites").
		Where(squirrel.Eq{"id": favouriteID}).
		PlaceholderFormat(squirrel.Dollar)

	favourite, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("execute delete: %w", err)
	}

	rowsAffected, err := favourite.RowsAffected()
	if err != nil {
		return fmt.Errorf("get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("favourite not found: %w", err)
	}

	return err
}
