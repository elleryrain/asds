package models

import (
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
)

type RespondPagination struct {
	Responds   []Respond
	Pagination pagination.Pagination
}

type Respond struct {
	ID             int
	TenderID       int
	OrganizationID int
	Price          int
	IsNDSPrice     bool
	IsWinner       bool
	CreatedAt      time.Time
}

func ConvertRespondModelToApi(respond Respond) api.Respond {
	return api.Respond{
		ID:             respond.ID,
		TenderID:       respond.TenderID,
		OrganizationID: respond.OrganizationID,
		Price:          respond.Price,
		IsNdsPrice:     respond.IsNDSPrice,
		IsWinner:       respond.IsWinner,
		CreatedAt:      respond.CreatedAt,
	}
}
