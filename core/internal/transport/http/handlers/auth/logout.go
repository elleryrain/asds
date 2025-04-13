package auth

import (
	"context"
	"fmt"
	"net/http"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

func (h *Handler) V1AuthLogoutPost(ctx context.Context, params api.V1AuthLogoutPostParams) (api.V1AuthLogoutPostRes, error) {
	if err := h.authSvc.Logout(ctx, params.UbratoSession); err != nil {
		return nil, fmt.Errorf("logout: %w", err)
	}

	cookie := &http.Cookie{
		Name:     "ubrato_session",
		Value:    "",
		Path:     "/",
		MaxAge:   -1,
		HttpOnly: true,
	}

	return &api.V1AuthLogoutPostNoContent{
		SetCookie: api.NewOptString(cookie.String()),
	}, nil
}
