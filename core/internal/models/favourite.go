package models

import (
	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
)

type FavouriteType int

const (
	FavouriteTypeInvalid FavouriteType = iota
	FavouriteTypeOrganization
	FavouriteTypeTender
)

var mapFavouriteType = map[FavouriteType]api.FavouriteType{
	FavouriteTypeOrganization: api.FavouriteTypeOrganization,
	FavouriteTypeTender:       api.FavouriteTypeTender,
}

var mapApiToFavouriteType = map[api.FavouriteType]FavouriteType{
	api.FavouriteTypeOrganization: FavouriteTypeOrganization,
	api.FavouriteTypeTender:       FavouriteTypeTender,
}

func ApiToFavouriteType(apiType api.FavouriteType) FavouriteType {
	objectType, ok := mapApiToFavouriteType[apiType]
	if !ok {
		return FavouriteTypeInvalid
	}
	return objectType
}

func (f FavouriteType) ToAPI() api.FavouriteType {
	return mapFavouriteType[f]
}

func APIToFavouriteType(apiType api.FavouriteType) FavouriteType {
	favourite, ok := mapApiToFavouriteType[apiType]
	if !ok {
		return FavouriteTypeInvalid
	}

	return favourite
}

type FavouriteObject interface {
	ToFavouriteObject() api.FavouritesObject
}

type FavouritePagination[T FavouriteObject] struct {
	Favourites []Favourite[T]
	Pagination pagination.Pagination
}

type Favourite[T FavouriteObject] struct {
	ID             int
	OrganizationID int
	ObjectType     FavouriteType
	ObjectID       int
	Object         T
}

func FavouriteModelToApi[T FavouriteObject](favourite Favourite[T]) api.Favourites {
	return api.Favourites{
		ID:             favourite.ID,
		OrganizationID: favourite.OrganizationID,
		ObjectType:     favourite.ObjectType.ToAPI(),
		Object:         favourite.Object.ToFavouriteObject(),
	}
}

func ToInt(id int64) int {
	return int(id)
}
