package jetstream

import (
	"context"
	"fmt"

	"github.com/rs/zerolog/log"
	"gitlab.ubrato.ru/ubrato/notification/internal/models"
	"gitlab.ubrato.ru/ubrato/notification/internal/service"
)

func (b *Broker) SaveNotification(ctx context.Context, notification models.Notification) error {
	newNotification, err := b.notificationService.Create(ctx, service.NotifictionCreateParams{
		UserID:       notification.UserID,
		Title:        notification.Title,
		Comment:      models.Optional[string]{Value: notification.Comment, Set: notification.Comment != ""},
		ActionButton: models.NewOptActionButton(notification.ActionButton),
		StatusBlock:  models.NewOptStatusBlock(notification.StatusBlock)})
	if err != nil {
		return fmt.Errorf("create notification: %w", err)
	}

	log.Debug().Interface("notification", newNotification).Msg("Notification stored")

	// Если пользователь подключен по SSE, то ему отправятся уведомления
	if userChan, ok := b.UserNotifications[notification.UserID]; ok {
		log.Debug().Int("userID", newNotification.UserID).Msg("Sending notification to user channel")
		userChan <- newNotification
	}

	return nil
}
