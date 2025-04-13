package service

import "gitlab.ubrato.ru/ubrato/notification/internal/models"

type NotifictionCreateParams struct {
	UserID       int
	Title        string
	Comment      models.Optional[string]
	ActionButton models.Optional[models.ActionButton]
	StatusBlock  models.Optional[models.StatusBlock]
}

type NotifictionGetParams struct {
	UserID     int
	OnlyUnread bool
}

type NotifictionUpdateParams struct {
	NotificationID int
}
