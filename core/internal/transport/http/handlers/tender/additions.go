package tender

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1TendersTenderIDAdditionsPost(ctx context.Context, req *api.V1TendersTenderIDAdditionsPostReq, params api.V1TendersTenderIDAdditionsPostParams) (api.V1TendersTenderIDAdditionsPostRes, error) {
	if err := h.tenderService.CreateAddition(ctx, service.AdditionCreateParams{
		TenderID:    params.TenderID,
		Title:       req.Title,
		Content:     req.Content,
		Attachments: req.Attachments,
	}); err != nil {
		return nil, fmt.Errorf("create addition: %w", err)
	}

	return &api.V1TendersTenderIDAdditionsPostOK{}, nil
}

func (h *Handler) V1TendersTenderIDAdditionsGet(ctx context.Context, params api.V1TendersTenderIDAdditionsGetParams) (api.V1TendersTenderIDAdditionsGetRes, error) {
	Addition, err := h.tenderService.GetAdditions(ctx, service.GetAdditionParams{TenderID: params.TenderID})
	if err != nil {
		return nil, fmt.Errorf("get tender: %w", err)
	}

	return &api.V1TendersTenderIDAdditionsGetOK{
		Data: convert.Slice[[]models.Addition, []api.Addition](Addition, models.ConvertAdditionModelToApi),
	}, nil
}
