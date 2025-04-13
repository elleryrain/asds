package tender

import (
	"context"
	"database/sql"
	"errors"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

func (s *Service) GetByID(ctx context.Context, tenderID int) (models.Tender, error) {
	organizationID := contextor.GetOrganizationID(ctx)

	tender, err := s.tenderStore.GetByID(ctx, s.psql.DB(), tenderID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Tender{}, cerr.Wrap(err, cerr.CodeNotFound, "tender not found", nil)
		}

		return models.Tender{}, err
	}

	if tender.IsDraft && tender.Organization.ID != organizationID {
		return models.Tender{}, cerr.Wrap(err, cerr.CodeNotPermitted, "not permitted get this tender", nil)
	}

	return tender, nil
}
