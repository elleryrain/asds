package handler

import (
	"context"

	eventsv1 "gitlab.ubrato.ru/ubrato/amo-sync/internal/gen/pb/events/v1"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/models"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/service/amoproxy"
)

func (h *Handler) UserRegistered(ctx context.Context, event *eventsv1.UserRegistered) error {
	amoContact, amoCompany := ConvertProtoUserToAmoModels(event.GetUser())

	contactID, err := h.amoProxySvc.CreateContact(ctx, amoproxy.CreateContactParams{
		Contact:         amoContact,
		ResponsibleUser: models.AmoUserEkaterinaFedorova,
		ExternalID:      int(event.GetUser().GetId()),
	})
	if err != nil {
		return err
	}

	companyID, err := h.amoProxySvc.CreateCompany(ctx, amoproxy.CreateCompanyParams{
		Company:         amoCompany,
		ResponsibleUser: models.AmoUserEkaterinaFedorova,
		ExternalID:      int(event.GetUser().GetOrganization().GetId()),
	})
	if err != nil {
		return err
	}

	err = h.amoProxySvc.CreateLink(ctx, amoproxy.CreateLink{
		LinkTypeFrom: "contacts",
		LinkIDFrom:   contactID,
		LinkTypeTo:   "companies",
		LinkIDTo:     companyID,
	})
	if err != nil {
		return err
	}

	_, err = h.amoProxySvc.CreateLead(ctx, amoproxy.CreateLeadParams{
		Name:            event.GetUser().GetOrganization().GetShortName(),
		ResponsibleUser: models.AmoUserEkaterinaFedorova,
		Pipeline:        models.AmoLeadPipelineIncomingRequests,
		Status:          models.AmoLeadIncomingRequestsPipelineStatusRegistrationCompleted,
		ContactID:       contactID,
		CompanyID:       companyID,
		ExternalID:      0,
	})
	if err != nil {
		return err
	}

	return nil
}
