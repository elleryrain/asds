package portfolio

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *PortfolioStore) Count(ctx context.Context, qe store.QueryExecutor, organizationID int) (int, error) {
	var count int

	query, args, err := squirrel.
		Select("COUNT(*)").
		From("portfolios").
		Where(squirrel.Eq{"organization_id": organizationID}).
		PlaceholderFormat(squirrel.Dollar).
		ToSql()
	if err != nil {
		return 0, fmt.Errorf("failed to build query: %w", err)
	}

	err = qe.QueryRowContext(ctx, query, args...).Scan(&count)
	if err != nil {
		return 0, fmt.Errorf("failed to execute query: %w", err)
	}

	return count, nil
}
