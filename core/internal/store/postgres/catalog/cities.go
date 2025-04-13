package catalog

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *CatalogStore) CreateCity(ctx context.Context, qe store.QueryExecutor, params store.CatalogCreateCityParams) (models.City, error) {
	builder := squirrel.
		Insert("cities").
		Columns(
			"name",
			"region_id",
		).
		Values(
			params.Name,
			params.RegionID,
		).
		Suffix(`
			RETURNING
				id,
				name,
				region_id
		`).
		PlaceholderFormat(squirrel.Dollar)

	var createdCity models.City

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdCity.ID,
		&createdCity.Name,
		&createdCity.Region.ID,
	)
	if err != nil {
		return models.City{}, fmt.Errorf("query row: %w", err)
	}

	return createdCity, nil
}

func (s *CatalogStore) GetCityByID(ctx context.Context, qe store.QueryExecutor, cityID int) (models.City, error) {
	builder := squirrel.
		Select(
			"c.id",
			"c.name",
			"c.region_id",
			"r.name",
		).
		From("cities c").
		Join("regions r ON r.id = c.region_id").
		Where(squirrel.Eq{"c.id": cityID}).
		PlaceholderFormat(squirrel.Dollar)

	var city models.City

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&city.ID,
		&city.Name,
		&city.Region.ID,
		&city.Region.Name,
	)
	if err != nil {
		return models.City{}, fmt.Errorf("query row: %w", err)
	}

	return city, nil
}

func (s *CatalogStore) ListCities(ctx context.Context, qe store.QueryExecutor, name string) ([]models.City, error) {
	builder := squirrel.
		Select(
			"c.id",
			"c.name",
			"c.region_id",
			"r.name",
		).
		From("cities c").
		Join("regions r ON r.id = c.region_id").
		PlaceholderFormat(squirrel.Dollar)

	if name != "" {
		builder = builder.Where(squirrel.ILike{"c.name": fmt.Sprintf("%s%%", name)})
	}

	var cities []models.City

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}

	for rows.Next() {
		var city models.City

		err = rows.Scan(
			&city.ID,
			&city.Name,
			&city.Region.ID,
			&city.Region.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		cities = append(cities, city)
	}

	return cities, nil
}

func (s *CatalogStore) GetCitiesByIDs(ctx context.Context, qe store.QueryExecutor, cityIDs []int) ([]models.City, error) {
	builder := squirrel.Select(
		"c.id",
		"c.name",
		"r.id AS region_id",
		"r.name AS region_name").
		From("cities AS c").
		Join("regions AS r ON c.region_id = r.id").
		Where(squirrel.Eq{"c.id": cityIDs}).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query cities: %w", err)
	}
	defer rows.Close()

	var cities []models.City
	for rows.Next() {
		var city models.City
		if err := rows.Scan(
			&city.ID,
			&city.Name,
			&city.Region.ID,
			&city.Region.Name,
		); err != nil {
			return nil, fmt.Errorf("scan city row: %w", err)
		}

		cities = append(cities, city)
	}

	return cities, nil
}
