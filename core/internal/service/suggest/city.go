package suggest

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (s *Service) City(ctx context.Context, name string) ([]models.City, error) {
	return s.catalogStore.ListCities(ctx, s.psql.DB(), name)
}
