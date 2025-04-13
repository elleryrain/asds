package favourite

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *FavouriteStore) Count(ctx context.Context, qe store.QueryExecutor, params store.FavouriteGetCountParams) (int, error) {
	builder := squirrel.Select("COUNT(*)").
		From("favourites AS f").
		Where(squirrel.Eq{"f.organization_id": params.OrganizationID}).
		Where(squirrel.Eq{"f.object_type": params.ObjectType}).
		PlaceholderFormat(squirrel.Dollar)

	var count int
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&count); err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return count, nil
}
