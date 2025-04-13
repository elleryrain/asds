package portfolio

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Create(ctx context.Context, params service.PortfolioCreateParams) (models.Portfolio, error) {	
	count, err := s.portfolioStore.Count(ctx, s.psql.DB(), params.OrganizationID)
	if err != nil {
		return models.Portfolio{}, fmt.Errorf("failed to count portfolios: %w", err)
	}

	if count >= 10 {
		return models.Portfolio{}, fmt.Errorf("portfolio limit reached: maximum 10 portfolios allowed")
	}

	return s.portfolioStore.Create(ctx, s.psql.DB(), store.PortfolioCreateParams{
		OrganizationID: params.OrganizationID,
		Title:          params.Title,
		Description:    params.Description,
		Attachments:    params.Attachments,
	})
}
