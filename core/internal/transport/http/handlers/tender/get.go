package tender

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1TendersTenderIDGet(ctx context.Context, params api.V1TendersTenderIDGetParams) (api.V1TendersTenderIDGetRes, error) {
	tender, err := h.tenderService.GetByID(ctx, params.TenderID)
	if err != nil {
		return nil, fmt.Errorf("get tender: %w", err)
	}

	return &api.V1TendersTenderIDGetOK{
		Data: models.ConvertTenderModelToApi(tender),
	}, nil
}

func (h *Handler) V1TendersGet(ctx context.Context, params api.V1TendersGetParams) (api.V1TendersGetRes, error) {
	tenders, err := h.tenderService.List(ctx, service.TenderListParams{
		VerifiedOnly: contextor.GetRole(ctx) < models.UserRoleEmployee,
		Page:         uint64(params.Page.Or(pagination.Page)),
		PerPage:      uint64(params.PerPage.Or(pagination.PerPage))})
	if err != nil {
		return nil, fmt.Errorf("get tenders: %w", err)
	}

	return &api.V1TendersGetOK{
		Data:       convert.Slice[[]models.Tender, []api.Tender](tenders.Tenders, models.ConvertTenderModelToApi),
		Pagination: pagination.ConvertPaginationToAPI(tenders.Pagination),
	}, nil
}

func (h *Handler) V1OrganizationsOrganizationIDTendersGet(
	ctx context.Context,
	params api.V1OrganizationsOrganizationIDTendersGetParams,
) (api.V1OrganizationsOrganizationIDTendersGetRes, error) {
	organizationID := contextor.GetOrganizationID(ctx)

	tenders, err := h.tenderService.List(ctx, service.TenderListParams{
		OrganizationID: models.Optional[int]{Value: params.OrganizationID, Set: true},
		VerifiedOnly:   params.OrganizationID != organizationID,
		WithDrafts:     organizationID == params.OrganizationID,
		Page:           uint64(params.Page.Or(pagination.Page)),
		PerPage:        uint64(params.PerPage.Or(pagination.PerPage)),
	})
	if err != nil {
		return nil, fmt.Errorf("get tender: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDTendersGetOK{
		Data:       convert.Slice[[]models.Tender, []api.Tender](tenders.Tenders, models.ConvertTenderModelToApi),
		Pagination: pagination.ConvertPaginationToAPI(tenders.Pagination),
	}, nil
}
