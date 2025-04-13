package verification

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *VerificationStore) Create(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestCreateParams) error {
	builder := squirrel.
		Insert("verification_requests").
		Columns(
			"object_type",
			"object_id",
			"attachments",
			"status",
		).
		Values(
			params.ObjectType,
			params.ObjectID,
			params.Attachments,
			models.VerificationStatusInReview,
		).
		PlaceholderFormat(squirrel.Dollar)

	_, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("query row: %w", err)
	}

	return nil
}
