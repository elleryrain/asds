package service

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type FavouriteCreateParams struct {
	OrganizationID int
	ObjectType     models.FavouriteType
	ObjectID       int
}

type FavouriteGetParams struct {
	OrganizationID int
	ObjectType     models.FavouriteType
	ObjectID       models.Optional[int]
	Page           uint64
	PerPage        uint64
}
