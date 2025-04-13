package amoproxy

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/amo-sync/internal/gateway/amocrm/dto"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/models"
)

type CreateContactParams struct {
	Contact         models.AmoContact
	ResponsibleUser models.AmoUser
	ExternalID      int
}

func (s *Service) CreateContact(ctx context.Context, params CreateContactParams) (int, error) {
	id, err := s.createContact(ctx, params)
	if err != nil {
		return 0, fmt.Errorf("create contact: %w", err)
	}

	err = s.amoStore.Create(ctx, s.psql.DB(), params.ExternalID, id, models.EntityContact)
	if err != nil {
		return 0, fmt.Errorf("create external id amo id contact link: %w", err)
	}

	return id, nil
}

func (s *Service) createContact(ctx context.Context, params CreateContactParams) (int, error) {
	contact := params.Contact

	resp, err := s.amoCRMGateway.CreateContacts(ctx, dto.CreateContactsRequest{
		{
			Name:              params.Contact.FirstName + " " + params.Contact.LastName,
			ResponsibleUserID: int(params.ResponsibleUser),
			CustomFieldsValues: []dto.CustomFieldsValue{
				{
					FieldID: ContactFieldIDFirstName,
					Values: []dto.Value{
						{
							Value: contact.FirstName,
						},
					},
				},
				{
					FieldID: ContactFieldIDLastName,
					Values: []dto.Value{
						{
							Value: contact.LastName,
						},
					},
				},
				{
					FieldID: ContactFieldIDMiddleName,
					Values: []dto.Value{
						{
							Value: contact.MiddleName,
						},
					},
				},
				{
					FieldID: ContactFieldIDPhone,
					Values: []dto.Value{
						{
							Value:  contact.Phone,
							EnumID: int(contact.PhoneType),
						},
					},
				},
				{
					FieldID: ContactFieldIDEmail,
					Values: []dto.Value{
						{
							Value:  contact.Email,
							EnumID: int(contact.EmailType),
						},
					},
				},
			},
		},
	})
	if err != nil {
		return 0, fmt.Errorf("create contacts: %w", err)
	}

	if len(resp.Embedded.Contacts) == 0 {
		return 0, fmt.Errorf("create contacts: empty response")
	}

	return resp.Embedded.Contacts[0].ID, nil
}
