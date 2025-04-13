package store

import (
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

type UserCreateParams struct {
	Email         string
	Phone         string
	PasswordHash  string
	TOTPSalt      string
	FirstName     string
	LastName      string
	MiddleName    models.Optional[string]
	EmailVerified bool
	AvatarURL     string
}

type UserUpdateParams struct {
	UserID     int
	Phone      models.Optional[string]
	FirstName  models.Optional[string]
	LastName   models.Optional[string]
	MiddleName models.Optional[string]
	AvatarURL  models.Optional[string]
}

type UserCreateEmployeeParams struct {
	UserID    int
	Role      models.UserRole
	Postition string
}

type UserGetParams struct {
	Email  string
	ID     int
	Role   models.Optional[[]models.UserRole]
	Offset models.Optional[uint64]
	Limit  models.Optional[uint64]
}

type ResetPasswordParams struct {
	UserID       int
	PasswordHash string
}
