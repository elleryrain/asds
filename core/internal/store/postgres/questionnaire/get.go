package questionnaire

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (s *QuestionnaireStore) Get(ctx context.Context, qe store.QueryExecutor, params store.QuestionnaireGetParams) ([]models.Questionnaire, error) {
	builder := squirrel.Select(
		"q.id",
		"q.answers",
		"q.is_completed",
		"q.completed_at",
		"q.created_at",
		"o.id",
		"o.brand_name",
		"o.full_name",
		"o.short_name",
		"o.inn",
		"o.okpo",
		"o.ogrn",
		"o.kpp",
		"o.tax_code",
		"o.address",
		"o.avatar_url",
		"o.emails",
		"o.phones",
		"o.messengers",
		"o.verification_status",
		"o.is_contractor",
		"o.is_banned",
		"o.created_at",
		"o.updated_at").
		From("questionnaire AS q").
		Join("organizations AS o ON o.id = q.organization_id").
		Limit(params.Limit).
		Offset(params.Offset).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	questionnaires := []models.Questionnaire{}
	for rows.Next() {
		var (
			questionnaire models.Questionnaire
			completedAt   sql.NullTime

			avatarURL sql.NullString
		)

		if err := rows.Scan(
			&questionnaire.ID,
			&questionnaire.Answers,
			&questionnaire.IsCompleted,
			&completedAt,
			&questionnaire.CreatedAt,
			&questionnaire.Organization.ID,
			&questionnaire.Organization.BrandName,
			&questionnaire.Organization.FullName,
			&questionnaire.Organization.ShortName,
			&questionnaire.Organization.INN,
			&questionnaire.Organization.OKPO,
			&questionnaire.Organization.OGRN,
			&questionnaire.Organization.KPP,
			&questionnaire.Organization.TaxCode,
			&questionnaire.Organization.Address,
			&avatarURL,
			&questionnaire.Organization.Emails,
			&questionnaire.Organization.Phones,
			&questionnaire.Organization.Messengers,
			&questionnaire.Organization.VerificationStatus,
			&questionnaire.Organization.IsContractor,
			&questionnaire.Organization.IsBanned,
			&questionnaire.Organization.CreatedAt,
			&questionnaire.Organization.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		questionnaire.Organization.AvatarURL = avatarURL.String
		questionnaire.CompletedAt = models.Optional[time.Time]{Value: completedAt.Time, Set: completedAt.Valid}

		questionnaires = append(questionnaires, questionnaire)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows err: %w", err)
	}

	return questionnaires, nil
}

func (s *QuestionnaireStore) GetStatus(ctx context.Context, qe store.QueryExecutor, organizationID int) (bool, error) {
	builder := squirrel.Select("q.is_completed").
		From("questionnaire AS q").
		Where(squirrel.Eq{"q.organization_id": organizationID}).
		PlaceholderFormat(squirrel.Dollar)

	var isCompleted bool
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&isCompleted); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, errstore.ErrQuestionnaireNotFound
		}
		return false, fmt.Errorf("query row: %w", err)
	}

	return isCompleted, nil
}
