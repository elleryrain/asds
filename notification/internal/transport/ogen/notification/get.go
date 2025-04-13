package notificationHandler

import (
	"context"

	api "gitlab.ubrato.ru/ubrato/notification/api/gen"
	"gitlab.ubrato.ru/ubrato/notification/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/notification/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/notification/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/notification/internal/models"
	"gitlab.ubrato.ru/ubrato/notification/internal/service"
)

func (s *Handler) V1NotificationsUserIDGet(ctx context.Context, params api.V1NotificationsUserIDGetParams) (api.V1NotificationsUserIDGetRes, error) {
	if contextor.GetUserID(ctx) != params.UserID {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "Нет доступа для получения уведомлений", nil)
	}

	notifications, err := s.svc.Get(ctx, service.NotifictionGetParams{UserID: params.UserID})
	if err != nil {
		return nil, err
	}

	return &api.V1NotificationsUserIDGetOK{
		Data: convert.Slice[[]models.Notification, []api.Notification](notifications, models.ConvertNotificationToAPI),
	}, nil
}
