package favourite

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql              DBTX
	favouriteStore    FavouriteStore
	tenderStore       TenderStore
	organizationStore OrganizationStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type FavouriteStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.FavouriteCreateParams) (int64, error)
	Delete(ctx context.Context, qe store.QueryExecutor, favouriteID int) error
	Get(ctx context.Context, qe store.QueryExecutor, params store.FavouriteGetParams) ([]models.Favourite[models.FavouriteObject], error)
	GetByID(ctx context.Context, qe store.QueryExecutor, favouriteID int) (models.Favourite[models.FavouriteObject], error)
	Count(ctx context.Context, qe store.QueryExecutor, params store.FavouriteGetCountParams) (int, error)
}

type OrganizationStore interface {
	GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Organization, error)
	GetContractors(ctx context.Context, qe store.QueryExecutor, params store.OrganizationContractorsGetParams) ([]models.Organization, error)
}

type TenderStore interface {
	List(ctx context.Context, qe store.QueryExecutor, params store.TenderListParams) ([]models.Tender, error)
	GetByID(ctx context.Context, qe store.QueryExecutor, tenderID int) (models.Tender, error)
}

func New(
	psql DBTX,
	favouriteStore FavouriteStore,
	tenderStore TenderStore,
	organizationStore OrganizationStore,
) *Service {
	return &Service{
		psql:              psql,
		favouriteStore:    favouriteStore,
		tenderStore:       tenderStore,
		organizationStore: organizationStore,
	}
}
