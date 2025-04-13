package auth

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (h *Handler) V1AuthUserGet(ctx context.Context) (api.V1AuthUserGetRes, error) {
	user, err := h.userSvc.GetByID(ctx, contextor.GetUserID(ctx))
	if err != nil {
		return nil, fmt.Errorf("get user by id: %w", err)
	}

	return &api.V1AuthUserGetOK{
		Data: models.ConvertRegularUserModelToApi(user),
	}, nil
}
