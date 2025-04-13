package portfolio

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *PortfolioStore) Delete(ctx context.Context, qe store.QueryExecutor, id int) error {
	builder := squirrel.Delete("portfolios").
		Where(squirrel.Eq{"id": id}).
		PlaceholderFormat(squirrel.Dollar)

	portfolio, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("execute delete: %w", err)
	}

	rowsAffected, err := portfolio.RowsAffected()
	if err != nil {
		return fmt.Errorf("get rows affected: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("portfolio not found: %w", err)
	}

	return err
}
