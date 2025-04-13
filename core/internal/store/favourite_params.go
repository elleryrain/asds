package store

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type FavouriteCreateParams struct {
	OrganizationID int
	ObjectType     models.FavouriteType
	ObjectID       int
}

type FavouriteGetParams struct {
	FavouriteID    models.Optional[int]
	OrganizationID models.Optional[int]
	ObjectType     models.Optional[models.FavouriteType]
	Offset         models.Optional[uint64]
	Limit          models.Optional[uint64]
}

type FavouriteGetCountParams struct {
	OrganizationID int
	ObjectType     models.FavouriteType
}
