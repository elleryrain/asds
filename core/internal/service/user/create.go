package user

import (
	"context"
	"errors"
	"fmt"

	"github.com/google/uuid"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/crypto"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) CreateEmployee(ctx context.Context, params service.UserCreateEmployeeParams) (models.EmployeeUser, error) {
	var user models.EmployeeUser

	err := s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		hashedPassword, err := crypto.Password(params.Password)
		if err != nil {
			return fmt.Errorf("hash password: %w", err)
		}

		createdUser, err := s.userStore.Create(ctx, qe, store.UserCreateParams{
			Email:         params.Email,
			Phone:         params.Phone,
			PasswordHash:  hashedPassword,
			TOTPSalt:      uuid.New().String(),
			FirstName:     params.FirstName,
			LastName:      params.LastName,
			MiddleName:    params.MiddleName,
			EmailVerified: true,
		})
		if err != nil {
			return fmt.Errorf("create user: %w", err)
		}

		err = s.userStore.CreateEmployee(ctx, qe, store.UserCreateEmployeeParams{
			UserID:    createdUser.ID,
			Role:      params.Role,
			Postition: params.Position,
		})
		if err != nil {
			return fmt.Errorf("create employee: %w", err)
		}

		fullUsers, err := s.userStore.Get(ctx, qe, store.UserGetParams{ID: createdUser.ID})
		if err != nil {
			return fmt.Errorf("find created employee: %w", err)
		}

		if len(fullUsers) == 0 {
			return errors.New("get created user")
		}

		fullUser := fullUsers[0]

		user = models.EmployeeUser{User: fullUser.User, Role: fullUser.Role, Position: fullUser.Position}

		return nil
	})
	if err != nil {
		return models.EmployeeUser{}, fmt.Errorf("create employee transaction: %w", err)
	}

	return user, nil
}
