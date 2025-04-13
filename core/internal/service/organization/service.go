package organization

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql              DBTX
	organizationStore OrganizationStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type OrganizationStore interface {
	Get(ctx context.Context, qe store.QueryExecutor, params store.OrganizationGetParams) ([]models.Organization, error)
	GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Organization, error)
	GetCustomer(ctx context.Context, qe store.QueryExecutor, organizationID int) (models.Organization, error)
	GetContractorByID(ctx context.Context, qe store.QueryExecutor, organizationID int) (models.Organization, error)
	GetContractors(ctx context.Context, qe store.QueryExecutor, params store.OrganizationContractorsGetParams) ([]models.Organization, error)
	Update(ctx context.Context, qe store.QueryExecutor, params store.OrganizationUpdateParams) error
	Count(ctx context.Context, qe store.QueryExecutor, params store.OrganizationGetCountParams) (int, error)
}

func New(psql DBTX, organizationStore OrganizationStore) *Service {
	return &Service{
		psql:              psql,
		organizationStore: organizationStore,
	}
}
