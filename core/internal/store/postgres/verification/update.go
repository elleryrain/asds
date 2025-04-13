package verification

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *VerificationStore) UpdateStatus(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestUpdateStatusParams) (store.VerificationObjectUpdateStatusResult, error) {
	builder := squirrel.Update("verification_requests").
		Set("reviewer_user_id", params.UserID).
		Set("status", params.Status).
		Set("reviewed_at", squirrel.Expr("CURRENT_TIMESTAMP")).
		Suffix(`
			RETURNING 
				object_id,
				object_type
		`).
		Where(squirrel.Eq{"id": params.RequestID}).
		PlaceholderFormat(squirrel.Dollar)

	if params.ReviewComment.Set {
		builder = builder.Set("review_comment", params.ReviewComment.Value)
	}

	result := store.VerificationObjectUpdateStatusResult{}
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&result.ObjectID,
		&result.ObjectType,
	); err != nil {
		return result, fmt.Errorf("query row: %w", err)
	}

	return result, nil
}
