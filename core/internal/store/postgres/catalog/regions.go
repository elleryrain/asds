package catalog

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *CatalogStore) CreateRegion(ctx context.Context, qe store.QueryExecutor, params store.CatalogCreateRegionParams) (models.Region, error) {
	builder := squirrel.
		Insert("regions").
		Columns(
			"name",
		).
		Values(
			params.Name,
		).
		Suffix(`
			RETURNING
				id,
				name
		`).
		PlaceholderFormat(squirrel.Dollar)

	var createdRegion models.Region

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdRegion.ID,
		&createdRegion.Name,
	)
	if err != nil {
		return models.Region{}, fmt.Errorf("query row: %w", err)
	}

	return createdRegion, nil
}

func (s *CatalogStore) GetRegionByID(ctx context.Context, qe store.QueryExecutor, regionID int) (models.Region, error) {
	builder := squirrel.
		Select(
			"id",
			"name",
		).
		From("regions").
		Where(squirrel.Eq{"id": regionID}).
		PlaceholderFormat(squirrel.Dollar)

	var region models.Region

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&region.ID,
		&region.Name,
	)
	if err != nil {
		return models.Region{}, fmt.Errorf("query row: %w", err)
	}

	return region, nil
}
