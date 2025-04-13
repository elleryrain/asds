package catalog

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	catalogService "gitlab.ubrato.ru/ubrato/core/internal/service/catalog"
)

func (h *Handler) V1CatalogRegionsPost(ctx context.Context, req *api.V1CatalogRegionsPostReq) (api.V1CatalogRegionsPostRes, error) {
	region, err := h.svc.CreateRegion(ctx, catalogService.CreateRegionParams{
		Name: req.GetName(),
	})
	if err != nil {
		return nil, fmt.Errorf("create region: %w", err)
	}

	return &api.V1CatalogRegionsPostCreated{
		Data: models.ConvertRegionModelToApi(region),
	}, nil
}
