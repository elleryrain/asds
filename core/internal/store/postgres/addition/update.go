package addition

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *AdditionStore) UpdateVerificationStatus(ctx context.Context, qe store.QueryExecutor, params store.AdditionUpdateVerifStatusParams) error {
	builder := squirrel.Update("additions").
		Set("verification_status", params.VerificationStatus).
		Set("updated_at", squirrel.Expr("CURRENT_TIMESTAMP")).
		Where(squirrel.Eq{"id": params.AdditionID}).
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
