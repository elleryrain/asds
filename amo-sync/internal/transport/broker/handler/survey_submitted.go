package handler

import (
	"context"
	"fmt"

	eventsv1 "gitlab.ubrato.ru/ubrato/amo-sync/internal/gen/pb/events/v1"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/models"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/service/amoproxy"
)

func (h *Handler) SurveySubmitted(ctx context.Context, event *eventsv1.SurveySubmitted) error {
	contactID, err := h.amoProxySvc.CreateContact(ctx, amoproxy.CreateContactParams{
		Contact: models.AmoContact{
			FirstName: event.Survey.GetName(),
			Phone:     event.Survey.GetPhone(),
		},
		ResponsibleUser: models.AmoUserEkaterinaFedorova,
		ExternalID:      0,
	})
	if err != nil {
		return fmt.Errorf("create contact: %w", err)
	}

	leadID, err := h.amoProxySvc.CreateLead(ctx, amoproxy.CreateLeadParams{
		Name:            event.Survey.GetName(),
		ResponsibleUser: models.AmoUserEkaterinaFedorova,
		Pipeline:        models.AmoLeadPipelineIncomingRequests,
		Status:          models.AmoLeadIncomingRequestsPipelineStatusNewLead,
		ContactID:       contactID,
		Tags:            []string{event.GetSurvey().GetType().String()},
		ExternalID:      0,
	})
	if err != nil {
		return fmt.Errorf("create lead: %w", err)
	}

	err = h.amoProxySvc.CreateNotes(ctx, amoproxy.CreateNotesParams{
		LeadID: leadID,
		Messages: []string{
			"Тип опроса: " + event.GetSurvey().GetType().String(),
			event.GetSurvey().GetQuestion(),
		},
	})
	if err != nil {
		return fmt.Errorf("create notes: %w", err)
	}

	return nil
}
