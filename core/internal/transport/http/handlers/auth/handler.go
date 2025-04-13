package auth

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/token"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	authService "gitlab.ubrato.ru/ubrato/core/internal/service/auth"
)

type Handler struct {
	logger  *slog.Logger
	authSvc AuthService
	userSvc UserService
}

type AuthService interface {
	SignUp(ctx context.Context, params authService.SignUpParams) (authService.SignUpResult, error)
	SignIn(ctx context.Context, params authService.SignInParams) (authService.SignInResult, error)
	Refresh(ctx context.Context, sessionToken string) (authService.SignInResult, error)
	ValidateAccessToken(ctx context.Context, accessToken string) (token.Claims, error)
	Logout(ctx context.Context, sessionToken string) error
}

type UserService interface {
	GetByID(ctx context.Context, userID int) (models.RegularUser, error)
}

func New(
	logger *slog.Logger,
	authSvc AuthService,
	userSvc UserService,
) *Handler {
	return &Handler{
		logger:  logger,
		authSvc: authSvc,
		userSvc: userSvc,
	}
}
