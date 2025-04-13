package catalog

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type CreateCityParams struct {
	Name     string
	RegionID int
}

func (s *Service) CreateCity(ctx context.Context, params CreateCityParams) (models.City, error) {
	var (
		city models.City
		err  error
	)

	err = s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		city, err = s.catalogStore.CreateCity(ctx, qe, store.CatalogCreateCityParams{
			Name:     params.Name,
			RegionID: params.RegionID,
		})
		if err != nil {
			return fmt.Errorf("create city: %w", err)
		}

		city, err = s.catalogStore.GetCityByID(ctx, qe, city.ID)
		if err != nil {
			return fmt.Errorf("get city: %w", err)
		}

		return nil
	})
	if err != nil {
		return models.City{}, fmt.Errorf("run transaction: %w", err)
	}

	return city, nil
}
