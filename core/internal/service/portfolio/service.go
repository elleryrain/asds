package portfolio

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql           DBTX
	portfolioStore PortfolioStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type PortfolioStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.PortfolioCreateParams) (models.Portfolio, error)
	Count(ctx context.Context, qe store.QueryExecutor, organizationID int) (int, error)
	Delete(ctx context.Context, qe store.QueryExecutor, id int) error
	Get(ctx context.Context, qe store.QueryExecutor, params store.PortfolioGetParams) ([]models.Portfolio, error)
	GetByID(ctx context.Context, qe store.QueryExecutor, portfolioID int) (models.Portfolio, error)
	Update(ctx context.Context, qe store.QueryExecutor, params store.PortfolioUpdateParams) (models.Portfolio, error)
}

func New(
	psql DBTX,
	portfolioStore PortfolioStore,
) *Service {
	return &Service{
		psql:           psql,
		portfolioStore: portfolioStore,
	}
}
