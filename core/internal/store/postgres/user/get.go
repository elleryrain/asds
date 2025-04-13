package user

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (s *UserStore) Get(ctx context.Context, qe store.QueryExecutor, params store.UserGetParams) ([]models.FullUser, error) {
	builder := squirrel.
		Select(
			"u.id",
			"u.email",
			"u.phone",
			"u.password_hash",
			"u.totp_salt",
			"u.first_name",
			"u.last_name",
			"u.middle_name",
			"u.avatar_url",
			"u.email_verified",
			"u.created_at",
			"u.updated_at",
			"o.id",
			"o.brand_name",
			"o.full_name",
			"o.short_name",
			"o.is_contractor",
			"o.is_banned",
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
			"o.created_at",
			"o.updated_at",
			"e.role",
			"e.position",
		).
		From("users AS u").
		LeftJoin("organization_users AS ou ON u.id = ou.user_id").
		LeftJoin("organizations AS o ON o.id = ou.organization_id").
		LeftJoin("employee AS e ON u.id = e.user_id").
		PlaceholderFormat(squirrel.Dollar)

	if params.Limit.Set {
		builder = builder.Limit(params.Limit.Value)
	}

	if params.Offset.Set {
		builder = builder.Offset(params.Offset.Value)
	}

	if params.Email != "" {
		builder = builder.Where(squirrel.Eq{"u.email": params.Email})
	}

	if params.ID != 0 {
		builder = builder.Where(squirrel.Eq{"u.id": params.ID})
	}

	var users []models.FullUser
	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}

	for rows.Next() {
		var (
			user           models.FullUser
			userMiddleName sql.NullString
			avatarURL      sql.NullString

			organizationID           sql.NullInt64
			organizationBrandName    sql.NullString
			organizationFullName     sql.NullString
			organizationShortName    sql.NullString
			organizationIsContractor sql.NullBool
			organizationIsBanned     sql.NullBool
			organizationINN          sql.NullString
			organizationOKPO         sql.NullString
			organizationOGRN         sql.NullString
			organizationKPP          sql.NullString
			organizationTaxCode      sql.NullString
			organizationAddress      sql.NullString
			organizationAvatarURL    sql.NullString
			organizationCreatedAt    sql.NullTime
			organizationUpdatedAt    sql.NullTime
			employeePosition         sql.NullString
			employeeRole             sql.NullInt16
		)

		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Phone,
			&user.PasswordHash,
			&user.TOTPSalt,
			&user.FirstName,
			&user.LastName,
			&userMiddleName,
			&avatarURL,
			&user.EmailVerified,
			&user.CreatedAt,
			&user.UpdatedAt,
			&organizationID,
			&organizationBrandName,
			&organizationFullName,
			&organizationShortName,
			&organizationIsContractor,
			&organizationIsBanned,
			&organizationINN,
			&organizationOKPO,
			&organizationOGRN,
			&organizationKPP,
			&organizationTaxCode,
			&organizationAddress,
			&organizationAvatarURL,
			&user.Organization.Emails,
			&user.Organization.Phones,
			&user.Organization.Messengers,
			&organizationCreatedAt,
			&organizationUpdatedAt,
			&employeeRole,
			&employeePosition,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		user.AvatarURL = avatarURL.String

		user.Organization.ID = int(organizationID.Int64)
		user.Organization.BrandName = organizationBrandName.String
		user.Organization.FullName = organizationFullName.String
		user.Organization.ShortName = organizationShortName.String
		user.Organization.IsContractor = organizationIsContractor.Bool
		user.Organization.IsBanned = organizationIsBanned.Bool
		user.Organization.INN = organizationINN.String
		user.Organization.OKPO = organizationOKPO.String
		user.Organization.OGRN = organizationOGRN.String
		user.Organization.KPP = organizationKPP.String
		user.Organization.TaxCode = organizationTaxCode.String
		user.Organization.Address = organizationAddress.String
		user.Organization.AvatarURL = organizationAvatarURL.String
		user.Organization.CreatedAt = organizationCreatedAt.Time
		user.Organization.UpdatedAt = organizationUpdatedAt.Time

		user.Position = employeePosition.String
		user.Role = models.UserRole(employeeRole.Int16)

		if userMiddleName.Valid {
			user.MiddleName = models.Optional[string]{Value: userMiddleName.String, Set: true}
		}

		users = append(users, user)
	}

	return users, nil
}

func (s *UserStore) GetWithOrganiztion(ctx context.Context, qe store.QueryExecutor, params store.UserGetParams) ([]models.RegularUser, error) {
	builder := squirrel.
		Select(
			"u.id",
			"u.email",
			"u.phone",
			"u.password_hash",
			"u.totp_salt",
			"u.first_name",
			"u.last_name",
			"u.middle_name",
			"u.avatar_url",
			"u.email_verified",
			"u.created_at",
			"u.updated_at",
			"o.id",
			"o.brand_name",
			"o.full_name",
			"o.short_name",
			"o.is_contractor",
			"o.is_banned",
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
			"o.created_at",
			"o.updated_at",
		).
		From("users AS u").
		Join("organization_users AS ou ON u.id = ou.user_id").
		Join("organizations AS o ON o.id = ou.organization_id").
		PlaceholderFormat(squirrel.Dollar)

	if params.Limit.Set {
		builder = builder.Limit(params.Limit.Value)
	}

	if params.Offset.Set {
		builder = builder.Offset(params.Offset.Value)
	}

	if params.Email != "" {
		builder = builder.Where(squirrel.Eq{"u.email": params.Email})
	}

	if params.ID != 0 {
		builder = builder.Where(squirrel.Eq{"u.id": params.ID})
	}

	var users []models.RegularUser

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}

	for rows.Next() {
		var (
			user                  models.RegularUser
			userAvatarURL         sql.NullString
			organizationAvatarURL sql.NullString
			middleName            sql.NullString
		)

		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Phone,
			&user.PasswordHash,
			&user.TOTPSalt,
			&user.FirstName,
			&user.LastName,
			&middleName,
			&userAvatarURL,
			&user.EmailVerified,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Organization.ID,
			&user.Organization.BrandName,
			&user.Organization.FullName,
			&user.Organization.ShortName,
			&user.Organization.IsContractor,
			&user.Organization.IsBanned,
			&user.Organization.INN,
			&user.Organization.OKPO,
			&user.Organization.OGRN,
			&user.Organization.KPP,
			&user.Organization.TaxCode,
			&user.Organization.Address,
			&organizationAvatarURL,
			&user.Organization.Emails,
			&user.Organization.Phones,
			&user.Organization.Messengers,
			&user.Organization.CreatedAt,
			&user.Organization.UpdatedAt,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		if middleName.Valid {
			user.MiddleName = models.Optional[string]{Value: middleName.String, Set: true}
		} else {
			user.MiddleName = models.Optional[string]{Set: false}
		}

		user.AvatarURL = userAvatarURL.String
		user.Organization.AvatarURL = organizationAvatarURL.String

		users = append(users, user)
	}

	return users, nil
}

func (s *UserStore) GetWithEmployee(ctx context.Context, qe store.QueryExecutor, params store.UserGetParams) ([]models.EmployeeUser, error) {
	builder := squirrel.
		Select(
			"u.id",
			"u.email",
			"u.phone",
			"u.password_hash",
			"u.totp_salt",
			"u.first_name",
			"u.last_name",
			"u.middle_name",
			"u.avatar_url",
			"u.email_verified",
			"u.created_at",
			"u.updated_at",
			"e.role",
			"e.position",
		).
		From("users AS u").
		Join("employee AS e ON u.id = e.user_id").
		PlaceholderFormat(squirrel.Dollar)

	if params.Limit.Set {
		builder = builder.Limit(params.Limit.Value)
	}

	if params.Offset.Set {
		builder = builder.Offset(params.Offset.Value)
	}

	if params.Role.Set {
		builder = builder.Where(squirrel.Eq{"e.role": params.Role.Value})
	}

	if params.Email != "" {
		builder = builder.Where(squirrel.Eq{"u.email": params.Email})
	}

	if params.ID != 0 {
		builder = builder.Where(squirrel.Eq{"u.id": params.ID})
	}

	var users []models.EmployeeUser

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}

	for rows.Next() {
		var (
			user          models.EmployeeUser
			userAvatarURL sql.NullString
			middleName    sql.NullString
		)

		err = rows.Scan(
			&user.ID,
			&user.Email,
			&user.Phone,
			&user.PasswordHash,
			&user.TOTPSalt,
			&user.FirstName,
			&user.LastName,
			&middleName,
			&userAvatarURL,
			&user.EmailVerified,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.Role,
			&user.Position,
		)
		if err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		if middleName.Valid {
			user.MiddleName = models.Optional[string]{Value: middleName.String, Set: true}
		} else {
			user.MiddleName = models.Optional[string]{Set: false}
		}

		user.AvatarURL = userAvatarURL.String

		users = append(users, user)
	}

	return users, nil
}

func (s *UserStore) GetUserIDByOrganizationID(ctx context.Context, qe store.QueryExecutor, organizationID int) (int, error) {
	builder := squirrel.
		Select("user_id").
		From("organization_users").
		Where(squirrel.Eq{"organization_id": organizationID}).
		PlaceholderFormat(squirrel.Dollar)

	var userID int
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&userID); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return 0, errstore.ErrUserNotFound
		}
		return 0, fmt.Errorf("query row: %w", err)
	}

	return userID, nil
}
