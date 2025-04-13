package suggest

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (h *Handler) V1SuggestCityGet(ctx context.Context, params api.V1SuggestCityGetParams) (api.V1SuggestCityGetRes, error) {
	cities, err := h.svc.City(ctx, params.Name.Value)
	if err != nil {
		return nil, fmt.Errorf("get cities: %w", err)
	}

	return &api.V1SuggestCityGetOK{
		Data: convert.Slice[[]models.City, []api.City](cities, models.ConvertCityModelToApi),
	}, nil
}
