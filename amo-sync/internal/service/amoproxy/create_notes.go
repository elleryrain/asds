package amoproxy

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/amo-sync/internal/gateway/amocrm/dto"
)

type CreateNotesParams struct {
	LeadID   int
	Messages []string
}

func (s *Service) CreateNotes(ctx context.Context, params CreateNotesParams) error {
	if len(params.Messages) == 0 {
		return nil
	}

	request := dto.CreateNotesRequest{
		EntityType: "leads",
	}

	for _, message := range params.Messages {
		request.Notes = append(request.Notes, dto.CreateNoteRequest{
			EntityID: params.LeadID,
			NoteType: "common",
			Params: dto.CreateNoteRequestParams{
				Text: message,
			},
		})
	}

	err := s.amoCRMGateway.CreateNotes(ctx, request)
	if err != nil {
		return fmt.Errorf("create notes: %w", err)
	}

	return nil
}
