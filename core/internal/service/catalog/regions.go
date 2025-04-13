package catalog

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type CreateRegionParams struct {
	Name string
}

func (s *Service) CreateRegion(ctx context.Context, params CreateRegionParams) (models.Region, error) {
	region, err := s.catalogStore.CreateRegion(ctx, s.psql.DB(), store.CatalogCreateRegionParams{
		Name: params.Name,
	})
	if err != nil {
		return models.Region{}, fmt.Errorf("create region: %w", err)
	}

	return region, nil
}
