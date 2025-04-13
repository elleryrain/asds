package amocrm

import (
	"context"

	"gitlab.ubrato.ru/ubrato/amo-sync/internal/gateway/amocrm/dto"
)

func (c *Client) CreateContacts(ctx context.Context, request dto.CreateContactsRequest) (dto.BasicResponse[dto.CreateContactsResponse], error) {
	var response dto.BasicResponse[dto.CreateContactsResponse]

	return response, c.do(ctx, endpointContacts, request, &response)
}
