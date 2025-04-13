package portfolio

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Update(ctx context.Context, params service.PortfolioUpdateParams) (models.Portfolio, error) {
	portfolio, err := s.portfolioStore.GetByID(ctx, s.psql.DB(), params.PortfolioID)
	if err != nil {
		return models.Portfolio{}, fmt.Errorf("get portfolio by id: %w", err)
	}

	if portfolio.OrganizationID != contextor.GetOrganizationID(ctx) {
		return models.Portfolio{}, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to update the portfolio", nil)
	}

	return s.portfolioStore.Update(ctx, s.psql.DB(), store.PortfolioUpdateParams{
		PortfolioID: params.PortfolioID,
		Title:       params.Title,
		Description: params.Description,
		Attachments: params.Attachments,
	})
}
