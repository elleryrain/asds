package amocrm

import (
	"context"

	"gitlab.ubrato.ru/ubrato/amo-sync/internal/gateway/amocrm/dto"
)

func (c *Client) CreateLeads(ctx context.Context, request dto.CreateLeadsRequest) (dto.BasicResponse[dto.CreateLeadsResponse], error) {
	var response dto.BasicResponse[dto.CreateLeadsResponse]

	return response, c.do(ctx, endpointLeads, request, &response)
}
