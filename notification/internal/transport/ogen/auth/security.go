package authHandler

import (
	"context"

	api "gitlab.ubrato.ru/ubrato/notification/api/gen"
	"gitlab.ubrato.ru/ubrato/notification/internal/lib/contextor"
)

func (h *Handler) HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error) {
	claims, err := h.authSvc.ValidateAccessToken(ctx, t.GetToken())
	if err != nil {
		return ctx, err
	}

	ctx = context.WithValue(ctx, contextor.UserIDKey, claims.UserID)
	ctx = context.WithValue(ctx, contextor.OrganizationIDKey, claims.OrganizationID)

	return ctx, nil
}
