package notificationSrv

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/notification/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/notification/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/notification/internal/service"
	"gitlab.ubrato.ru/ubrato/notification/internal/store"
)

func (s *Service) Update(ctx context.Context, params service.NotifictionUpdateParams) error {
	notification, err := s.notificationStore.GetByID(ctx, s.psql.DB(), params.NotificationID)
	if err != nil {
		return fmt.Errorf("get notification by id: %w", err)
	}

	if contextor.GetUserID(ctx) != notification.UserID {
		return cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to read notification", nil)
	}
	
	return s.notificationStore.Update(ctx, s.psql.DB(), store.NotifictionUpdateParams{
		NotificationID: params.NotificationID,
	})
}
