package user

import (
	"context"
	"errors"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	usersrv "gitlab.ubrato.ru/ubrato/core/internal/service/user"
)

func (h *Handler) V1UsersConfirmPasswordPost(ctx context.Context, req *api.V1UsersConfirmPasswordPostReq) (api.V1UsersConfirmPasswordPostRes, error) {
	if !req.GetUserID().Set || req.GetUserID().Value <= 0 {
		return nil, errors.New("invalid user id")
	}

	if err := h.svc.ConfirmResetPassword(ctx, usersrv.ResetPasswordParams{
		UserID:   req.GetUserID().Value,
		Code:     req.GetCode(),
		Password: string(req.GetPassword()),
	}); err != nil {
		return nil, fmt.Errorf("confirm reset password: %w", err)
	}

	return &api.V1UsersConfirmPasswordPostOK{}, nil

}

func (h *Handler) V1UsersRequestResetPasswordPost(ctx context.Context, req *api.V1UsersRequestResetPasswordPostReq) (api.V1UsersRequestResetPasswordPostRes, error) {
	if err := h.svc.ReqResetPassword(ctx, string(req.GetEmail())); err != nil {
		return nil, fmt.Errorf("req reset password: %w", err)
	}

	return &api.V1UsersRequestResetPasswordPostOK{}, nil
}
