package notificationHandler

import (
	"context"
	"errors"

	api "gitlab.ubrato.ru/ubrato/notification/api/gen"
	"gitlab.ubrato.ru/ubrato/notification/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/notification/internal/service"
	"gitlab.ubrato.ru/ubrato/notification/internal/store/errstore"
)

func (s *Handler) V1NotificationsNotificationIDReadPost(ctx context.Context, params api.V1NotificationsNotificationIDReadPostParams) (api.V1NotificationsNotificationIDReadPostRes, error) {
	if err := s.svc.Update(ctx, service.NotifictionUpdateParams{
		NotificationID: params.NotificationID,
	}); err != nil {
		if errors.Is(err, errstore.ErrNotificationNotFound) {
			return nil, cerr.Wrap(err, cerr.CodeNotFound, "Уведомление не найдено", nil)
		}

		return nil, err
	}

	return &api.V1NotificationsNotificationIDReadPostOK{}, nil
}
