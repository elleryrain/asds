package store

import (
	"time"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
)

type PortfolioCreateParams struct {
	OrganizationID int
	Title          string
	Description    string
	Attachments    []string
	CreatedAt      time.Time
}

type PortfolioGetParams struct {
	PortfolioID    models.Optional[int]
	OrganizationID models.Optional[int]
}

type PortfolioUpdateParams struct {
	PortfolioID int
	Title       models.Optional[string]
	Description models.Optional[string]
	Attachments models.Optional[[]string]
}
