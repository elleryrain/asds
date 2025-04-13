package favourite

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Get(ctx context.Context, params service.FavouriteGetParams) (models.FavouritePagination[models.FavouriteObject], error) {
	favourites, err := s.favouriteStore.Get(ctx, s.psql.DB(), store.FavouriteGetParams{
		OrganizationID: models.NewOptional(params.OrganizationID),
		ObjectType:     models.NewOptional(params.ObjectType),
		Limit:          models.NewOptional(params.PerPage),
		Offset:         models.Optional[uint64]{Value: params.Page * params.PerPage, Set: (params.Page * params.PerPage) != 0},
	})
	if err != nil {
		return models.FavouritePagination[models.FavouriteObject]{}, fmt.Errorf("get object req: %w", err)
	}

	count, err := s.favouriteStore.Count(ctx, s.psql.DB(), store.FavouriteGetCountParams{
		OrganizationID: params.OrganizationID,
		ObjectType:     params.ObjectType,
	})
	if err != nil {
		return models.FavouritePagination[models.FavouriteObject]{}, fmt.Errorf("get count favourites: %w", err)
	}

	var objectID []int
	for _, req := range favourites {
		objectID = append(objectID, req.ObjectID)
	}

	objectMap := map[int]models.FavouriteObject{}
	switch params.ObjectType {
	case models.FavouriteTypeOrganization:
		organizations, err := s.organizationStore.GetContractors(ctx, s.psql.DB(), store.OrganizationContractorsGetParams{})
		if err != nil {
			return models.FavouritePagination[models.FavouriteObject]{}, fmt.Errorf("get organizations: %w", err)
		}

		for _, organization := range organizations {
			objectMap[organization.ID] = models.OrganizationWithProfile{
				Organization: organization,
				Profile:      organization.ContractorInfo,
			}
		}

	case models.FavouriteTypeTender:
		tenders, err := s.tenderStore.List(ctx, s.psql.DB(), store.TenderListParams{
			TenderIDs: models.Optional[[]int]{Value: objectID, Set: true}})
		if err != nil {
			return models.FavouritePagination[models.FavouriteObject]{}, fmt.Errorf("get tenders: %w", err)
		}

		for _, tender := range tenders {
			objectMap[tender.ID] = tender
		}
	}

	for i := range favourites {
		if object, ok := objectMap[favourites[i].ObjectID]; ok {
			favourites[i].Object = object
		}
	}

	return models.FavouritePagination[models.FavouriteObject]{
		Favourites: favourites,
		Pagination: pagination.New(params.Page, params.PerPage, uint64(count)),
	}, nil
}
