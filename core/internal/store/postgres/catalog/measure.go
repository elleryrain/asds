package catalog

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *CatalogStore) GetMeasurements(ctx context.Context, qe store.QueryExecutor) ([]models.Measure, error) {
	builder := squirrel.
		Select(
			"m.id",
			"m.name",
		).
		From("measurements m").
		PlaceholderFormat(squirrel.Dollar)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	var measurements []models.Measure
	for rows.Next() {
		var measure models.Measure

		if err := rows.Scan(
			&measure.ID,
			&measure.Name,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		measurements = append(measurements, measure)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration error: %w", err)
	}

	return measurements, nil
}

func (s *CatalogStore) GetMeasurementsByIDs(ctx context.Context, qe store.QueryExecutor, measureIDs []int) (map[int]models.Measure, error) {
	builder := squirrel.Select(
		"id",
		"name").
		From("measurements").
		Where(squirrel.Eq{"id": measureIDs}).
		PlaceholderFormat(squirrel.Dollar)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query measurements: %w", err)
	}
	defer rows.Close()

	measurements := make(map[int]models.Measure)
	for rows.Next() {
		var measurement models.Measure
		if err := rows.Scan(
			&measurement.ID,
			&measurement.Name,
		); err != nil {
			return nil, fmt.Errorf("scan measurement row: %w", err)
		}

		measurements[measurement.ID] = measurement
	}

	return measurements, nil
}
