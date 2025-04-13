package tender

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Update(ctx context.Context, params service.TenderUpdateParams) (models.Tender, error) {
	tender, err := s.tenderStore.GetByID(ctx, s.psql.DB(), params.ID)
	if err != nil {
		return models.Tender{}, fmt.Errorf("get tender: %w", err)
	}

	if tender.Organization.ID != params.OrganizationID {
		cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to edit this tender", nil)
	}

	id, err := s.tenderStore.Update(ctx, s.psql.DB(), store.TenderUpdateParams{
		ID:              params.ID,
		Name:            params.Name,
		ServiceIDs:      params.ServiceIDs,
		ObjectIDs:       params.ObjectIDs,
		Price:           params.Price,
		IsContractPrice: params.IsContractPrice,
		IsNDSPrice:      params.IsNDSPrice,
		IsDraft:         params.IsDraft,
		CityID:          params.CityID,
		FloorSpace:      params.FloorSpace,
		Description:     params.Description,
		Wishes:          params.Wishes,
		Specification:   params.Specification,
		Attachments:     params.Attachments,
		ReceptionStart:  params.ReceptionStart,
		ReceptionEnd:    params.ReceptionEnd,
		WorkStart:       params.WorkStart,
		WorkEnd:         params.WorkEnd,
	})
	if err != nil {
		return models.Tender{}, fmt.Errorf("update tender: %w", err)
	}

	tender, err = s.tenderStore.GetByID(ctx, s.psql.DB(), id)
	if err != nil {
		return models.Tender{}, fmt.Errorf("get tender: %w", err)
	}

	return tender, nil
}
