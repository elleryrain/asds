package user

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *UserStore) Create(ctx context.Context, qe store.QueryExecutor, params store.UserCreateParams) (models.User, error) {
	builder := squirrel.
		Insert("users").
		Columns(
			"email",
			"phone",
			"password_hash",
			"totp_salt",
			"first_name",
			"last_name",
			"middle_name",
			"avatar_url",
			"email_verified",
			"is_banned",
		).
		Values(
			params.Email,
			params.Phone,
			params.PasswordHash,
			params.TOTPSalt,
			params.FirstName,
			params.LastName,
			sql.NullString{Valid: params.MiddleName.Set, String: params.MiddleName.Value},
			sql.NullString{Valid: params.AvatarURL != "", String: params.AvatarURL},
			params.EmailVerified,
			false,
		).
		Suffix(`
			RETURNING
				id,
				email,
				phone,
				password_hash,
				totp_salt,
				first_name,
				last_name,
				middle_name,
				avatar_url,
				email_verified,
				created_at,
				updated_at
		`).
		PlaceholderFormat(squirrel.Dollar)

	var (
		createdUser models.User
		avatarURL   sql.NullString
		middleName  sql.NullString
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdUser.ID,
		&createdUser.Email,
		&createdUser.Phone,
		&createdUser.PasswordHash,
		&createdUser.TOTPSalt,
		&createdUser.FirstName,
		&createdUser.LastName,
		&middleName,
		&avatarURL,
		&createdUser.EmailVerified,
		&createdUser.CreatedAt,
		&createdUser.UpdatedAt,
	)
	if err != nil {
		return models.User{}, fmt.Errorf("query row: %w", err)
	}

	if middleName.Valid {
		createdUser.MiddleName = models.Optional[string]{Value: middleName.String, Set: true}
	}

	createdUser.AvatarURL = avatarURL.String

	return createdUser, nil
}

func (s *UserStore) CreateEmployee(ctx context.Context, qe store.QueryExecutor, params store.UserCreateEmployeeParams) error {
	builder := squirrel.
		Insert("employee").
		Columns(
			"user_id",
			"role",
			"position",
		).
		Values(
			params.UserID,
			params.Role,
			params.Postition,
		).
		PlaceholderFormat(squirrel.Dollar)

	_, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("query row: %w", err)
	}

	return nil
}
