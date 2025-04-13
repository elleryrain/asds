package favourite

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (s *FavouriteStore) Get(ctx context.Context, qe store.QueryExecutor, params store.FavouriteGetParams) ([]models.Favourite[models.FavouriteObject], error) {
	builder := squirrel.Select(
		"f.id",
		"f.organization_id",
		"f.object_type",
		"f.object_id").
		From("favourites AS f").
		PlaceholderFormat(squirrel.Dollar)

	if params.FavouriteID.Set {
		builder = builder.Where(squirrel.Eq{"f.id": params.FavouriteID.Value})
	}

	if params.OrganizationID.Set {
		builder = builder.Where(squirrel.Eq{"f.organization_id": params.OrganizationID.Value})
	}

	if params.ObjectType.Set {
		builder = builder.Where(squirrel.Eq{"f.object_type": params.ObjectType.Value})
	}

	if params.Limit.Set {
		builder = builder.Limit(params.Limit.Value)
	}

	if params.Offset.Set {
		builder = builder.Offset(params.Offset.Value)
	}

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query context: %w", err)
	}
	defer rows.Close()

	favourites := []models.Favourite[models.FavouriteObject]{}
	for rows.Next() {
		var favourite models.Favourite[models.FavouriteObject]

		if err := rows.Scan(
			&favourite.ID,
			&favourite.OrganizationID,
			&favourite.ObjectType,
			&favourite.ObjectID,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		favourites = append(favourites, favourite)
	}

	return favourites, nil
}

func (s *FavouriteStore) GetByID(ctx context.Context, qe store.QueryExecutor, favouriteID int) (models.Favourite[models.FavouriteObject], error) {
	favourites, err := s.Get(ctx, qe, store.FavouriteGetParams{
		FavouriteID: models.NewOptional(favouriteID)})
	if err != nil {
		return models.Favourite[models.FavouriteObject]{}, fmt.Errorf("get objects: %w", err)
	}

	if len(favourites) == 0 {
		return models.Favourite[models.FavouriteObject]{}, errstore.ErrFavouriteNotFound
	}

	return favourites[0], nil
}
