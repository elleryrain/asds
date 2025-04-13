package amocrm

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/amo-sync/internal/gateway/amocrm/dto"
)

func (c *Client) CreateNotes(ctx context.Context, request dto.CreateNotesRequest) error {
	return c.do(ctx, fmt.Sprintf(endpointNotes, request.EntityType), request.Notes, nil)
}
