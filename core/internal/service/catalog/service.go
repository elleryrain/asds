package catalog

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql         DBTX
	catalogStore CatalogStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type CatalogStore interface {
	GetServices(ctx context.Context, qe store.QueryExecutor) (models.Services, error)
	GetObjects(ctx context.Context, qe store.QueryExecutor) (models.Objects, error)
	CreateRegion(ctx context.Context, qe store.QueryExecutor, params store.CatalogCreateRegionParams) (models.Region, error)
	CreateCity(ctx context.Context, qe store.QueryExecutor, params store.CatalogCreateCityParams) (models.City, error)
	GetRegionByID(ctx context.Context, qe store.QueryExecutor, regionID int) (models.Region, error)
	GetCityByID(ctx context.Context, qe store.QueryExecutor, cityID int) (models.City, error)
	CreateObject(ctx context.Context, qe store.QueryExecutor, params store.CatalogCreateObjectParams) (models.Object, error)
	CreateService(ctx context.Context, qe store.QueryExecutor, params store.CatalogCreateServiceParams) (models.Service, error)
	GetMeasurements(ctx context.Context, qe store.QueryExecutor) ([]models.Measure, error)
}

func New(
	psql DBTX,
	catalogStore CatalogStore,
) *Service {
	return &Service{
		psql:         psql,
		catalogStore: catalogStore,
	}
}
