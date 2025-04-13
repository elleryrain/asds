package respond

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql         DBTX
	respondStore RespondStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type RespondStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.RespondCreateParams) error
	Get(ctx context.Context, qe store.QueryExecutor, params store.RespondGetParams) ([]models.Respond, error)
	Count(ctx context.Context, qe store.QueryExecutor, params store.RespondGetCountParams) (int, error)
}

func New(
	psql DBTX,
	respondStore RespondStore,
) *Service {
	return &Service{
		psql:         psql,
		respondStore: respondStore,
	}
}
