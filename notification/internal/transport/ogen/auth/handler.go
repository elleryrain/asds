package authHandler

import (
	"context"

	"gitlab.ubrato.ru/ubrato/notification/internal/lib/token"
)

type Handler struct {
	authSvc AuthService
}

type AuthService interface {
	ValidateAccessToken(ctx context.Context, accessToken string) (token.Claims, error)
}

func New(authSvc AuthService) *Handler {
	return &Handler{
		authSvc: authSvc,
	}
}
