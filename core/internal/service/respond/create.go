package respond

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Create(ctx context.Context, params service.RespondCreateParams) error {
	return s.respondStore.Create(ctx, s.psql.DB(), store.RespondCreateParams{
		TenderID:       params.TenderID,
		OrganizationID: params.OrganizationID,
		Price:          params.Price,
		IsNds:          params.IsNds,
	})
}
