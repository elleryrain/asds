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

func (s *CatalogStore) CreateObject(ctx context.Context, qe store.QueryExecutor, params store.CatalogCreateObjectParams) (models.Object, error) {
	builder := squirrel.
		Insert("objects").
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
		object   models.Object
		parentID sql.NullInt64
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&object.ID,
		&object.Name,
		&parentID,
	)
	if err != nil {
		return models.Object{}, fmt.Errorf("query row: %w", err)
	}

	object.ParentID = int(parentID.Int64)

	return object, nil
}

func (s *CatalogStore) GetObjects(ctx context.Context, qe store.QueryExecutor) (models.Objects, error) {
	builder := squirrel.
		Select(
			"o.id",
			"o.name",
			"o.parent_id",
		).
		From("objects o").
		LeftJoin("objects o2 ON o2.id = o.parent_id;").
		PlaceholderFormat(squirrel.Dollar)

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	var objects models.Objects

	for rows.Next() {
		var (
			object   models.Object
			parentID sql.NullInt64
		)

		err = rows.Scan(
			&object.ID,
			&object.Name,
			&parentID,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		object.ParentID = int(parentID.Int64)

		objects = append(objects, object)
	}

	return objects, nil
}

func (s *CatalogStore) GetObjectsByIDs(ctx context.Context, qe store.QueryExecutor, objectIDs []int) (map[int]models.Object, error) {
	query := `WITH RECURSIVE object_hierarchy AS (
		SELECT 
			id,
			name,
			parent_id
		FROM 
			objects
		WHERE 
			id = ANY($1::bigint[])

		UNION ALL

		SELECT 
			o.id,
			o.name,
			o.parent_id
		FROM 
			objects o
		JOIN 
			object_hierarchy sh ON o.id = sh.parent_id
	)

	SELECT
		id,
		parent_id,
		name
	FROM 
		object_hierarchy;`

	rows, err := qe.QueryContext(ctx, query, pq.Array(objectIDs))
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	objects := map[int]models.Object{}

	for rows.Next() {
		var (
			object   models.Object
			parentID sql.NullInt64
		)

		err = rows.Scan(
			&object.ID,
			&parentID,
			&object.Name,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		object.ParentID = int(parentID.Int64)

		objects[object.ID] = object
	}

	return objects, nil
}
