package questionnaire

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql               DBTX
	questionnaireStore QuestionnaireStore
	organizationStore  OrganizationStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type QuestionnaireStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.QuestionnaireCreateParams) error
	Get(ctx context.Context, qe store.QueryExecutor, params store.QuestionnaireGetParams) ([]models.Questionnaire, error)
	GetStatus(ctx context.Context, qe store.QueryExecutor, organizationID int) (bool, error)
}

type OrganizationStore interface {
	GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Organization, error)
}

func New(
	psql DBTX,
	questionnaireStore QuestionnaireStore,
	organizationStore OrganizationStore,
) *Service {
	return &Service{
		psql:               psql,
		questionnaireStore: questionnaireStore,
		organizationStore:  organizationStore,
	}
}
