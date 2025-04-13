package addition

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *AdditionStore) CreateAddition(ctx context.Context, qe store.QueryExecutor, params store.AdditionCreateParams) (int, error) {
	builder := squirrel.
		Insert("additions").
		Columns(
			"tender_id",
			"title",
			"content",
			"verification_status",
			"attachments").
		Values(
			params.TenderID,
			params.Title,
			params.Content,
			models.VerificationStatusInReview,
			pq.Array(params.Attachments)).
		Suffix(`RETURNING id`).
		PlaceholderFormat(squirrel.Dollar)

	var id int
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&id); err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return id, nil
}
