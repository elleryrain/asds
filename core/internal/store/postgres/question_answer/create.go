package questionanswer

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (s *QuestionAnswerStore) Create(ctx context.Context, qe store.QueryExecutor, params store.CreateQuestionAnswerParams) (models.QuestionAnswer, error) {
	builder := squirrel.
		Insert("question_answer").
		Columns(
			"tender_id",
			"author_organization_id",
			"parent_id",
			"type",
			"content",
			"verification_status").
		Values(
			params.TenderID,
			params.AuthorOrganizationID,
			sql.NullInt64{Int64: int64(params.ParentID.Value), Valid: params.ParentID.Set},
			params.Type,
			params.Content,
			models.VerificationStatusInReview).
		Suffix(`
            RETURNING
                id,
                tender_id,
                author_organization_id,
                parent_id,
                type,
                content,
				verification_status`).
		PlaceholderFormat(squirrel.Dollar)

	var (
		questionAnswer models.QuestionAnswer
		parentID       sql.NullInt64
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&questionAnswer.ID,
		&questionAnswer.TenderID,
		&questionAnswer.AuthorOrganizationID,
		&parentID,
		&questionAnswer.Type,
		&questionAnswer.Content,
		&questionAnswer.VerificationStatus,
	)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			switch pqErr.Code {
			case errstore.FKViolation:
				return models.QuestionAnswer{}, errstore.ErrQuestionAnswerFKViolation
			case errstore.UniqueConstraint:
				return models.QuestionAnswer{}, errstore.ErrQuestionAnswerUniqueViolation
			}
		}
		return models.QuestionAnswer{}, fmt.Errorf("query row: %w", err)
	}

	questionAnswer.ParentID = models.Optional[int]{Value: int(parentID.Int64), Set: parentID.Valid}

	return questionAnswer, nil
}
