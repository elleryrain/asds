package questionnaire

import (
	"context"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (s *QuestionnaireStore) Create(ctx context.Context, qe store.QueryExecutor, params store.QuestionnaireCreateParams) error {
	builder := squirrel.Insert("questionnaire").Columns(
		"organization_id",
		"answers",
		"is_completed",
	).Values(
		params.OrganizationID,
		params.Answers,
		params.IsCompleted,
	).PlaceholderFormat(squirrel.Dollar)

	if _, err := builder.RunWith(qe).ExecContext(ctx); err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) && pqErr.Code == errstore.UniqueConstraint {
			return errstore.ErrQuestionnaireExist
		}

		return fmt.Errorf("query row: %w", err)
	}

	return nil
}
