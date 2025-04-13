package catalog

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) GetObjects(ctx context.Context) (models.Objects, error) {
	objects, err := s.catalogStore.GetObjects(ctx, s.psql.DB())
	if err != nil {
		return nil, fmt.Errorf("get objects catalog: %w", err)
	}

	return objects, nil
}

type CreateObjectParams struct {
	Name     string
	ParentID int
}

func (s *Service) CreateObject(ctx context.Context, params CreateObjectParams) (models.Object, error) {
	object, err := s.catalogStore.CreateObject(ctx, s.psql.DB(), store.CatalogCreateObjectParams{
		Name:     params.Name,
		ParentID: params.ParentID,
	})
	if err != nil {
		return models.Object{}, fmt.Errorf("create object: %w", err)
	}

	return object, nil
}
