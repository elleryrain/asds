package user

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	usersrv "gitlab.ubrato.ru/ubrato/core/internal/service/user"
)

type Handler struct {
	logger *slog.Logger
	svc    Service
}

type Service interface {
	GetByID(ctx context.Context, tenderID int) (models.RegularUser, error)
	Get(ctx context.Context, params service.UserGetParams) (models.UserPagination, error)
	Update(ctx context.Context, params service.UserUpdateParams) error

	ReqEmailVerification(ctx context.Context, email string) error
	ConfirmEmail(ctx context.Context, params usersrv.ConfirmEmailParams) error
	ReqResetPassword(ctx context.Context, email string) error
	ConfirmResetPassword(ctx context.Context, params usersrv.ResetPasswordParams) error
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
