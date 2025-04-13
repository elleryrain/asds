package tender

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) List(ctx context.Context, params service.TenderListParams) (models.TendersPagination, error) {
	tenders, err := s.tenderStore.List(ctx, s.psql.DB(), store.TenderListParams{
		OrganizationID: params.OrganizationID,
		VerifiedOnly:   params.VerifiedOnly,
		WithDrafts:     params.WithDrafts,
		Limit:          models.NewOptional(params.PerPage),
		Offset:         models.Optional[uint64]{Value: params.Page * params.PerPage, Set: (params.Page * params.PerPage) != 0}})
	if err != nil {
		return models.TendersPagination{}, fmt.Errorf("get tenders: %w", err)
	}

	count, err := s.tenderStore.Count(ctx, s.psql.DB(), store.TenderGetCountParams{
		OrganizationID: params.OrganizationID,
		VerifiedOnly:   params.VerifiedOnly,
		WithDrafts:     params.WithDrafts})
	if err != nil {
		return models.TendersPagination{}, fmt.Errorf("get count tenders: %w", err)
	}

	return models.TendersPagination{
		Tenders: tenders,
		Pagination: pagination.New(params.Page, params.PerPage, uint64(count)),
	}, nil
}
