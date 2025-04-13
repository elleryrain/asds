package tender

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql              DBTX
	broker            Broker
	tenderStore       TenderStore
	additionStore     AdditionStore
	verificationStore VerificationStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type Broker interface {
	Publish(ctx context.Context, subject broker.Topic, data []byte) error
}

type TenderStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.TenderCreateParams) (int, error)
	GetByID(ctx context.Context, qe store.QueryExecutor, tenderID int) (models.Tender, error)
	List(ctx context.Context, qe store.QueryExecutor, params store.TenderListParams) ([]models.Tender, error)
	Update(ctx context.Context, qe store.QueryExecutor, params store.TenderUpdateParams) (int, error)
	Count(ctx context.Context, qe store.QueryExecutor, params store.TenderGetCountParams) (int, error)
}

type AdditionStore interface {
	CreateAddition(ctx context.Context, qe store.QueryExecutor, params store.AdditionCreateParams) (int, error)
	Get(ctx context.Context, qe store.QueryExecutor, params store.AdditionGetParams) ([]models.Addition, error)
}

type VerificationStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestCreateParams) error
}

func New(
	psql DBTX,
	tenderStore TenderStore,
	additionStore AdditionStore,
	verificationStore VerificationStore,
	broker Broker,
) *Service {
	return &Service{
		psql:              psql,
		tenderStore:       tenderStore,
		additionStore:     additionStore,
		verificationStore: verificationStore,
		broker:            broker,
	}
}
