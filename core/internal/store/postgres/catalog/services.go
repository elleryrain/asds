package catalog

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *CatalogStore) CreateService(ctx context.Context, qe store.QueryExecutor, params store.CatalogCreateServiceParams) (models.Service, error) {
	builder := squirrel.
		Insert("services").
		Columns(
			"name",
			"parent_id",
		).
		Values(
			params.Name,
			sql.NullInt64{Int64: int64(params.ParentID), Valid: params.ParentID != 0},
		).
		Suffix(`
			RETURNING
				id,
				name,
				parent_id
		`).
		PlaceholderFormat(squirrel.Dollar)

	var (
		service  models.Service
		parentID sql.NullInt64
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&service.ID,
		&service.Name,
		&parentID,
	)
	if err != nil {
		return models.Service{}, fmt.Errorf("query row: %w", err)
	}

	service.ParentID = int(parentID.Int64)

	return service, nil
}

func (s *CatalogStore) GetServices(ctx context.Context, qe store.QueryExecutor) (models.Services, error) {
	builder := squirrel.
		Select(
			"s.id",
			"s.name",
			"s.parent_id",
		).
		From("services s").
		LeftJoin("services s2 ON s2.id = s.parent_id;").
		PlaceholderFormat(squirrel.Dollar)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	var services models.Services

	for rows.Next() {
		var (
			service  models.Service
			parentID sql.NullInt64
		)

		err = rows.Scan(
			&service.ID,
			&service.Name,
			&parentID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		service.ParentID = int(parentID.Int64)

		services = append(services, service)
	}

	return services, nil
}

func (s *CatalogStore) GetServicesByIDs(ctx context.Context, qe store.QueryExecutor, serviceIDs []int) (map[int]models.Service, error) {
	query := `WITH RECURSIVE service_hierarchy AS (
		SELECT 
			id,
			parent_id,
			name
		FROM 
			services
		WHERE 
			id = ANY($1::bigint[])

		UNION ALL

		SELECT 
			s.id,
			s.parent_id,
			s.name
		FROM 
			services s
		JOIN 
			service_hierarchy sh ON s.id = sh.parent_id
	)

	SELECT
		id,
		parent_id,
		name
	FROM
		service_hierarchy;`

	rows, err := qe.QueryContext(ctx, query, pq.Array(serviceIDs))
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	services := map[int]models.Service{}

	for rows.Next() {
		var (
			service  models.Service
			parentID sql.NullInt64
		)

		err = rows.Scan(
			&service.ID,
			&parentID,
			&service.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		service.ParentID = int(parentID.Int64)

		services[service.ID] = service
	}

	return services, nil
}
