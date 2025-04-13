package notificationHandler

import (
	"context"

	"gitlab.ubrato.ru/ubrato/notification/internal/lib/token"
	"gitlab.ubrato.ru/ubrato/notification/internal/models"
	"gitlab.ubrato.ru/ubrato/notification/internal/service"
)

type Handler struct {
	authSvc AuthService
	svc     Service
	broker  Broker
}

type Broker interface {
	CreateUserChan(userID int) chan models.Notification
	DeleteUserChan(userID int)
}

type AuthService interface {
	ValidateAccessToken(ctx context.Context, accessToken string) (token.Claims, error)
}

type Service interface {
	Create(ctx context.Context, params service.NotifictionCreateParams) (models.Notification, error)
	Get(ctx context.Context, params service.NotifictionGetParams) ([]models.Notification, error)
	Update(ctx context.Context, params service.NotifictionUpdateParams) error
}

func New(
	svc Service,
	authSvc AuthService,
	broker Broker) *Handler {
	return &Handler{
		svc:     svc,
		authSvc: authSvc,
		broker: broker,
	}
}
