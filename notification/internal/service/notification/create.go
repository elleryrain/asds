package notificationSrv

import (
	"context"

	"gitlab.ubrato.ru/ubrato/notification/internal/models"
	"gitlab.ubrato.ru/ubrato/notification/internal/service"
	"gitlab.ubrato.ru/ubrato/notification/internal/store"
)

func (s *Service) Create(ctx context.Context, params service.NotifictionCreateParams) (models.Notification, error) {
	return s.notificationStore.Create(ctx, s.psql.DB(), store.NotifictionCreateParams{
		UserID:       params.UserID,
		Title:        params.Title,
		Comment:      params.Comment,
		ActionButton: params.ActionButton,
		StatusBlock:  params.StatusBlock,
	})
}
