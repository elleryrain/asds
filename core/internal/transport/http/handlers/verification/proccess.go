package verification

import (
	"context"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1VerificationsRequestIDAprovePost(ctx context.Context, params api.V1VerificationsRequestIDAprovePostParams) (api.V1VerificationsRequestIDAprovePostRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	if err := h.verificationService.UpdateStatus(ctx, service.VerificationRequestUpdateStatusParams{
		UserID: contextor.GetUserID(ctx),
		RequesID: params.RequestID,
		Status: models.VerificationStatusApproved,
	}); err != nil {
		return nil, err
	}

	return &api.V1VerificationsRequestIDAprovePostOK{}, nil
}

func (h *Handler) V1VerificationsRequestIDDenyPost(ctx context.Context, req *api.V1VerificationsRequestIDDenyPostReq, params api.V1VerificationsRequestIDDenyPostParams) (api.V1VerificationsRequestIDDenyPostRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	if err := h.verificationService.UpdateStatus(ctx, service.VerificationRequestUpdateStatusParams{
		UserID: contextor.GetUserID(ctx),
		RequesID: params.RequestID,
		Status: models.VerificationStatusDeclined,
		ReviewComment: models.NewOptional(req.ReviewComment),
	}); err != nil {
		return nil, err
	}

	return &api.V1VerificationsRequestIDDenyPostOK{}, nil
}
