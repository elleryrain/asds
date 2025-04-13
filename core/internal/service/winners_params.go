package service

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type WinnersCreateParams struct {
	OrganizationID int
	TenderID       int
}

type WinnerUpdateParams struct {
	WinnerID int
	Accepted models.AcceptedStatus
}
