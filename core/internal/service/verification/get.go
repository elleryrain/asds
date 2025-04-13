package verification

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/deduplicate"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"

	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) GetByID(ctx context.Context, requestID int) (models.VerificationRequest[models.VerificationObject], error) {
	request, err := s.verificationStore.GetByIDWithEmptyObject(ctx, s.psql.DB(), requestID)
	if err != nil {
		return models.VerificationRequest[models.VerificationObject]{}, fmt.Errorf("get req by id=%v: %w", requestID, err)
	}

	var object models.VerificationObject
	switch request.ObjectType {
	case models.ObjectTypeOrganization:
		object, err = s.organizationStore.GetByID(ctx, s.psql.DB(), request.ObjectID)

	case models.ObjectTypeTender:
		object, err = s.tenderStore.GetByID(ctx, s.psql.DB(), request.ObjectID)

	case models.ObjectTypeAddition:
		object, err = s.additionStore.GetByID(ctx, s.psql.DB(), request.ObjectID)

	case models.ObjectTypeQuestionAnswer:
		object, err = s.questionAnswerStore.GetByID(ctx, s.psql.DB(), request.ObjectID)
	}
	if err != nil {
		return models.VerificationRequest[models.VerificationObject]{}, fmt.Errorf("get object type=%v by id=%v: %w", request.ObjectType, request.ObjectID, err)
	}

	request.Object = object

	return request, nil
}

func (s *Service) Get(ctx context.Context, params service.VerificationRequestsObjectGetParams) (models.VerificationRequestPagination[models.VerificationObject], error) {
	requests, err := s.verificationStore.GetWithEmptyObject(ctx, s.psql.DB(), store.VerificationRequestsObjectGetParams{
		ObjectType: models.NewOptional(params.ObjectType),
		ObjectID:   params.ObjectID,
		Status:     params.Status,
		Limit:      models.Optional[uint64]{Value: params.PerPage, Set: params.PerPage != 0},
		Offset:     models.Optional[uint64]{Value: params.Page * params.PerPage, Set: (params.Page * params.PerPage) != 0}})
	if err != nil {
		return models.VerificationRequestPagination[models.VerificationObject]{}, fmt.Errorf("get object req: %w", err)
	}

	count, err := s.verificationStore.Count(ctx, s.psql.DB(), store.VerificationRequestsObjectGetCountParams{
		ObjectType: models.NewOptional(params.ObjectType),
		ObjectID:   params.ObjectID,
		Status:     params.Status})
	if err != nil {
		return models.VerificationRequestPagination[models.VerificationObject]{}, fmt.Errorf("get count objects: %w", err)
	}

	var objectIDs []int
	for _, req := range requests {
		objectIDs = append(objectIDs, req.ObjectID)
	}
	objectIDs = deduplicate.Deduplicate(objectIDs)

	objectMap := map[int]models.VerificationObject{}
	switch params.ObjectType {
	case models.ObjectTypeOrganization:
		organizations, err := s.organizationStore.Get(ctx, s.psql.DB(), store.OrganizationGetParams{
			OrganizationIDs: objectIDs})
		if err != nil {
			return models.VerificationRequestPagination[models.VerificationObject]{}, fmt.Errorf("get organizations: %w", err)
		}

		for _, organization := range organizations {
			objectMap[organization.ID] = organization
		}

	case models.ObjectTypeTender:
		tenders, err := s.tenderStore.List(ctx, s.psql.DB(), store.TenderListParams{
			TenderIDs: models.Optional[[]int]{Value: objectIDs, Set: true}})
		if err != nil {
			return models.VerificationRequestPagination[models.VerificationObject]{}, fmt.Errorf("get tenders: %w", err)
		}

		for _, tender := range tenders {
			objectMap[tender.ID] = tender
		}

	case models.ObjectTypeAddition:
		additions, err := s.additionStore.Get(ctx, s.psql.DB(), store.AdditionGetParams{
			AdditionIDs: models.Optional[[]int]{Value: objectIDs, Set: true}})
		if err != nil {
			return models.VerificationRequestPagination[models.VerificationObject]{}, fmt.Errorf("get additions: %w", err)
		}

		for _, addition := range additions {
			objectMap[addition.ID] = addition
		}

	case models.ObjectTypeQuestionAnswer:
		questionAnswers, err := s.questionAnswerStore.Get(ctx, s.psql.DB(), store.QuestionAnswerGetParams{
			QuestionAnswerIDs: models.Optional[[]int]{Value: objectIDs, Set: true}})
		if err != nil {
			return models.VerificationRequestPagination[models.VerificationObject]{}, fmt.Errorf("get question-answers: %w", err)
		}

		for _, qa := range questionAnswers {
			objectMap[qa.Question.ID] = qa
			if qa.Answer.Set {
				objectMap[qa.Answer.Value.ID] = qa
			}
		}

	default:
		return models.VerificationRequestPagination[models.VerificationObject]{}, fmt.Errorf("invalid object type: %v", params.ObjectType)
	}

	for i := range requests {
		if object, ok := objectMap[requests[i].ObjectID]; ok {
			requests[i].Object = object
		}
	}

	return models.VerificationRequestPagination[models.VerificationObject]{
		VerificationRequests: requests,
		Pagination:           pagination.New(params.Page, params.PerPage, uint64(count)),
	}, nil
}
