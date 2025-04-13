package store

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type WinnersCreateParams struct {
	TenderID       int
	OrganizationID int
}

type WinnerUpdateParams struct {
	WinnerID int
	Accepted models.AcceptedStatus
}
