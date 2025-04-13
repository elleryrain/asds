package suggest

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

type Handler struct {
	logger *slog.Logger
	svc    Service
}

type Service interface {
	City(ctx context.Context, name string) ([]models.City, error)
	Company(ctx context.Context, INN string) (string, error)
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
