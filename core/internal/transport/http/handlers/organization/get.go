package organization

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1OrganizationsGet(ctx context.Context, params api.V1OrganizationsGetParams) (api.V1OrganizationsGetRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	organizations, err := h.organizationService.Get(ctx, service.OrganizationGetParams{
		Page:    uint64(params.Page.Or(pagination.Page)),
		PerPage: uint64(params.PerPage.Or(pagination.PerPage))})
	if err != nil {
		return nil, fmt.Errorf("get organizations: %w", err)
	}

	return &api.V1OrganizationsGetOK{
		Data:       convert.Slice[[]models.Organization, []api.Organization](organizations.Organizations, models.ConvertOrganizationModelToApi),
		Pagination: pagination.ConvertPaginationToAPI(organizations.Pagination),
	}, nil
}

func (h *Handler) V1OrganizationsOrganizationIDGet(ctx context.Context, params api.V1OrganizationsOrganizationIDGetParams) (api.V1OrganizationsOrganizationIDGetRes, error) {
	// TODO: check role
	organization, err := h.organizationService.GetByID(ctx, params.OrganizationID)
	if err != nil {
		return nil, fmt.Errorf("get organization by id: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDGetOK{
		Data: models.ConvertOrganizationModelToApi(organization),
	}, nil
}

func (h *Handler) V1OrganizationsContractorsGet(ctx context.Context, params api.V1OrganizationsContractorsGetParams) (api.V1OrganizationsContractorsGetRes, error) {
	organizations, err := h.organizationService.GetContractors(ctx, service.OrganizationContractorsGetParams{
		Page:    uint64(params.Page.Or(pagination.Page)),
		PerPage: uint64(params.PerPage.Or(pagination.PerPage))})
	if err != nil {
		return nil, fmt.Errorf("get organizations: %w", err)
	}

	return &api.V1OrganizationsContractorsGetOK{
		Data: convert.Slice[[]models.Organization, []api.V1OrganizationsContractorsGetOKDataItem](
			organizations.Organizations, models.ConvertContractorModelToApi,
		),
		Pagination: pagination.ConvertPaginationToAPI(organizations.Pagination),
	}, nil
}
