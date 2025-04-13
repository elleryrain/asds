package user

import (
	"context"
	"errors"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	usersrv "gitlab.ubrato.ru/ubrato/core/internal/service/user"
)

func (h *Handler) V1UsersRequestEmailVerificationPost(ctx context.Context, req *api.V1UsersRequestEmailVerificationPostReq) (api.V1UsersRequestEmailVerificationPostRes, error) {
	err := h.svc.ReqEmailVerification(ctx, string(req.GetEmail()))
	if err != nil {
		return nil, fmt.Errorf("req email verification: %w", err)
	}

	return &api.V1UsersRequestEmailVerificationPostOK{}, nil
}

func (h *Handler) V1UsersConfirmEmailPost(ctx context.Context, req *api.V1UsersConfirmEmailPostReq) (api.V1UsersConfirmEmailPostRes, error) {
	if !req.GetUserID().Set || req.GetUserID().Value <= 0 {
		return nil, errors.New("invalid user id")
	}

	if err := h.svc.ConfirmEmail(ctx,
		usersrv.ConfirmEmailParams{
			UserID: req.GetUserID().Value,
			Code:   req.GetCode()}); err != nil {
		return nil, fmt.Errorf("confirm email: %w", err)
	}

	return &api.V1UsersConfirmEmailPostOK{}, nil
}
