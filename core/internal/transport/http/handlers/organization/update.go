package organization

import (
	"context"
	"errors"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (h *Handler) V1OrganizationsOrganizationIDProfileBrandPut(
	ctx context.Context,
	req *api.V1OrganizationsOrganizationIDProfileBrandPutReq,
	params api.V1OrganizationsOrganizationIDProfileBrandPutParams,
) (api.V1OrganizationsOrganizationIDProfileBrandPutRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to edit the organization", nil)
	}

	if err := h.organizationService.UpdateBrand(ctx, service.OrganizationUpdateBrandParams{
		OrganizationID: params.OrganizationID,
		Brand:          models.Optional[string]{Value: req.GetBrand().Value, Set: req.GetBrand().Set},
		AvatarURL:      models.Optional[string]{Value: req.AvatarURL.Value.String(), Set: req.GetAvatarURL().Set},
	}); err != nil {
		if errors.Is(err, errstore.ErrOrganizationNotFound) {
			return nil, cerr.Wrap(err, cerr.CodeNotFound, "Организация не найдена", map[string]interface{}{
				"organization_id": params.OrganizationID,
			})
		}

		return nil, fmt.Errorf("update organization brand: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDProfileBrandPutOK{}, nil
}

func (h *Handler) V1OrganizationsOrganizationIDProfileContactsPut(ctx context.Context, req *api.V1OrganizationsOrganizationIDProfileContactsPutReq, params api.V1OrganizationsOrganizationIDProfileContactsPutParams) (api.V1OrganizationsOrganizationIDProfileContactsPutRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to edit the organization", nil)
	}

	if err := h.organizationService.UpdateContacts(ctx, service.OrganizationUpdateContactsParams{
		OrganizationID: params.OrganizationID,
		Emails: models.Optional[models.ContactInfos]{
			Value: convert.Slice[[]api.ContactInfo, models.ContactInfos](req.GetEmails(), models.ConvertAPIToContactInfo),
			Set:   len(req.GetEmails()) != 0,
		},
		Phones: models.Optional[models.ContactInfos]{
			Value: convert.Slice[[]api.ContactInfo, models.ContactInfos](req.GetPhones(), models.ConvertAPIToContactInfo),
			Set:   len(req.GetPhones()) != 0,
		},
		Messengers: models.Optional[models.ContactInfos]{
			Value: convert.Slice[[]api.ContactInfo, models.ContactInfos](req.GetMessengers(), models.ConvertAPIToContactInfo),
			Set:   len(req.GetMessengers()) != 0,
		},
	}); err != nil {
		if errors.Is(err, errstore.ErrOrganizationNotFound) {
			return nil, cerr.Wrap(err, cerr.CodeNotFound, "Организация не найдена", map[string]interface{}{
				"organization_id": params.OrganizationID,
			})
		}

		return nil, fmt.Errorf("update organization contacts: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDProfileContactsPutOK{}, nil
}