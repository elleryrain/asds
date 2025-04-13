package verification

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	eventsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/events/v1"
	modelsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/models/v1"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func (s *Service) UpdateStatus(ctx context.Context, params service.VerificationRequestUpdateStatusParams) error {
	return s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		result, err := s.verificationStore.UpdateStatus(ctx, qe, store.VerificationRequestUpdateStatusParams{
			UserID:        params.UserID,
			RequestID:     params.RequesID,
			Status:        params.Status,
			ReviewComment: params.ReviewComment,
		})
		if err != nil {
			return fmt.Errorf("update request status: %w", err)
		}

		notification := &modelsv1.Notification{
			User: &modelsv1.NotifiedUser{},
			Verification: &modelsv1.Verification{
				Status:  modelsv1.Status(params.Status),
				Comment: params.ReviewComment.Value,
			},
			Object: &modelsv1.Object{
				Id:   int32(result.ObjectID),
				Type: modelsv1.ObjectType(result.ObjectType),
			},
		}

		switch result.ObjectType {
		case models.ObjectTypeOrganization:
			if err = s.organizationStore.UpdateVerificationStatus(ctx, qe, store.OrganizationUpdateVerifStatusParams{
				OrganizationID:     result.ObjectID,
				VerificationStatus: params.Status}); err != nil {
				return fmt.Errorf("update organization verif status: %w", err)
			}

			isContractor, err := s.organizationStore.GetIsContractorByID(ctx, qe, result.ObjectID)
			if err != nil {
				return fmt.Errorf("get organization is_contractor by id: %w", err)
			}
			notification.User.IsContractor = isContractor

			return s.SendUserNotify(ctx, notification, broker.NotifyOrganizationVerification, result.ObjectID)

		case models.ObjectTypeTender:
			status := models.InvalidStatus
			switch params.Status {
			case models.VerificationStatusApproved:
				status = models.ReceptionStatus
			case models.VerificationStatusDeclined:
				status = models.RemovedByModeratorStatus
			}

			err = s.tenderStore.UpdateVerificationStatus(ctx, qe, store.TenderUpdateVerifStatusParams{
				TenderID:           result.ObjectID,
				VerificationStatus: params.Status,
				Status:             status})
			if err != nil {
				return fmt.Errorf("update tender verif status: %w", err)
			}

			tenderNotifyInfo, err := s.tenderStore.GetTenderNotifyInfoByObjectID(ctx, qe, store.TenderNotifyInfoParams{TenderID: models.NewOptional(result.ObjectID)})
			if err != nil {
				return fmt.Errorf("get tender notify: %w", err)
			}

			notification.Object.Tender = &modelsv1.Tender{
				Id:             int32(tenderNotifyInfo.ID),
				Title:          tenderNotifyInfo.Name,
				ReceptionStart: timestamppb.New(tenderNotifyInfo.ReceptionStart),
			}

			return s.SendUserNotify(ctx, notification, broker.NotifyTenderVerification, tenderNotifyInfo.Organization.ID)

		case models.ObjectTypeAddition:
			if err = s.additionStore.UpdateVerificationStatus(ctx, qe, store.AdditionUpdateVerifStatusParams{
				AdditionID:         result.ObjectID,
				VerificationStatus: params.Status}); err != nil {
				return fmt.Errorf("update tender ddition verif status: %w", err)
			}

			tenderNotifyInfo, err := s.tenderStore.GetTenderNotifyInfoByObjectID(ctx, qe, store.TenderNotifyInfoParams{AdditionID: models.NewOptional(result.ObjectID)})
			if err != nil {
				return fmt.Errorf("get tender notify: %w", err)
			}

			notification.Object.Tender = &modelsv1.Tender{
				Id:    int32(tenderNotifyInfo.ID),
				Title: tenderNotifyInfo.Name,
			}

			return s.SendUserNotify(ctx, notification, broker.NotifyTenderAdditionVerification, tenderNotifyInfo.Organization.ID)

		case models.ObjectTypeQuestionAnswer:
			if err = s.questionAnswerStore.UpdateVerificationStatus(ctx, qe, store.QuestionAnswerVerifStatusUpdateParams{
				QuestionAnswerID:   result.ObjectID,
				VerificationStatus: params.Status}); err != nil {
				return fmt.Errorf("update qa verif status: %w", err)
			}

			// уведомления
			tenderNotifyInfo, err := s.tenderStore.GetTenderNotifyInfoByObjectID(ctx, qe, store.TenderNotifyInfoParams{QuestionAnswerID: models.NewOptional(result.ObjectID)})
			if err != nil {
				return fmt.Errorf("get tender notify: %w", err)
			}

			eventTender := &modelsv1.Tender{
				Id:    int32(tenderNotifyInfo.ID),
				Title: tenderNotifyInfo.Name,
			}
			notification.Object.Tender = eventTender

			// получение вопроса с ответом
			qa, err := s.questionAnswerStore.GetByID(ctx, qe, result.ObjectID)
			if err != nil {
				return fmt.Errorf("update request status: %w", err)
			}

			// 1. Верификация
			// уведомление для исполнителя о верификации вопроса
			organizationID := qa.Question.AuthorOrganizationID
			if qa.Answer.Set {
				// уведомление для заказчика о верификации ответа
				organizationID = qa.Answer.Value.AuthorOrganizationID
			}

			notification.User.IsContractor = !qa.Answer.Set

			// Отправка автору вопроса/ответа уведомления о верификации
			if err = s.SendUserNotify(ctx, notification, broker.NotifyTenderQuestionAnswerVerification, organizationID); err != nil {
				return fmt.Errorf("send verification notify: %w", err)
			}

			// 2. Вопрос ответ тендера
			if params.Status == models.VerificationStatusApproved {
				if qa.Answer.Set {
					// Отправка уведомления исполнителю о наличии ответа
					organizationID = qa.Question.AuthorOrganizationID
				} else {
					// Отправка уведомления заказчику о наличии вопроса	
					organizationID = tenderNotifyInfo.Organization.ID
				}


				return s.SendUserNotify(ctx, &modelsv1.Notification{
					User: &modelsv1.NotifiedUser{IsContractor: qa.Answer.Set},
					Object: &modelsv1.Object{
						Id:     int32(result.ObjectID),
						Type:   modelsv1.ObjectType(result.ObjectType),
						Tender: eventTender,
					},
				}, broker.NotifyTenderQA, organizationID)
			}
		default:
			return fmt.Errorf("invalid object type: %v", result.ObjectType)
		}

		return nil
	})
}

func (s *Service) SendUserNotify(ctx context.Context, notify *modelsv1.Notification, topic broker.Topic, organizationID int) error {
	userID, err := s.userStore.GetUserIDByOrganizationID(ctx, s.psql.DB(), organizationID)
	if err != nil {
		return fmt.Errorf("get userID by orgID: %w", err)
	}
	notify.User.Id = int32(userID)

	b, err := proto.Marshal(&eventsv1.SentNotification{Notification: notify})
	if err != nil {
		return fmt.Errorf("marhal notification proto: %w", err)
	}

	return s.broker.Publish(ctx, topic, b)
}
