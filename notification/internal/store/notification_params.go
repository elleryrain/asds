package store

import "gitlab.ubrato.ru/ubrato/notification/internal/models"

type NotifictionCreateParams struct {
	UserID       int
	Title        string
	Comment      models.Optional[string]
	ActionButton models.Optional[models.ActionButton]
	StatusBlock  models.Optional[models.StatusBlock]
}

type NotifictionGetParams struct {
	UserID         models.Optional[int]
	NotificationID models.Optional[int]
}

type NotifictionUpdateParams struct {
	NotificationID int
}
