package models

import (
	"net/url"
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
)

type Portfolio struct {
	ID             int
	OrganizationID int
	Title          string
	Description    string
	Attachments    []string
	CreatedAt      time.Time
	UpdatedAt      Optional[time.Time]
}

func ConvertPortfolioModelToApi(portfolio Portfolio) api.Portfolio {
	return api.Portfolio{
		ID:          portfolio.ID,
		Title:       portfolio.Title,
		Description: api.Description(portfolio.Description),
		Attachments: convert.Slice[[]string, []url.URL](
			portfolio.Attachments, func(s string) url.URL { return stringToUrl(s) },
		),
		CreatedAt: portfolio.CreatedAt,
		UpdatedAt: api.OptDateTime{Value: portfolio.UpdatedAt.Value, Set: portfolio.UpdatedAt.Set},
	}
}
