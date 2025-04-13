package store

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type RespondCreateParams struct {
	TenderID       int
	OrganizationID int
	Price          int
	IsNds          bool
}

type RespondGetParams struct {
	TenderID int
	Offset   models.Optional[uint64]
	Limit    models.Optional[uint64]
}

type RespondGetCountParams struct {
	TenderID int
}

type RespondUpdateParams struct {
	IsWinner       bool
	TenderID       int
	OrganizationID int
}
