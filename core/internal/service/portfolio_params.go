package service

import "gitlab.ubrato.ru/ubrato/core/internal/models"

type PortfolioCreateParams struct {
	OrganizationID int
	Title          string
	Description    string
	Attachments    []string
}

type PortfolioGetParams struct {
	OrganizationID int
}

type PortfolioUpdateParams struct {
	PortfolioID int
	Title       models.Optional[string]
	Description models.Optional[string]
	Attachments models.Optional[[]string]
}
