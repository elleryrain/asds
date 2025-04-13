package auth

import (
	"context"
	"fmt"
	"net/http"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service/auth"
)

func (h *Handler) V1AuthSignupPost(ctx context.Context, req *api.V1AuthSignupPostReq) (api.V1AuthSignupPostRes, error) {
	resp, err := h.authSvc.SignUp(ctx, auth.SignUpParams{
		Email:        string(req.GetEmail()),
		Phone:        string(req.GetPhone()),
		Password:     string(req.GetPassword()),
		FirstName:    string(req.GetFirstName()),
		LastName:     string(req.GetLastName()),
		MiddleName:   models.Optional[string]{Value: string(req.MiddleName.Value), Set: req.MiddleName.Set},
		AvatarURL:    req.AvatarURL.Value.String(),
		INN:          string(req.GetInn()),
		IsContractor: req.GetIsContractor(),
	})
	if err != nil {
		return nil, fmt.Errorf("signup: %w", err)
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

	return &api.V1AuthSignupPostCreatedHeaders{
		SetCookie: api.NewOptString(cookie.String()),
		Response: api.V1AuthSignupPostCreated{
			Data: api.V1AuthSignupPostCreatedData{
				User:        models.ConvertRegularUserModelToApi(resp.User),
				AccessToken: resp.AccessToken,
			},
		},
	}, nil
}
