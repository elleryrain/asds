package survey

import (
	"context"
	"log/slog"

	surveyService "gitlab.ubrato.ru/ubrato/core/internal/service/survey"
)

type Handler struct {
	logger *slog.Logger
	svc    Service
}

type Service interface {
	Response(ctx context.Context, params surveyService.ResponseParams) error
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
