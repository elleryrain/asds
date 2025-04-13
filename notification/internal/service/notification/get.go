package notificationSrv

import (
	"context"

	"gitlab.ubrato.ru/ubrato/notification/internal/models"
	"gitlab.ubrato.ru/ubrato/notification/internal/service"
	"gitlab.ubrato.ru/ubrato/notification/internal/store"
)

func (s *Service) Get(ctx context.Context, params service.NotifictionGetParams) ([]models.Notification, error) {
	return s.notificationStore.Get(ctx, s.psql.DB(), store.NotifictionGetParams{
		UserID: models.NewOptional(params.UserID),
	})
}
