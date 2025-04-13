package catalog

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	catalogService "gitlab.ubrato.ru/ubrato/core/internal/service/catalog"
)

func (h *Handler) V1CatalogCitiesPost(ctx context.Context, req *api.V1CatalogCitiesPostReq) (api.V1CatalogCitiesPostRes, error) {
	city, err := h.svc.CreateCity(ctx, catalogService.CreateCityParams{
		Name:     req.GetName(),
		RegionID: req.GetRegionID(),
	})
	if err != nil {
		return nil, fmt.Errorf("create city: %w", err)
	}

	return &api.V1CatalogCitiesPostCreated{
		Data: api.V1CatalogCitiesPostCreatedData{
			City: models.ConvertCityModelToApi(city),
		},
	}, nil
}
