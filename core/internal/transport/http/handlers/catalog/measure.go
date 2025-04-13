package catalog

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (h *Handler) V1CatalogMeasurementsGet(ctx context.Context) (api.V1CatalogMeasurementsGetRes, error) {
	measurements, err := h.svc.GetMeasurements(ctx)
	if err != nil {
		return nil, fmt.Errorf("get measurements: %w", err)
	}

	return &api.V1CatalogMeasurementsGetOK{
		Data: convert.Slice[[]models.Measure, []api.Measure](measurements, models.ConvertMeasureToAPI),
	}, nil
}
