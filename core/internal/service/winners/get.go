package winners

import (
	"context"
	"database/sql"
	"errors"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (s *Service) Get(ctx context.Context, tenderID int) ([]models.Winners, error) {
	tender, err := s.tenderStore.GetByID(ctx, s.psql.DB(), tenderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return []models.Winners{}, cerr.Wrap(err, cerr.CodeNotFound, "tender not found", nil)
		}
		return []models.Winners{}, err
	}

	if tender.Organization.ID != contextor.GetOrganizationID(ctx) {
		return []models.Winners{}, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to create the winner", nil)
	}

	return s.winnersStore.Get(ctx, s.psql.DB(), tenderID)
}
