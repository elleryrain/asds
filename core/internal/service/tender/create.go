package tender

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	eventsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/events/v1"
	modelsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/models/v1"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) Create(ctx context.Context, params service.TenderCreateParams) (models.Tender, error) {
	id, err := s.tenderStore.Create(ctx, s.psql.DB(), store.TenderCreateParams{
		Name:            params.Name,
		CityID:          params.CityID,
		OrganizationID:  params.OrganizationID,
		ServiceIDs:      params.ServiceIDs,
		ObjectIDs:       params.ObjectIDs,
		Price:           params.Price,
		IsContractPrice: params.IsContractPrice,
		IsNDSPrice:      params.IsNDSPrice,
		IsDraft:         params.IsDraft,
		FloorSpace:      params.FloorSpace,
		Description:     params.Description,
		Wishes:          params.Wishes,
		Specification:   params.Specification,
		Attachments:     params.Attachments,
		ReceptionStart:  params.ReceptionStart,
		ReceptionEnd:    params.ReceptionEnd,
		WorkStart:       params.WorkStart,
		WorkEnd:         params.WorkEnd,
	})
	if err != nil {
		return models.Tender{}, fmt.Errorf("create tender: %w", err)
	}

	if !params.IsDraft {
		err = s.verificationStore.Create(ctx, s.psql.DB(), store.VerificationRequestCreateParams{
			ObjectID:   id,
			ObjectType: models.ObjectTypeTender,
		})
		if err != nil {
			return models.Tender{}, fmt.Errorf("create verification request: %w", err)
		}

		// Уведомления
		notification := &modelsv1.Notification{
			User: &modelsv1.NotifiedUser{
				Id: *proto.Int32(int32(contextor.GetUserID(ctx))),
			},
			Verification: &modelsv1.Verification{
				Status: modelsv1.Status(models.VerificationStatusInReview),
			},
			Object: &modelsv1.Object{
				Id:   int32(id),
				Type: modelsv1.ObjectType(models.ObjectTypeTender),
				Tender: &modelsv1.Tender{
					Title:          params.Name,
					ReceptionStart: timestamppb.New(params.ReceptionStart),
				},
			},
		}

		b, err := proto.Marshal(&eventsv1.SentNotification{Notification: notification})
		if err != nil {
			return models.Tender{}, fmt.Errorf("marhal notification proto: %w", err)
		}

		err = s.broker.Publish(ctx, broker.NotifyTenderVerification, b)
		if err != nil {
			return models.Tender{}, fmt.Errorf("notification: %w", err)
		}
	}

	tender, err := s.tenderStore.GetByID(ctx, s.psql.DB(), id)
	if err != nil {
		return models.Tender{}, fmt.Errorf("get tender: %w", err)
	}

	return tender, nil
}
