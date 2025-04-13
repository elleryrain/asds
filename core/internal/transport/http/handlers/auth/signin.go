package auth

import (
	"context"
	"fmt"
	"net/http"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service/auth"
)

func (h *Handler) V1AuthSigninPost(ctx context.Context, req *api.V1AuthSigninPostReq) (api.V1AuthSigninPostRes, error) {
	resp, err := h.authSvc.SignIn(ctx, auth.SignInParams{
		Email:    string(req.Email),
		Password: string(req.Password),
	})
	if err != nil {
		return nil, fmt.Errorf("sign in: %w", err)
	}

	cookie := http.Cookie{
		Name:     "ubrato_session",
		Value:    resp.Session.ID,
		Path:     "/",
		Expires:  resp.Session.ExpiresAt,
		HttpOnly: true,
		Secure:   true,
		SameSite: http.SameSiteNoneMode,
	}

	return &api.V1AuthSigninPostOKHeaders{
		SetCookie: api.NewOptString(cookie.String()),
		Response: api.V1AuthSigninPostOK{
			Data: api.V1AuthSigninPostOKData{
				User:        models.ConvertRegularUserModelToApi(resp.User),
				AccessToken: resp.AccessToken,
			},
		},
	}, nil
}
