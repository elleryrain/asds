package tender

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type (
	TenderStore struct {
		catalogStore CatalogStore
	}

	CatalogStore interface {
		GetObjectsByIDs(ctx context.Context, qe store.QueryExecutor, objectIDs []int) (map[int]models.Object, error)
		GetServicesByIDs(ctx context.Context, qe store.QueryExecutor, serviceIDs []int) (map[int]models.Service, error)
	}
)

func NewTenderStore(catalogStore CatalogStore) *TenderStore {
	return &TenderStore{catalogStore: catalogStore}
}
