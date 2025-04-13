package user

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) GetByID(ctx context.Context, userID int) (models.RegularUser, error) {
	users, err := s.userStore.Get(ctx, s.psql.DB(), store.UserGetParams{ID: userID})
	if err != nil {
		return models.RegularUser{}, err
	}

	if len(users) == 0 {
		return models.RegularUser{}, cerr.Wrap(
			fmt.Errorf("user not found"),
			cerr.CodeNotFound,
			fmt.Sprintf("user with %d id not found", userID),
			nil,
		)
	}

	return models.RegularUser{User: users[0].User, Organization: users[0].Organization}, nil
}

func (s *Service) Get(ctx context.Context, params service.UserGetParams) (models.UserPagination, error) {
	var (
		usersFull []api.V1UsersGetOKDataItem
		users     []models.FullUser
		roles     []models.UserRole
		count     int
	)

	countUser, err := s.userStore.CountUsers(ctx, s.psql.DB())
	if err != nil {
		return models.UserPagination{}, fmt.Errorf("get count users: %w", err)
	}

	countEmployee, err := s.userStore.CountEmployee(ctx, s.psql.DB(), params.Role)
	if err != nil {
		return models.UserPagination{}, fmt.Errorf("get count employee: %w", err)
	}

	if params.Role.Set {
		for _, r := range params.Role.Value {
			if r < models.UserRoleEmployee {
				roles = append(roles, r)
			}
		}
	}

	switch {
	case !params.Role.Set:
		user, err := s.userStore.Get(ctx, s.psql.DB(), store.UserGetParams{
			Offset: models.Optional[uint64]{Value: params.Page * params.PerPage, Set: (params.Page * params.PerPage) != 0},
			Limit:  models.NewOptional(params.PerPage)})
		if err != nil {
			return models.UserPagination{}, fmt.Errorf("get all users: %w", err)
		}

		count = countUser + countEmployee
		users = user

	case len(roles) != 0:
		user, err := s.userStore.GetWithOrganiztion(ctx, s.psql.DB(), store.UserGetParams{
			Offset: models.Optional[uint64]{Value: params.Page * params.PerPage, Set: (params.Page * params.PerPage) != 0},
			Limit:  models.NewOptional(params.PerPage)})
		if err != nil {
			return models.UserPagination{}, fmt.Errorf("get rgular users: %w", err)
		}

		count = countUser
		users = models.ConvertRegularToFull(user)

	default:
		user, err := s.userStore.GetWithEmployee(ctx, s.psql.DB(), store.UserGetParams{
			Role:   models.NewOptional(params.Role.Value),
			Offset: models.Optional[uint64]{Value: params.Page * params.PerPage, Set: (params.Page * params.PerPage) != 0},
			Limit:  models.NewOptional(params.PerPage)})
		if err != nil {
			return models.UserPagination{}, fmt.Errorf("get employee users: %w", err)
		}

		count = countEmployee
		users = models.ConvertEmployeeToFull(user)
	}

	for _, userFull := range users {
		var user api.V1UsersGetOKDataItem

		if userFull.Role < models.UserRoleEmployee {
			userFull.RegularUser.User = userFull.User
			user.Type = api.RegularUserV1UsersGetOKDataItem
			user.RegularUser = models.ConvertRegularUserModelToApi(userFull.RegularUser)
		} else if userFull.Role >= models.UserRoleEmployee {
			userFull.EmployeeUser.User = userFull.User
			user.Type = api.EmployeeUserV1UsersGetOKDataItem
			user.EmployeeUser = models.ConvertEmployeeUserModelToApi(userFull.EmployeeUser)
		}

		usersFull = append(usersFull, user)
	}

	return models.UserPagination{
		Users:      usersFull,
		Pagination: pagination.New(params.Page, params.PerPage, uint64(count)),
	}, nil
}