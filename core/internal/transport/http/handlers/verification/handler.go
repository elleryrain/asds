package verification

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

type Handler struct {
	logger              *slog.Logger
	verificationService VerificationService
}

type VerificationService interface {
	Create(ctx context.Context, params service.VerificationRequestCreateParams) error
	UpdateStatus(ctx context.Context, params service.VerificationRequestUpdateStatusParams) error
	GetByID(ctx context.Context, requestID int) (models.VerificationRequest[models.VerificationObject], error)
	Get(ctx context.Context, params service.VerificationRequestsObjectGetParams) (models.VerificationRequestPagination[models.VerificationObject], error)
}

func New(logger *slog.Logger, verificationService VerificationService) *Handler {
	return &Handler{
		logger:              logger,
		verificationService: verificationService,
	}
}
