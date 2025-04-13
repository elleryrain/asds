package organization

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (s *Service) UpdateBrand(ctx context.Context, params service.OrganizationUpdateBrandParams) error {
	return s.organizationStore.Update(ctx, s.psql.DB(), store.OrganizationUpdateParams{
		OrganizationID: params.OrganizationID,
		Brand:          params.Brand,
		AvatarURL:      params.AvatarURL,
	})
}

func (s *Service) UpdateContacts(ctx context.Context, params service.OrganizationUpdateContactsParams) error {
	return s.organizationStore.Update(ctx, s.psql.DB(), store.OrganizationUpdateParams{
		OrganizationID: params.OrganizationID,
		Emails:         params.Emails,
		Phones:         params.Phones,
		Messengers:     params.Messengers,
	})
}

func (s *Service) UpdateCustomer(ctx context.Context, params service.OrganizationUpdateCustomerParams) (models.Organization, error) {
	if err := s.organizationStore.Update(ctx, s.psql.DB(), store.OrganizationUpdateParams{
		OrganizationID: params.OrganizationID,
		CustomerInfo: models.NewOptional(models.CustomerInfo{
			Description: params.Description,
			CityIDs:     params.CityIDs,
		}),
	}); err != nil {
		return models.Organization{}, fmt.Errorf("update organization: %w", err)
	}

	return s.organizationStore.GetCustomer(ctx, s.psql.DB(), params.OrganizationID)
}

func (s *Service) UpdateContractor(ctx context.Context, params service.OrganizationUpdateContractorParams) (models.Organization, error) {
	organization, err := s.organizationStore.GetByID(ctx, s.psql.DB(), params.OrganizationID)
	if err != nil {
		return models.Organization{}, fmt.Errorf("get organization by id: %w", err)
	}

	if !organization.IsContractor {
		return models.Organization{}, errstore.ErrOrganizationNotAContractor
	}

	if err := s.organizationStore.Update(ctx, s.psql.DB(), store.OrganizationUpdateParams{
		OrganizationID: params.OrganizationID,
		ContractorInfo: models.NewOptional(models.ContractorInfo{
			Description: params.Description,
			CityIDs:     params.CityIDs,
			ObjectIDs:   params.ObjectIDs,
			Services:    params.Services,
		}),
	}); err != nil {
		return models.Organization{}, fmt.Errorf("update organization: %w", err)
	}

	return s.organizationStore.GetContractorByID(ctx, s.psql.DB(), params.OrganizationID)
}
