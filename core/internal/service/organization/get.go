package organization

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Get(ctx context.Context, params service.OrganizationGetParams) (models.OrganizationsPagination, error) {
	organizations, err := s.organizationStore.Get(ctx, s.psql.DB(), store.OrganizationGetParams{
		IsContractor: params.IsContractor,
		Limit:        models.NewOptional(params.PerPage),
		Offset:       models.Optional[uint64]{Value: params.Page * params.PerPage, Set: (params.Page * params.PerPage) != 0}})
	if err != nil {
		return models.OrganizationsPagination{}, fmt.Errorf("get organizations: %w", err)
	}

	count, err := s.organizationStore.Count(ctx, s.psql.DB(), store.OrganizationGetCountParams{
		IsContractor: params.IsContractor})
	if err != nil {
		return models.OrganizationsPagination{}, fmt.Errorf("get count organizations: %w", err)
	}

	return models.OrganizationsPagination{
		Organizations: organizations,
		Pagination:    pagination.New(params.Page, params.PerPage, uint64(count)),
	}, nil
}

func (s *Service) GetByID(ctx context.Context, id int) (models.Organization, error) {
	return s.organizationStore.GetByID(ctx, s.psql.DB(), id)
}

func (s *Service) GetContractors(ctx context.Context, params service.OrganizationContractorsGetParams) (models.OrganizationsPagination, error) {
	organizations, err := s.organizationStore.GetContractors(ctx, s.psql.DB(), store.OrganizationContractorsGetParams{
		Limit:  models.NewOptional(params.PerPage),
		Offset: models.Optional[uint64]{Value: params.Page * params.PerPage, Set: (params.Page * params.PerPage) != 0}})
	if err != nil {
		return models.OrganizationsPagination{}, fmt.Errorf("get contactors organizations: %w", err)
	}

	count, err := s.organizationStore.Count(ctx, s.psql.DB(), store.OrganizationGetCountParams{
		IsContractor: models.Optional[bool]{Value: true, Set: true}})
	if err != nil {
		return models.OrganizationsPagination{}, fmt.Errorf("get count organizations: %w", err)
	}

	return models.OrganizationsPagination{
		Organizations: organizations,
		Pagination:    pagination.New(params.Page, params.PerPage, uint64(count)),
	}, nil
}
