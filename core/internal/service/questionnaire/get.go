package questionnaire

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Get(ctx context.Context, params service.QuestionnaireGetParams) ([]models.Questionnaire, error) {
	return s.questionnaireStore.Get(ctx, s.psql.DB(), store.QuestionnaireGetParams{
		Limit:  params.Limit,
		Offset: params.Offset,
	})
}

func (s *Service) GetStatus(ctx context.Context, organizationID int) (bool, error) {
	return s.questionnaireStore.GetStatus(ctx, s.psql.DB(), organizationID)
}
