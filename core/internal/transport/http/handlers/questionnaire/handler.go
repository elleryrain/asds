package questionnaire

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

type Handler struct {
	logger               *slog.Logger
	questionnaireService QuestionnaireService
}

type QuestionnaireService interface {
	Create(ctx context.Context, params service.QuestionnaireCreateParams) error
	Get(ctx context.Context, params service.QuestionnaireGetParams) ([]models.Questionnaire, error)
	GetStatus(ctx context.Context, organizationID int) (bool, error)
}

func New(
	logger *slog.Logger,
	questionnaireService QuestionnaireService) *Handler {
	return &Handler{
		logger:               logger,
		questionnaireService: questionnaireService,
	}
}
