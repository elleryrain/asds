package questionanswer

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

type Service struct {
	psql                DBTX
	broker              Broker
	questionAnswerStore QuestionAnswerStore
	tenderStore         TenderStore
	verificationStore   VerificationStore
}

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type Broker interface {
	Publish(ctx context.Context, subject broker.Topic, data []byte) error
}

type QuestionAnswerStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.CreateQuestionAnswerParams) (models.QuestionAnswer, error)
	GetWithAccess(ctx context.Context, qe store.QueryExecutor, params store.QuestionAnswerGetWithAccessParams) ([]models.QuestionWithAnswer, error)
	GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.QuestionWithAnswer, error)
}

type TenderStore interface {
	GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Tender, error)
	GetTenderNotifyInfoByObjectID(ctx context.Context, qe store.QueryExecutor, params store.TenderNotifyInfoParams) (models.Tender, error)
}

type VerificationStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestCreateParams) error
}

func New(
	psql DBTX,
	questionAnswerStore QuestionAnswerStore,
	tenderStore TenderStore,
	verificationStore VerificationStore,
	broker Broker,
) *Service {
	return &Service{
		psql:                psql,
		broker:              broker,
		questionAnswerStore: questionAnswerStore,
		tenderStore:         tenderStore,
		verificationStore:   verificationStore,
	}
}
