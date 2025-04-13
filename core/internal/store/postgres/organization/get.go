package organization

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	sq "github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/deduplicate"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (s *OrganizationStore) GetCustomer(ctx context.Context, qe store.QueryExecutor, organizationID int) (models.Organization, error) {
	builder := sq.Select(
		"o.id",
		"o.brand_name",
		"o.full_name",
		"o.short_name",
		"o.inn",
		"o.okpo",
		"o.ogrn",
		"o.kpp",
		"o.tax_code",
		"o.address",
		"o.avatar_url",
		"o.emails",
		"o.phones",
		"o.messengers",
		"o.verification_status",
		"o.is_contractor",
		"o.is_banned",
		"o.customer_info",
		"o.created_at",
		"o.updated_at",
	).From("organizations AS o").
		Where(sq.Eq{"o.id": organizationID}).
		PlaceholderFormat(sq.Dollar)

	var (
		organization models.Organization
		avatarURL    sql.NullString
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&organization.ID,
		&organization.BrandName,
		&organization.FullName,
		&organization.ShortName,
		&organization.INN,
		&organization.OKPO,
		&organization.OGRN,
		&organization.KPP,
		&organization.TaxCode,
		&organization.Address,
		&avatarURL,
		&organization.Emails,
		&organization.Phones,
		&organization.Messengers,
		&organization.VerificationStatus,
		&organization.IsContractor,
		&organization.IsBanned,
		&organization.CustomerInfo,
		&organization.CreatedAt,
		&organization.UpdatedAt,
	)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Organization{}, errstore.ErrOrganizationNotFound
		}
		return models.Organization{}, fmt.Errorf("query row: %w", err)
	}

	cities, err := s.catalogStore.GetCitiesByIDs(ctx, qe, deduplicate.Deduplicate(organization.CustomerInfo.CityIDs))
	if err != nil {

		return models.Organization{}, fmt.Errorf("get cities by ids: %w", err)
	}
	organization.CustomerInfo.Cities = cities

	return organization, nil
}

func (s *OrganizationStore) GetContractorByID(ctx context.Context, qe store.QueryExecutor, organizationID int) (models.Organization, error) {
	builder := sq.Select(
		"o.id",
		"o.brand_name",
		"o.full_name",
		"o.short_name",
		"o.inn",
		"o.okpo",
		"o.ogrn",
		"o.kpp",
		"o.tax_code",
		"o.address",
		"o.avatar_url",
		"o.emails",
		"o.phones",
		"o.messengers",
		"o.verification_status",
		"o.is_contractor",
		"o.is_banned",
		"o.contractor_info",
		"o.created_at",
		"o.updated_at",
	).From("organizations AS o").
		Where(sq.Eq{"o.id": organizationID}).
		PlaceholderFormat(sq.Dollar)

	var (
		organization models.Organization
		avatarURL    sql.NullString
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&organization.ID,
		&organization.BrandName,
		&organization.FullName,
		&organization.ShortName,
		&organization.INN,
		&organization.OKPO,
		&organization.OGRN,
		&organization.KPP,
		&organization.TaxCode,
		&organization.Address,
		&avatarURL,
		&organization.Emails,
		&organization.Phones,
		&organization.Messengers,
		&organization.VerificationStatus,
		&organization.IsContractor,
		&organization.IsBanned,
		&organization.ContractorInfo,
		&organization.CreatedAt,
		&organization.UpdatedAt)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Organization{}, errstore.ErrOrganizationNotFound
		}
		return models.Organization{}, fmt.Errorf("query row: %w", err)
	}

	organization.AvatarURL = avatarURL.String

	if !organization.IsContractor {
		return models.Organization{}, errstore.ErrOrganizationNotAContractor
	}

	var (
		measureIDs []int
		serviceIDs []int
	)

	for _, service := range organization.ContractorInfo.Services {
		measureIDs = append(measureIDs, service.MeasureID)
		serviceIDs = append(serviceIDs, service.ServiceID)
	}

	services, err := s.catalogStore.GetServicesByIDs(ctx, qe, deduplicate.Deduplicate(serviceIDs))
	if err != nil {
		return models.Organization{}, fmt.Errorf("get services by ids: %w", err)
	}

	measurements, err := s.catalogStore.GetMeasurementsByIDs(ctx, qe, deduplicate.Deduplicate(measureIDs))
	if err != nil {
		return models.Organization{}, fmt.Errorf("get measurements by ids: %w", err)
	}

	for i, serviceWithPrice := range organization.ContractorInfo.Services {
		if service, ok := services[serviceWithPrice.ServiceID]; ok {
			organization.ContractorInfo.Services[i].Service = service
		}

		if measure, ok := measurements[serviceWithPrice.MeasureID]; ok {
			organization.ContractorInfo.Services[i].Measure = measure
		}
	}

	cities, err := s.catalogStore.GetCitiesByIDs(ctx, qe, deduplicate.Deduplicate(organization.ContractorInfo.CityIDs))
	if err != nil {
		return models.Organization{}, fmt.Errorf("get cities by ids: %w", err)
	}
	organization.ContractorInfo.Cities = cities

	objects, err := s.catalogStore.GetObjectsByIDs(ctx, qe, deduplicate.Deduplicate(organization.ContractorInfo.ObjectIDs))
	if err != nil {
		return models.Organization{}, fmt.Errorf("get objects by ids: %w", err)
	}

	for _, object := range objects {
		organization.ContractorInfo.Objects = append(organization.ContractorInfo.Objects, object)
	}

	return organization, nil
}

func (s *OrganizationStore) GetContractors(ctx context.Context, qe store.QueryExecutor, params store.OrganizationContractorsGetParams) ([]models.Organization, error) {
	builder := sq.Select(
		"o.id",
		"o.brand_name",
		"o.full_name",
		"o.short_name",
		"o.inn",
		"o.okpo",
		"o.ogrn",
		"o.kpp",
		"o.tax_code",
		"o.address",
		"o.avatar_url",
		"o.emails",
		"o.phones",
		"o.messengers",
		"o.verification_status",
		"o.is_contractor",
		"o.is_banned",
		"o.contractor_info",
		"o.created_at",
		"o.updated_at",
	).From("organizations AS o").
		Where(sq.Eq{"o.is_contractor": true}).
		PlaceholderFormat(sq.Dollar)

	if params.Limit.Set {
		builder = builder.Limit(params.Limit.Value)
	}

	if params.Offset.Set {
		builder = builder.Offset(params.Offset.Value)
	}

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	organizations := []models.Organization{}
	for rows.Next() {
		var (
			organization models.Organization
			avatarURL    sql.NullString
		)

		if err = rows.Scan(
			&organization.ID,
			&organization.BrandName,
			&organization.FullName,
			&organization.ShortName,
			&organization.INN,
			&organization.OKPO,
			&organization.OGRN,
			&organization.KPP,
			&organization.TaxCode,
			&organization.Address,
			&avatarURL,
			&organization.Emails,
			&organization.Phones,
			&organization.Messengers,
			&organization.VerificationStatus,
			&organization.IsContractor,
			&organization.IsBanned,
			&organization.ContractorInfo,
			&organization.CreatedAt,
			&organization.UpdatedAt); err != nil {
			if errors.Is(err, sql.ErrNoRows) {
				return []models.Organization{}, errstore.ErrOrganizationNotFound
			}
			return []models.Organization{}, fmt.Errorf("scan row: %w", err)
		}

		organization.AvatarURL = avatarURL.String

		if !organization.IsContractor {
			return []models.Organization{}, errstore.ErrOrganizationNotAContractor
		}

		var (
			measureIDs []int
			serviceIDs []int
		)

		for _, service := range organization.ContractorInfo.Services {
			measureIDs = append(measureIDs, service.MeasureID)
			serviceIDs = append(serviceIDs, service.ServiceID)
		}

		services, err := s.catalogStore.GetServicesByIDs(ctx, qe, deduplicate.Deduplicate(serviceIDs))
		if err != nil {
			return []models.Organization{}, fmt.Errorf("get services by ids: %w", err)
		}

		measurements, err := s.catalogStore.GetMeasurementsByIDs(ctx, qe, deduplicate.Deduplicate(measureIDs))
		if err != nil {
			return []models.Organization{}, fmt.Errorf("get measurements by ids: %w", err)
		}

		for i, serviceWithPrice := range organization.ContractorInfo.Services {
			if service, ok := services[serviceWithPrice.ServiceID]; ok {
				organization.ContractorInfo.Services[i].Service = service
			}

			if measure, ok := measurements[serviceWithPrice.MeasureID]; ok {
				organization.ContractorInfo.Services[i].Measure = measure
			}
		}

		cities, err := s.catalogStore.GetCitiesByIDs(ctx, qe, deduplicate.Deduplicate(organization.ContractorInfo.CityIDs))
		if err != nil {
			return []models.Organization{}, fmt.Errorf("get cities by ids: %w", err)
		}
		organization.ContractorInfo.Cities = cities

		objects, err := s.catalogStore.GetObjectsByIDs(ctx, qe, deduplicate.Deduplicate(organization.ContractorInfo.ObjectIDs))
		if err != nil {
			return []models.Organization{}, fmt.Errorf("get objects by ids: %w", err)
		}

		for _, object := range objects {
			organization.ContractorInfo.Objects = append(organization.ContractorInfo.Objects, object)
		}

		organizations = append(organizations, organization)
	}

	return organizations, nil
}

func (s *OrganizationStore) Get(ctx context.Context, qe store.QueryExecutor, params store.OrganizationGetParams) ([]models.Organization, error) {
	builder := sq.Select(
		"o.id",
		"o.brand_name",
		"o.full_name",
		"o.short_name",
		"o.inn",
		"o.okpo",
		"o.ogrn",
		"o.kpp",
		"o.tax_code",
		"o.address",
		"o.avatar_url",
		"o.emails",
		"o.phones",
		"o.messengers",
		"o.verification_status",
		"o.is_contractor",
		"o.is_banned",
		"o.created_at",
		"o.updated_at").
		From("organizations AS o").
		PlaceholderFormat(sq.Dollar)

	if params.Limit.Set {
		builder = builder.Limit(params.Limit.Value)
	}

	if params.Offset.Set {
		builder = builder.Offset(params.Offset.Value)
	}

	if params.IsContractor.Set {
		builder = builder.Where(
			sq.Eq{
				"o.is_contractor":       params.IsContractor.Value,
				"o.verification_status": models.VerificationStatusApproved,
				"o.is_banned":           false,
			})
	}

	if len(params.OrganizationIDs) != 0 {
		builder = builder.Where(sq.Eq{"o.id": params.OrganizationIDs})
	}

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	organizations := []models.Organization{}
	for rows.Next() {
		var (
			organization models.Organization
			avatarURL    sql.NullString
		)

		if err = rows.Scan(
			&organization.ID,
			&organization.BrandName,
			&organization.FullName,
			&organization.ShortName,
			&organization.INN,
			&organization.OKPO,
			&organization.OGRN,
			&organization.KPP,
			&organization.TaxCode,
			&organization.Address,
			&avatarURL,
			&organization.Emails,
			&organization.Phones,
			&organization.Messengers,
			&organization.VerificationStatus,
			&organization.IsContractor,
			&organization.IsBanned,
			&organization.CreatedAt,
			&organization.UpdatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		organization.AvatarURL = avatarURL.String

		organizations = append(organizations, organization)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("row iteration: %w", rows.Err())
	}

	return organizations, nil
}

func (s *OrganizationStore) GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Organization, error) {
	organizations, err := s.Get(ctx, qe, store.OrganizationGetParams{
		OrganizationIDs: []int{id}})
	if err != nil {
		return models.Organization{}, fmt.Errorf("get organizations: %w", err)
	}

	if len(organizations) == 0 {
		return models.Organization{}, errstore.ErrOrganizationNotFound
	}

	return organizations[0], nil
}

func (s *OrganizationStore) GetIsContractorByID(ctx context.Context, qe store.QueryExecutor, id int) (bool, error) {
	builder := squirrel.Select("is_contractor ").
		Where(squirrel.Eq{"id": id}).
		From("organizations").
		PlaceholderFormat(squirrel.Dollar)

	var isContractor bool
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&isContractor); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return false, errstore.ErrOrganizationNotFound
		}
		return false, fmt.Errorf("db: %w", err)
	}

	return isContractor, nil
}
