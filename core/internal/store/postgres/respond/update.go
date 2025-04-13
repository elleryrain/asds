package respond

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *RespondStore) Update(ctx context.Context, qe store.QueryExecutor, params store.RespondUpdateParams) error {
	builder := squirrel.Update("tender_responses").
		Where(squirrel.Eq{"tender_id": params.TenderID}).
		Where(squirrel.Eq{"organization_id": params.OrganizationID}).
		Set("is_winner", params.IsWinner).
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
