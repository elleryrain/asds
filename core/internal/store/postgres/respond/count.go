package respond

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *RespondStore) Count(ctx context.Context, qe store.QueryExecutor, params store.RespondGetCountParams) (int, error) {
	builder := squirrel.Select("COUNT(*)").
		From("tender_responses AS tr").
		Where(squirrel.Eq{"tr.tender_id": params.TenderID}).
		PlaceholderFormat(squirrel.Dollar)

	var count int
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&count); err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return count, nil
}
