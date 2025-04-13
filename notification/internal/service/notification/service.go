package notificationSrv

import (
	"context"

	"gitlab.ubrato.ru/ubrato/notification/internal/models"
	"gitlab.ubrato.ru/ubrato/notification/internal/store"
)

type Service struct {
	psql              DBTX
	notificationStore NotificationStore
}

type DBTX interface {
	DB() store.Querier
	InTransaction(ctx context.Context, fn store.TransactionFunc) error
}

type NotificationStore interface {
	Create(ctx context.Context, psql store.Querier, params store.NotifictionCreateParams) (models.Notification, error)
	Get(ctx context.Context, psql store.Querier, params store.NotifictionGetParams) ([]models.Notification, error)
	GetByID(ctx context.Context, qe store.Querier, notificationID int) (models.Notification, error)
	Update(ctx context.Context, psql store.Querier, params store.NotifictionUpdateParams) error
}

func New(
	psql DBTX,
	notificationStore NotificationStore,
) *Service {
	return &Service{
		psql:              psql,
		notificationStore: notificationStore,
	}
}
