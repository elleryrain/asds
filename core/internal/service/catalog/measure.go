package catalog

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (s *Service) GetMeasurements(ctx context.Context) ([]models.Measure, error) {
	return s.catalogStore.GetMeasurements(ctx, s.psql.DB())
}
