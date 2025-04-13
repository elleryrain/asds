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

func (h *Handler) V1OrganizationsOrganizationIDProfileContractorGet(ctx context.Context, params api.V1OrganizationsOrganizationIDProfileContractorGetParams) (api.V1OrganizationsOrganizationIDProfileContractorGetRes, error) {
	organization, err := h.organizationService.GetContractorByID(ctx, params.OrganizationID)
	switch {
	case errors.Is(err, errstore.ErrOrganizationNotFound):
		return nil, cerr.Wrap(err, cerr.CodeNotFound, "Организация не найдена", map[string]interface{}{
			"organization_id": params.OrganizationID,
		})
	case errors.Is(err, errstore.ErrOrganizationNotAContractor):
		return nil, cerr.Wrap(err, cerr.CodeUnprocessableEntity, "Организация не является исполнителем", map[string]interface{}{
			"organization_id": params.OrganizationID,
		})
	case err != nil:
		return nil, fmt.Errorf("get organizations customer profile")
	}

	return &api.V1OrganizationsOrganizationIDProfileContractorGetOK{
		Data: api.V1OrganizationsOrganizationIDProfileContractorGetOKData{
			Organization: models.ConvertOrganizationModelToApi(organization),
			Profile:      models.ConvertContractorInfoToApi(organization.ContractorInfo),
		},
	}, nil
}

func (h *Handler) V1OrganizationsOrganizationIDProfileCustomerGet(ctx context.Context, params api.V1OrganizationsOrganizationIDProfileCustomerGetParams) (api.V1OrganizationsOrganizationIDProfileCustomerGetRes, error) {
	organization, err := h.organizationService.GetCustomer(ctx, params.OrganizationID)
	if err != nil {
		if errors.Is(err, errstore.ErrOrganizationNotFound) {
			return nil, cerr.Wrap(err, cerr.CodeNotFound, "Организация не найдена", map[string]interface{}{
				"organization_id": params.OrganizationID,
			})
		}
		return nil, fmt.Errorf("get organizations customer profile: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDProfileCustomerGetOK{
		Data: api.V1OrganizationsOrganizationIDProfileCustomerGetOKData{
			Organization: models.ConvertOrganizationModelToApi(organization),
			Profile:      models.ConvertCustomerInfoToApi(organization.CustomerInfo),
		},
	}, nil
}

func (h *Handler) V1OrganizationsOrganizationIDProfileCustomerPut(ctx context.Context, req *api.V1OrganizationsOrganizationIDProfileCustomerPutReq, params api.V1OrganizationsOrganizationIDProfileCustomerPutParams) (api.V1OrganizationsOrganizationIDProfileCustomerPutRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to edit the organization", nil)
	}

	organization, err := h.organizationService.UpdateCustomer(ctx, service.OrganizationUpdateCustomerParams{
		OrganizationID: params.OrganizationID,
		Description:    models.Optional[string]{Value: string(req.GetDescription().Value), Set: req.GetDescription().Set},
		CityIDs:        req.GetCityIds()})
	if err != nil {
		if errors.Is(err, errstore.ErrOrganizationNotFound) {
			return nil, cerr.Wrap(err, cerr.CodeNotFound, "Организация не найдена", map[string]interface{}{
				"organization_id": params.OrganizationID,
			})
		}

		return nil, fmt.Errorf("update organization brand: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDProfileCustomerPutOK{
		Data: api.V1OrganizationsOrganizationIDProfileCustomerPutOKData{
			Organization: models.ConvertOrganizationModelToApi(organization),
			Profile:      models.ConvertCustomerInfoToApi(organization.CustomerInfo),
		},
	}, nil
}

func (h *Handler) V1OrganizationsOrganizationIDProfileContractorPut(ctx context.Context, req *api.V1OrganizationsOrganizationIDProfileContractorPutReq, params api.V1OrganizationsOrganizationIDProfileContractorPutParams) (api.V1OrganizationsOrganizationIDProfileContractorPutRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to edit the organization", nil)
	}

	organization, err := h.organizationService.UpdateContractor(ctx, service.OrganizationUpdateContractorParams{
		OrganizationID: params.OrganizationID,
		Description:    models.Optional[string]{Value: string(req.GetDescription().Value), Set: req.GetDescription().Set},
		CityIDs:        req.GetCityIds(),
		Services: convert.Slice[[]api.V1OrganizationsOrganizationIDProfileContractorPutReqServicesItem, []models.ServiceWithPrice](
			req.GetServices(), func(i api.V1OrganizationsOrganizationIDProfileContractorPutReqServicesItem) models.ServiceWithPrice {
				return models.ServiceWithPrice{
					ServiceID: i.ServiceID,
					MeasureID: i.MeasureID,
					Price:     i.Price,
				}
			}),
		ObjectIDs: req.GetObjectsIds()})

	switch {
	case errors.Is(err, errstore.ErrOrganizationNotFound):
		return nil, cerr.Wrap(err, cerr.CodeNotFound, "Организация не найдена", map[string]interface{}{
			"organization_id": params.OrganizationID,
		})
	case errors.Is(err, errstore.ErrOrganizationNotAContractor):
		return nil, cerr.Wrap(err, cerr.CodeUnprocessableEntity, "Организация не является исполнителем", map[string]interface{}{
			"organization_id": params.OrganizationID,
		})
	case err != nil:
		return nil, fmt.Errorf("update organization brand: %w", err)

	}

	return &api.V1OrganizationsOrganizationIDProfileContractorPutOK{
		Data: api.V1OrganizationsOrganizationIDProfileContractorPutOKData{
			Organization: models.ConvertOrganizationModelToApi(organization),
			Profile:      models.ConvertContractorInfoToApi(organization.ContractorInfo),
		},
	}, nil
}
