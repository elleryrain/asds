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

func (h *Handler) V1TendersTenderIDRespondPost(
	ctx context.Context,
	req *api.V1TendersTenderIDRespondPostReq,
	params api.V1TendersTenderIDRespondPostParams,
) (api.V1TendersTenderIDRespondPostRes, error) {
	err := h.respondService.Create(ctx, service.RespondCreateParams{
		TenderID:       params.TenderID,
		OrganizationID: contextor.GetOrganizationID(ctx),
		Price:          req.Price,
		IsNds:          req.IsNds,
	})
	if err != nil {
		return nil, fmt.Errorf("respond: %w", err)
	}

	return &api.V1TendersTenderIDRespondPostOK{}, nil
}

func (h *Handler) V1TendersTenderIDRespondGet(
	ctx context.Context,
	params api.V1TendersTenderIDRespondGetParams,
) (api.V1TendersTenderIDRespondGetRes, error) {
	respond, err := h.respondService.Get(ctx, service.RespondGetParams{
		TenderID: params.TenderID,
		Page:     uint64(params.Page.Or(pagination.Page)),
		PerPage:  uint64(params.PerPage.Or(pagination.PerPage))})
	if err != nil {
		return nil, fmt.Errorf("get responds: %w", err)
	}

	return &api.V1TendersTenderIDRespondGetOK{
		Data:       convert.Slice[[]models.Respond, []api.Respond](respond.Responds, models.ConvertRespondModelToApi),
		Pagination: pagination.ConvertPaginationToAPI(respond.Pagination),
	}, nil
}
