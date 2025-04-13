package user

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql      DBTX
	userStore UserStore
	broker    Broker
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type UserStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.UserCreateParams) (models.User, error)
	CreateEmployee(ctx context.Context, qe store.QueryExecutor, params store.UserCreateEmployeeParams) error
	Get(ctx context.Context, qe store.QueryExecutor, params store.UserGetParams) ([]models.FullUser, error)
	GetWithOrganiztion(ctx context.Context, qe store.QueryExecutor, params store.UserGetParams) ([]models.RegularUser, error)
	GetWithEmployee(ctx context.Context, qe store.QueryExecutor, params store.UserGetParams) ([]models.EmployeeUser, error)
	Update(ctx context.Context, qe store.QueryExecutor, params store.UserUpdateParams) error
	CountUsers(ctx context.Context, qe store.QueryExecutor) (int, error)
	CountEmployee(ctx context.Context, qe store.QueryExecutor, role models.Optional[[]models.UserRole]) (int, error)

	SetEmailVerified(ctx context.Context, qe store.QueryExecutor, userID int) error
	ResetPassword(ctx context.Context, qe store.QueryExecutor, params store.ResetPasswordParams) error
}

type Broker interface {
	Publish(ctx context.Context, subject broker.Topic, data []byte) error
}

func New(
	psql DBTX,
	userStore UserStore,
	broker Broker,
) *Service {
	return &Service{
		psql:      psql,
		userStore: userStore,
		broker:    broker,
	}
}
