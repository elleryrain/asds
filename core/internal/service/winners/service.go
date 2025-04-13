package winners

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql         DBTX
	broker       Broker
	winnersStore WinnersStore
	tenderStore  TenderStore
	respondStore RespondStore
	userStore    UserStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type Broker interface {
	Publish(ctx context.Context, subject broker.Topic, data []byte) error
}

type WinnersStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.WinnersCreateParams) (models.Winners, error)
	Get(ctx context.Context, qe store.QueryExecutor, tenderID int) ([]models.Winners, error)
	UpdateStatus(ctx context.Context, qe store.QueryExecutor, params store.WinnerUpdateParams) error
	GetOrganizationIDByWinnerID(ctx context.Context, qe store.QueryExecutor, winnerID int) (int, error)
	GetTenderIDByWinnerID(ctx context.Context, qe store.QueryExecutor, winnerID int) (int, error)
	Count(ctx context.Context, qe store.QueryExecutor, tenderID int) (int, error)
}

type TenderStore interface {
	GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Tender, error)
	UpdateStatus(ctx context.Context, qe store.QueryExecutor, params store.TenderUpdateStatusParams) error
}

type RespondStore interface {
	Update(ctx context.Context, qe store.QueryExecutor, params store.RespondUpdateParams) error
}

type UserStore interface {
	GetUserIDByOrganizationID(ctx context.Context, qe store.QueryExecutor, organizationID int) (int, error)
}

func New(
	psql DBTX,
	broker Broker,
	winnersStore WinnersStore,
	tenderStore TenderStore,
	respondStore RespondStore,
	userStore UserStore,
) *Service {
	return &Service{
		psql:         psql,
		broker:       broker,
		winnersStore: winnersStore,
		tenderStore:  tenderStore,
		respondStore: respondStore,
		userStore:    userStore,
	}
}
