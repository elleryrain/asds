package respond

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *RespondStore) Create(ctx context.Context, qe store.QueryExecutor, params store.RespondCreateParams) error {
	builder := squirrel.
		Insert("tender_responses").
		Columns(
			"tender_id",
			"organization_id",
			"price",
			"is_nds_price",
		).
		Values(
			params.TenderID,
			params.OrganizationID,
			params.Price,
			params.IsNds,
		).
		PlaceholderFormat(squirrel.Dollar)

	_, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("query row: %w", err)
	}

	return nil
}
