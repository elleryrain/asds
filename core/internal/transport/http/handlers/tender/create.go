package tender

import (
	"context"
	"fmt"
	"net/url"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1TendersPost(ctx context.Context, req *api.V1TendersPostReq) (api.V1TendersPostRes, error) {
	tender, err := h.tenderService.Create(ctx, service.TenderCreateParams{
		Name:            req.GetName(),
		OrganizationID:  contextor.GetOrganizationID(ctx),
		CityID:          req.GetCity(),
		Price:           int(req.GetPrice() * 100),
		IsContractPrice: req.GetIsContractPrice(),
		IsNDSPrice:      req.GetIsNdsPrice(),
		IsDraft:         req.GetIsDraft().Value,
		FloorSpace:      req.GetFloorSpace(),
		Description:     req.GetDescription().Value,
		Wishes:          req.GetWishes().Value,
		Specification:   req.Specification.Value.String(),
		Attachments: convert.Slice[[]url.URL, []string](
			req.GetAttachments(), func(u url.URL) string { return u.String() },
		),
		ServiceIDs:     req.GetServices(),
		ObjectIDs:      req.GetObjects(),
		ReceptionStart: req.GetReceptionStart(),
		ReceptionEnd:   req.GetReceptionEnd(),
		WorkStart:      req.GetWorkStart(),
		WorkEnd:        req.GetWorkEnd(),
	})
	if err != nil {
		return nil, fmt.Errorf("create tender: %w", err)
	}

	return &api.V1TendersPostCreated{
		Data: models.ConvertTenderModelToApi(tender),
	}, nil
}
