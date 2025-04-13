package amoproxy

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/amo-sync/internal/gateway/amocrm/dto"
)

type CreateLink struct {
	LinkTypeFrom string
	LinkIDFrom   int
	LinkTypeTo   string
	LinkIDTo     int
}

func (s *Service) CreateLink(ctx context.Context, params CreateLink) error {
	_, err := s.amoCRMGateway.CreateLinks(ctx, dto.CreateLinksRequest{
		EntityType: params.LinkTypeFrom,
		EntityID:   params.LinkIDFrom,
		Links: []dto.CreateLinkRequest{
			{
				ToEntityType: params.LinkTypeTo,
				ToEntityID:   params.LinkIDTo,
			},
		},
	})
	if err != nil {
		return fmt.Errorf("create link: %w", err)
	}

	return nil
}
