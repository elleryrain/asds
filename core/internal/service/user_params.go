package service

import (
	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

type UserCreateEmployeeParams struct {
	Email      string
	Phone      string
	Password   string
	FirstName  string
	LastName   string
	MiddleName models.Optional[string]
	Role       models.UserRole
	Position   string
}

type UserUpdateParams struct {
	UserID     int
	Phone      models.Optional[string]
	FirstName  models.Optional[string]
	LastName   models.Optional[string]
	MiddleName models.Optional[string]
	AvatarURL  models.Optional[string]
}

type UserGetParams struct {
	Email   string
	ID      int
	Role    models.Optional[[]models.UserRole]
	Page    uint64
	PerPage uint64
}
