package favourite

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *FavouriteStore) Create(ctx context.Context, qe store.QueryExecutor, params store.FavouriteCreateParams) (int64, error) {
	builder := squirrel.Insert("favourites").Columns(
		"organization_id",
		"object_type",
		"object_id",
	).Values(
		params.OrganizationID,
		params.ObjectType,
		params.ObjectID,
	).PlaceholderFormat(squirrel.Dollar).Suffix("RETURNING id")

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	var id int64
	if rows.Next() {
		if err := rows.Scan(&id); err != nil {
			return 0, fmt.Errorf("scan id: %w", err)
		}
	}

	return id, nil
}
