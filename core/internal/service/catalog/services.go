package catalog

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) GetServices(ctx context.Context) (models.Services, error) {
	services, err := s.catalogStore.GetServices(ctx, s.psql.DB())
	if err != nil {
		return nil, fmt.Errorf("get services catalog: %w", err)
	}

	return services, nil
}

type CreateServiceParams struct {
	Name     string
	ParentID int
}

func (s *Service) CreateService(ctx context.Context, params CreateServiceParams) (models.Service, error) {
	service, err := s.catalogStore.CreateService(ctx, s.psql.DB(), store.CatalogCreateServiceParams{
		Name:     params.Name,
		ParentID: params.ParentID,
	})
	if err != nil {
		return models.Service{}, fmt.Errorf("create service: %w", err)
	}

	return service, nil
}
