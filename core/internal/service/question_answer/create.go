package questionanswer

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

func (s *Service) Create(ctx context.Context, params service.CreateQuestionAnswerParams) (models.QuestionAnswer, error) {
	if params.Type == models.QuestionAnswerTypeAnswer {
		question, err := s.questionAnswerStore.GetByID(ctx, s.psql.DB(), params.ParentID.Value)
		if err != nil {
			return models.QuestionAnswer{}, fmt.Errorf("get question: %w", err)
		}

		if question.Question.VerificationStatus != models.VerificationStatusApproved {
			return models.QuestionAnswer{}, cerr.Wrap(
				fmt.Errorf("question status not approved"),
				cerr.CodeUnprocessableEntity,
				"Нельзя отправить ответ, так как вопрос не прошел модерацию", nil)
		}
	}

	var questionAnswer models.QuestionAnswer
	if err := s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		newQuestionAnswer, err := s.questionAnswerStore.Create(ctx, qe, store.CreateQuestionAnswerParams{
			TenderID:             params.TenderID,
			AuthorOrganizationID: params.AuthorOrganizationID,
			ParentID:             params.ParentID,
			Type:                 params.Type,
			Content:              params.Content})
		if err != nil {
			return fmt.Errorf("create question-answer %w", err)
		}
		questionAnswer = newQuestionAnswer

		err = s.verificationStore.Create(ctx, qe, store.VerificationRequestCreateParams{
			ObjectID:   newQuestionAnswer.ID,
			ObjectType: models.ObjectTypeQuestionAnswer})
		if err != nil {
			return fmt.Errorf("create verification request: %w", err)
		}

		// уведомление
		tenderNotifyInfo, err := s.tenderStore.GetTenderNotifyInfoByObjectID(ctx, qe, store.TenderNotifyInfoParams{QuestionAnswerID: models.NewOptional(questionAnswer.ID)})
		if err != nil {
			return fmt.Errorf("get tender info id=%v: %w", params.TenderID, err)
		}

		b, err := proto.Marshal(&eventsv1.SentNotification{
			Notification: &modelsv1.Notification{
				User: &modelsv1.NotifiedUser{
					Id: *proto.Int32(int32(contextor.GetUserID(ctx))),
					IsContractor: params.Type == models.QuestionAnswerTypeQuestion,
				},
				Verification: &modelsv1.Verification{
					Status: modelsv1.Status_STATUS_IN_REVIEW,
				},
				Object: &modelsv1.Object{
					Id:   int32(newQuestionAnswer.ID),
					Type: modelsv1.ObjectType_ObjectTypeQuestionAnswer,
					Tender: &modelsv1.Tender{
						Id:    int32(params.TenderID),
						Title: tenderNotifyInfo.Name,
					},
				},
			}})
		if err != nil {
			return fmt.Errorf("marhal notification proto: %w", err)
		}

		err = s.broker.Publish(ctx, broker.NotifyTenderQuestionAnswerVerification, b)
		if err != nil {
			return fmt.Errorf("notification: %w", err)
		}

		return nil
	}); err != nil {
		return models.QuestionAnswer{}, fmt.Errorf("run transaction: %w", err)
	}

	return questionAnswer, nil
}
