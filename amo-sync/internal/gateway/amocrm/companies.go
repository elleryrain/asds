package amocrm

import (
	"context"

	"gitlab.ubrato.ru/ubrato/amo-sync/internal/gateway/amocrm/dto"
)

func (c *Client) CreateCompanies(ctx context.Context, request dto.CreateCompaniesRequest) (dto.BasicResponse[dto.CreateCompaniesResponse], error) {
	var response dto.BasicResponse[dto.CreateCompaniesResponse]

	return response, c.do(ctx, endpointCompanies, request, &response)
}
