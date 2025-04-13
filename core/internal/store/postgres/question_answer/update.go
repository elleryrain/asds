package questionanswer

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *QuestionAnswerStore) UpdateVerificationStatus(ctx context.Context, qe store.QueryExecutor, params store.QuestionAnswerVerifStatusUpdateParams) (error) {
	builder := squirrel.Update("question_answer").
		Set("verification_status", params.VerificationStatus).
		Where(squirrel.Eq{"id": params.QuestionAnswerID}).
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
