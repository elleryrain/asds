package suggest

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/gateway/dadata"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql          DBTX
	dadataGateway DadataGateway
	catalogStore  CatalogStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type DadataGateway interface {
	FindByINN(ctx context.Context, INN string) (dadata.FindByInnResponse, error)
}

type CatalogStore interface {
	ListCities(ctx context.Context, qe store.QueryExecutor, name string) ([]models.City, error)
}

func New(
	psql DBTX,
	dadataGateway DadataGateway,
	catalogStore CatalogStore,
) *Service {
	return &Service{
		psql:          psql,
		dadataGateway: dadataGateway,
		catalogStore:  catalogStore,
	}
}
