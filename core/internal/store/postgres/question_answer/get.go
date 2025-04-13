package questionanswer

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (s *QuestionAnswerStore) GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.QuestionWithAnswer, error) {
	questionAnswer, err := s.Get(ctx, qe, store.QuestionAnswerGetParams{
		QuestionAnswerIDs: models.NewOptional([]int{id})})
	if err != nil {
		return models.QuestionWithAnswer{}, fmt.Errorf("get questionAnswer: %w", err)
	}

	if len(questionAnswer) == 0 {
		return models.QuestionWithAnswer{}, errstore.ErrQuestionAnswerNotFound
	}

	return questionAnswer[0], nil
}

func (s *QuestionAnswerStore) Get(ctx context.Context, qe store.QueryExecutor, params store.QuestionAnswerGetParams) ([]models.QuestionWithAnswer, error) {
	builder := squirrel.
		Select(
			"q.id",
			"q.tender_id",
			"q.type",
			"q.content",
			"q.author_organization_id",
			"q.verification_status",

			"a.id ",
			"a.tender_id",
			"a.type",
			"a.content",
			"a.parent_id",
			"a.author_organization_id",
			"a.verification_status",
		).
		From("question_answer q").
		LeftJoin("question_answer a ON a.parent_id = q.id").
		PlaceholderFormat(squirrel.Dollar)

	if params.QuestionAnswerIDs.Set {
		builder = builder.Where(squirrel.Expr("q.id IN (?)", squirrel.
			Select("DISTINCT COALESCE(parent_id, id)").
			From("question_answer").
			Where(squirrel.Eq{"id": params.QuestionAnswerIDs.Value})))
	}

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	var questionWithAnswers []models.QuestionWithAnswer

	for rows.Next() {
		var (
			qa                         models.QuestionWithAnswer
			answerID                   sql.NullInt64
			answerTenderID             sql.NullInt64
			answerType                 sql.NullInt16
			answerContent              sql.NullString
			answerParentID             sql.NullInt64
			answerAuthorOrganizationID sql.NullInt64
			answerVerificationStatus   sql.NullInt16
		)

		err = rows.Scan(
			&qa.Question.ID,
			&qa.Question.TenderID,
			&qa.Question.Type,
			&qa.Question.Content,
			&qa.Question.AuthorOrganizationID,
			&qa.Question.VerificationStatus,
			&answerID,
			&answerTenderID,
			&answerType,
			&answerContent,
			&answerParentID,
			&answerAuthorOrganizationID,
			&answerVerificationStatus,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		if answerParentID.Valid {
			qa.Answer = models.NewOptional(models.QuestionAnswer{
				ID:                   int(answerID.Int64),
				TenderID:             int(answerTenderID.Int64),
				AuthorOrganizationID: int(answerAuthorOrganizationID.Int64),
				ParentID:             models.NewOptional(int(answerParentID.Int64)),
				Type:                 models.QuestionAnswerType(answerType.Int16),
				Content:              answerContent.String,
				VerificationStatus:   models.VerificationStatus(answerVerificationStatus.Int16),
			})
		}

		questionWithAnswers = append(questionWithAnswers, qa)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("row iteration: %w", rows.Err())
	}

	return questionWithAnswers, nil
}

// GetWithAccess получает вопросы-ответы с фильтрацией:
//
// 1. Создатель тендера: вопросы со статусом approved, все ответы.
//
// 2. Авторизированный пользователь: все свои вопросы; остальные вопросы и ответы со статусом approved.
//
// 3. Неавторизированный пользователь: вопросы и ответы статусом approved.
func (s *QuestionAnswerStore) GetWithAccess(ctx context.Context, qe store.QueryExecutor, params store.QuestionAnswerGetWithAccessParams) ([]models.QuestionWithAnswer, error) {
	builder := squirrel.
		Select(
			"q.id",
			"q.tender_id",
			"q.type",
			"q.content",
			"q.author_organization_id",
			"q.verification_status",

			"a.id ",
			"a.tender_id",
			"a.type",
			"a.content",
			"a.parent_id",
			"a.author_organization_id",
			"a.verification_status",
		).
		From("question_answer q").
		Where(squirrel.Eq{
			"q.tender_id": params.TenderID,
			"q.type":      models.QuestionAnswerTypeQuestion}).
		OrderBy("q.created_at DESC").
		PlaceholderFormat(squirrel.Dollar)

	if params.IsTenderCreator {
		builder = builder.LeftJoin("question_answer a ON a.parent_id = q.id")
	} else {
		builder = builder.LeftJoin("question_answer a ON a.parent_id = q.id AND a.verification_status = ?", models.VerificationStatusApproved)
	}

	switch {
	case params.OrganizationID.Set && !params.IsTenderCreator:
		// Авторизированный пользователь
		builder = builder.Where(squirrel.Or{
			squirrel.Eq{"q.verification_status": models.VerificationStatusApproved},
			squirrel.Eq{"q.author_organization_id": params.OrganizationID.Value},
		})
	default:
		builder = builder.Where(squirrel.Eq{"q.verification_status": models.VerificationStatusApproved})
	}

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	var questionWithAnswers []models.QuestionWithAnswer

	for rows.Next() {
		var (
			qa                         models.QuestionWithAnswer
			answerID                   sql.NullInt64
			answerTenderID             sql.NullInt64
			answerType                 sql.NullInt16
			answerContent              sql.NullString
			answerParentID             sql.NullInt64
			answerAuthorOrganizationID sql.NullInt64
			answerVerificationStatus   sql.NullInt16
		)

		err = rows.Scan(
			&qa.Question.ID,
			&qa.Question.TenderID,
			&qa.Question.Type,
			&qa.Question.Content,
			&qa.Question.AuthorOrganizationID,
			&qa.Question.VerificationStatus,
			&answerID,
			&answerTenderID,
			&answerType,
			&answerContent,
			&answerParentID,
			&answerAuthorOrganizationID,
			&answerVerificationStatus,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		if answerParentID.Valid {
			qa.Answer = models.NewOptional(models.QuestionAnswer{
				ID:                   int(answerID.Int64),
				TenderID:             int(answerTenderID.Int64),
				AuthorOrganizationID: int(answerAuthorOrganizationID.Int64),
				ParentID:             models.NewOptional(int(answerParentID.Int64)),
				Type:                 models.QuestionAnswerType(answerType.Int16),
				Content:              answerContent.String,
				VerificationStatus:   models.VerificationStatus(answerVerificationStatus.Int16),
			})
		}

		questionWithAnswers = append(questionWithAnswers, qa)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("row iteration: %w", rows.Err())
	}

	return questionWithAnswers, nil
}

func (s *QuestionAnswerStore) GetAuthorOrganizationIDByID(ctx context.Context, qe store.QueryExecutor, qeID int) (int, error) {
	builder := squirrel.
		Select("author_organization_id").
		From("question_answer").
		Where(squirrel.Eq{"id": qeID}).
		PlaceholderFormat(squirrel.Dollar)

	var authorOrganizationID int
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&authorOrganizationID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errstore.ErrQuestionAnswerNotFound
		}
		return 0, fmt.Errorf("query row: %w", err)
	}

	return authorOrganizationID, nil
}
