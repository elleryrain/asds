package amocrm

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/amo-sync/internal/gateway/amocrm/dto"
)

func (c *Client) CreateLinks(ctx context.Context, request dto.CreateLinksRequest) (dto.BasicResponse[dto.CreateLinksResponse], error) {
	var response dto.BasicResponse[dto.CreateLinksResponse]

	return response, c.do(ctx, fmt.Sprintf(endpointLink, request.EntityType, request.EntityID), request.Links, &response)
}
