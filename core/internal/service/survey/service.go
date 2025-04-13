package survey

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql   DBTX
	broker Broker
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type Broker interface {
	Publish(ctx context.Context, subject broker.Topic, data []byte) error
}

func New(
	psql DBTX,
	broker Broker,
) *Service {
	return &Service{
		psql:   psql,
		broker: broker,
	}
}
