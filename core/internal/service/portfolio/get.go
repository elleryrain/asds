package portfolio

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Get(ctx context.Context, params service.PortfolioGetParams) ([]models.Portfolio, error) {
	return s.portfolioStore.Get(ctx, s.psql.DB(), store.PortfolioGetParams{
		OrganizationID: models.NewOptional(params.OrganizationID),
	})
}