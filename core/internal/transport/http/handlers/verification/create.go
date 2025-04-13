package verification

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1VerificationsOrganizationsOrganizationIDPost(ctx context.Context, req []api.Attachment, params api.V1VerificationsOrganizationsOrganizationIDPostParams) (api.V1VerificationsOrganizationsOrganizationIDPostRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to request verification of the organization", nil)
	}

	if err := h.verificationService.Create(ctx, service.VerificationRequestCreateParams{
		ObjectID:    params.OrganizationID,
		ObjectType:  models.ObjectTypeOrganization,
		Attachments: convert.Slice[[]api.Attachment, []models.Attachment](req, models.ConvertAPIToAttachment),
	}); err != nil {
		return nil, fmt.Errorf("create verif req: %w", err)
	}

	return &api.V1VerificationsOrganizationsOrganizationIDPostOK{}, nil
}
