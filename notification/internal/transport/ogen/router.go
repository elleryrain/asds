package ogen

import (
	"context"
	"net/http"

	api "gitlab.ubrato.ru/ubrato/notification/api/gen"
)

var _ api.Handler = new(Router)

type Router struct {
	Auth
	Error
	Notification
}

type Error interface {
	HandleError(ctx context.Context, w http.ResponseWriter, r *http.Request, err error)
}

type Auth interface {
	HandleBearerAuth(ctx context.Context, operationName string, t api.BearerAuth) (context.Context, error)
}

type Notification interface {
	V1NotificationsNotificationIDReadPost(ctx context.Context, params api.V1NotificationsNotificationIDReadPostParams) (api.V1NotificationsNotificationIDReadPostRes, error)
	V1NotificationsUserIDStreamGet(ctx context.Context, params api.V1NotificationsUserIDStreamGetParams) (api.V1NotificationsUserIDStreamGetRes, error) 
	V1NotificationsUserIDGet(ctx context.Context, params api.V1NotificationsUserIDGetParams) (api.V1NotificationsUserIDGetRes, error)
	V1GetUserNotificationsBySSE() http.HandlerFunc
}

type RouterParams struct {
	Error        Error
	Auth         Auth
	Notification Notification
}

func NewRouter(params RouterParams) *Router {
	return &Router{
		Auth:         params.Auth,
		Error:        params.Error,
		Notification: params.Notification,
	}
}
