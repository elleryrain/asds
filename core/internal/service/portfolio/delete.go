package portfolio

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
)

func (s *Service) Delete(ctx context.Context, portfolioID int) error {
	portfolio, err := s.portfolioStore.GetByID(ctx, s.psql.DB(), portfolioID)
	if err != nil {
		return fmt.Errorf("get portfolio by id: %w", err)
	}

	if portfolio.OrganizationID != contextor.GetOrganizationID(ctx) {
		return cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to delete the portfolio", nil)
	}

	return s.portfolioStore.Delete(ctx, s.psql.DB(), portfolioID)
}
