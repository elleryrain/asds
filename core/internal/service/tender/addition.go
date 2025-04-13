package tender

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	eventsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/events/v1"
	modelsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/models/v1"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"google.golang.org/protobuf/proto"
)

func (s *Service) CreateAddition(ctx context.Context, params service.AdditionCreateParams) error {
	tender, err := s.tenderStore.GetByID(ctx, s.psql.DB(), params.TenderID)
	if err != nil {
		return fmt.Errorf("get tender: %w", err)
	}

	if tender.Organization.ID != contextor.GetOrganizationID(ctx) {
		return cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "Недостаточно прав для добавления дополнительной информации", nil)
	}

	if tender.VerificationStatus != models.VerificationStatusApproved {
		return cerr.Wrap(
			fmt.Errorf("tender status not approved"),
			cerr.CodeUnprocessableEntity,
			"Нельзя добавить дополнительную информацию, так как тендер не прошел модерацию", nil)
	}

	if err := s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		additionID, err := s.additionStore.CreateAddition(ctx, qe, store.AdditionCreateParams{
			TenderID:    params.TenderID,
			Title:       params.Title,
			Content:     params.Content,
			Attachments: params.Attachments})
		if err != nil {
			return fmt.Errorf("creating addition %w", err)
		}

		err = s.verificationStore.Create(ctx, qe, store.VerificationRequestCreateParams{
			ObjectID:   additionID,
			ObjectType: models.ObjectTypeAddition})
		if err != nil {
			return fmt.Errorf("create verification request: %w", err)
		}

		// уведомления
		b, err := proto.Marshal(&eventsv1.SentNotification{
			Notification: &modelsv1.Notification{
				User: &modelsv1.NotifiedUser{
					Id: *proto.Int32(int32(contextor.GetUserID(ctx))),
				},
				Verification: &modelsv1.Verification{
					Status: modelsv1.Status_STATUS_IN_REVIEW,
				},
				Object: &modelsv1.Object{
					Id:   int32(additionID),
					Type: modelsv1.ObjectType_ObjectTypeAddition,
					Tender: &modelsv1.Tender{
						Id:    int32(tender.ID),
						Title: tender.Name,
					},
				},
			}})
		if err != nil {
			return fmt.Errorf("marhal notification proto: %w", err)
		}

		err = s.broker.Publish(ctx, broker.NotifyTenderAdditionVerification, b)
		if err != nil {
			return fmt.Errorf("notification: %w", err)
		}

		return nil
	}); err != nil {
		return fmt.Errorf("run transaction: %w", err)
	}

	return nil
}

func (s *Service) GetAdditions(ctx context.Context, params service.GetAdditionParams) ([]models.Addition, error) {
	tender, err := s.tenderStore.GetByID(ctx, s.psql.DB(), params.TenderID)
	if err != nil {
		return nil, fmt.Errorf("get tender: %w", err)
	}

	additions, err := s.additionStore.Get(ctx, s.psql.DB(), store.AdditionGetParams{
		TenderID:     models.NewOptional(params.TenderID),
		VerifiedOnly: tender.Organization.ID != contextor.GetOrganizationID(ctx)})
	if err != nil {
		return nil, fmt.Errorf("get Addition: %w", err)
	}

	return additions, nil
}
