package amoproxy

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/amo-sync/internal/gateway/amocrm/dto"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/models"
)

type CreateLeadParams struct {
	Name            string
	ResponsibleUser models.AmoUser
	Pipeline        models.AmoLeadPipeline
	Status          models.AmoLeadPipelineStatus
	CompanyID       int
	ContactID       int
	Tags            []string
	ExternalID      int
}

func (s *Service) CreateLead(ctx context.Context, params CreateLeadParams) (int, error) {
	id, err := s.createLead(ctx, params)
	if err != nil {
		return 0, fmt.Errorf("create lead: %w", err)
	}

	err = s.amoStore.Create(ctx, s.psql.DB(), params.ExternalID, id, models.EntityLead)
	if err != nil {
		return 0, fmt.Errorf("create external id amo id lead link: %w", err)
	}

	return id, nil
}

func (s *Service) createLead(ctx context.Context, params CreateLeadParams) (int, error) {
	createLeadsRequest := dto.CreateLeadsRequest{
		{
			Name:              params.Name,
			ResponsibleUserID: int(params.ResponsibleUser),
			StatusID:          int(params.Status),
			PipelineID:        int(params.Pipeline),
			Embedded:          dto.CreateLeadRequestEmbedded{},
		},
	}

	if params.ContactID != 0 {
		createLeadsRequest[0].Embedded.Contacts = []dto.CreateLeadRequestEmbeddedContacts{
			{
				ID: params.ContactID,
			},
		}
	}

	if params.CompanyID != 0 {
		createLeadsRequest[0].Embedded.Companies = []dto.CreateLeadRequestEmbeddedCompanies{
			{
				ID: params.CompanyID,
			},
		}
	}

	for _, tag := range params.Tags {
		createLeadsRequest[0].TagsToAdd = append(createLeadsRequest[0].TagsToAdd, dto.CreateLeadRequestTagsToAdd{
			Name: tag,
		})
	}

	resp, err := s.amoCRMGateway.CreateLeads(ctx, createLeadsRequest)
	if err != nil {
		return 0, fmt.Errorf("create lead: %w", err)
	}

	if len(resp.Embedded.Leads) == 0 {
		return 0, fmt.Errorf("create lead: empty response")
	}

	return resp.Embedded.Leads[0].ID, nil
}
