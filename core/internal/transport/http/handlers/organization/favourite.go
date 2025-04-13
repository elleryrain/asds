package organization

import (
	"context"
	"errors"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (h *Handler) V1OrganizationsOrganizationIDFavouritesPost(
	ctx context.Context,
	req *api.V1OrganizationsOrganizationIDFavouritesPostReq,
	params api.V1OrganizationsOrganizationIDFavouritesPostParams) (
	api.V1OrganizationsOrganizationIDFavouritesPostRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to create a favourite", nil)
	}

	id, err := h.favouriteService.Create(ctx, service.FavouriteCreateParams{
		OrganizationID: params.OrganizationID,
		ObjectType:     models.APIToFavouriteType(params.ObjectType),
		ObjectID:       req.ObjectID,
	})
	if err != nil {
		return nil, fmt.Errorf("create favourite: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDFavouritesPostCreated{
		ID: models.ToInt(id),
	}, nil
}

func (h *Handler) V1OrganizationsOrganizationIDFavouritesGet(
	ctx context.Context,
	params api.V1OrganizationsOrganizationIDFavouritesGetParams) (
	api.V1OrganizationsOrganizationIDFavouritesGetRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to get a favourite", nil)
	}

	favourites, err := h.favouriteService.Get(ctx, service.FavouriteGetParams{
		OrganizationID: params.OrganizationID,
		ObjectType:     models.ApiToFavouriteType(params.ObjectType),
		Page:           uint64(params.Page.Or(pagination.Page)),
		PerPage:        uint64(params.PerPage.Or(pagination.PerPage))})
	if err != nil {
		return nil, fmt.Errorf("get favourite: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDFavouritesGetOK{
		Data:       convert.Slice[[]models.Favourite[models.FavouriteObject], []api.Favourites](favourites.Favourites, models.FavouriteModelToApi),
		Pagination: pagination.ConvertPaginationToAPI(favourites.Pagination),
	}, nil
}

func (h *Handler) V1OrganizationsFavouritesFavouriteIDDelete(
	ctx context.Context,
	params api.V1OrganizationsFavouritesFavouriteIDDeleteParams) (
	api.V1OrganizationsFavouritesFavouriteIDDeleteRes, error) {
	if err := h.favouriteService.Delete(ctx, params.FavouriteID); err != nil {
		if errors.Is(err, errstore.ErrFavouriteNotFound) {
			return nil, cerr.Wrap(err, cerr.CodeNotFound, "Избранное не существует", nil)
		}
		return nil, fmt.Errorf("favourite delete: %w", err)
	}

	return &api.V1OrganizationsFavouritesFavouriteIDDeleteNoContent{}, nil
}
