package verification

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
)

func (s *Service) Create(ctx context.Context, params service.VerificationRequestCreateParams) error {
	return s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		if err := s.verificationStore.Create(ctx, qe, store.VerificationRequestCreateParams{
			ObjectID:    params.ObjectID,
			ObjectType:  params.ObjectType,
			Attachments: params.Attachments,
		}); err != nil {
			return fmt.Errorf("create verification: %w", err)
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
				Id:   int32(params.ObjectID),
				Type: modelsv1.ObjectType(params.ObjectType),
			},
		}

		var topic broker.Topic
		switch params.ObjectType {
		case models.ObjectTypeOrganization:
			topic = broker.NotifyOrganizationVerification

			err := s.organizationStore.UpdateVerificationStatus(ctx, qe, store.OrganizationUpdateVerifStatusParams{
				OrganizationID:     params.ObjectID,
				VerificationStatus: models.VerificationStatusInReview,
			})
			if err != nil {
				return fmt.Errorf("update organization verification status: %w", err)
			}
		default:
			return fmt.Errorf("invalid object type: %v", params.ObjectType)
		}

		b, err := proto.Marshal(&eventsv1.SentNotification{Notification: notification})
		if err != nil {
			return fmt.Errorf("marhal notification proto: %w", err)
		}

		err = s.broker.Publish(ctx, topic, b)
		if err != nil {
			return fmt.Errorf("notification: %w", err)
		}

		return nil
	})
}
