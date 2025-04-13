package favourite

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
)

func (s *Service) Delete(ctx context.Context, favouriteID int) error {
	favorite, err := s.favouriteStore.GetByID(ctx, s.psql.DB(), favouriteID)
	if err != nil {
		return fmt.Errorf("get favorite by id: %w", err)
	}

	if favorite.OrganizationID != contextor.GetOrganizationID(ctx) {
		return cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to delete object from favourite", nil)
	}

	return s.favouriteStore.Delete(ctx, s.psql.DB(), favouriteID)
}
