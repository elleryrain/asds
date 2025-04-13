package amoproxy

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/amo-sync/internal/gateway/amocrm/dto"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/models"
)

type CreateCompanyParams struct {
	Company         models.AmoCompany
	ResponsibleUser models.AmoUser
	ExternalID      int
}

func (s *Service) CreateCompany(ctx context.Context, params CreateCompanyParams) (int, error) {
	id, err := s.createCompany(ctx, params)
	if err != nil {
		return 0, fmt.Errorf("create company: %w", err)
	}

	err = s.amoStore.Create(ctx, s.psql.DB(), params.ExternalID, id, models.EntityCompany)
	if err != nil {
		return 0, fmt.Errorf("create external id amo id company link: %w", err)
	}

	return id, nil
}

func (s *Service) createCompany(ctx context.Context, params CreateCompanyParams) (int, error) {
	company := params.Company

	resp, err := s.amoCRMGateway.CreateCompanies(ctx, dto.CreateCompaniesRequest{
		{
			Name:              company.ShortName,
			ResponsibleUserID: int(params.ResponsibleUser),
			CustomFieldsValues: []dto.CustomFieldsValue{
				{
					FieldID: CompanyFieldIDBrandName,
					Values: []dto.Value{
						{
							Value: company.BrandName,
						},
					},
				},
				{
					FieldID: CompanyFieldIDShortName,
					Values: []dto.Value{
						{
							Value: company.ShortName,
						},
					},
				},
				{
					FieldID: CompanyFieldIDFullName,
					Values: []dto.Value{
						{
							Value: company.FullName,
						},
					},
				},
				{
					FieldID: CompanyFieldIDRole,
					Values: []dto.Value{
						{
							EnumID: int(company.Role),
						},
					},
				},
				{
					FieldID: CompanyFieldIDINN,
					Values: []dto.Value{
						{
							Value: company.INN,
						},
					},
				},
				{
					FieldID: CompanyFieldIDKPP,
					Values: []dto.Value{
						{
							Value: company.KPP,
						},
					},
				},
				{
					FieldID: CompanyFieldIDOGRN,
					Values: []dto.Value{
						{
							Value: company.OGRN,
						},
					},
				},
				{
					FieldID: CompanyFieldIDOKPO,
					Values: []dto.Value{
						{
							Value: company.OKPO,
						},
					},
				},
				{
					FieldID: CompanyFieldIDTaxCode,
					Values: []dto.Value{
						{
							Value: company.TaxCode,
						},
					},
				},
				{
					FieldID: CompanyFieldIDRegisteredAt,
					Values: []dto.Value{
						{
							Value: company.RegisteredAt.Unix(),
						},
					},
				},
			},
		},
	})
	if err != nil {
		return 0, fmt.Errorf("create companies: %w", err)
	}

	if len(resp.Embedded.Companies) == 0 {
		return 0, fmt.Errorf("create companies: empty response")
	}

	return resp.Embedded.Companies[0].ID, nil
}
