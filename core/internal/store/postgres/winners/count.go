package winners

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *WinnersStore) Count(ctx context.Context, qe store.QueryExecutor, tenderID int) (int, error) {
	var count int

	query, args, err := squirrel.
		Select("COUNT(*)").
		From("winners").
		Where(squirrel.Eq{"tender_id": tenderID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("build query: %w", err)
	}

	err = qe.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("execute query: %w", err)
	}

	return count, nil
}
