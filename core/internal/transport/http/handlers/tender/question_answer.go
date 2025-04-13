package tender

import (
	"context"
	"errors"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (h *Handler) V1TendersTenderIDQuestionAnswerPost(ctx context.Context, req *api.V1TendersTenderIDQuestionAnswerPostReq, params api.V1TendersTenderIDQuestionAnswerPostParams) (api.V1TendersTenderIDQuestionAnswerPostRes, error) {
	questionAnswer, err := h.questionAnswerService.Create(ctx, service.CreateQuestionAnswerParams{
		TenderID:             params.TenderID,
		AuthorOrganizationID: contextor.GetOrganizationID(ctx),
		ParentID:             models.Optional[int]{Value: req.GetParentID().Value, Set: req.GetParentID().Set},
		Type:                 models.APIToQuestionAnswerType(params.Type),
		Content:              req.Content})

	switch {
	case errors.Is(err, errstore.ErrQuestionAnswerNotFound):
		return nil, cerr.Wrap(err, cerr.CodeNotFound, "Вопрос не найден", nil)
	case errors.Is(err, errstore.ErrQuestionAnswerUniqueViolation):
		return nil, cerr.Wrap(err, cerr.CodeConflict, "Ответ на вопрос уже существует", nil)
	case err != nil:
		return nil, fmt.Errorf("create question-answer: %w", err)
	}

	return &api.V1TendersTenderIDQuestionAnswerPostCreated{
		Data: models.ConvertQuestionAnswerToAPI(questionAnswer),
	}, nil
}

func (h *Handler) V1TendersTenderIDQuestionAnswerGet(ctx context.Context, params api.V1TendersTenderIDQuestionAnswerGetParams) (api.V1TendersTenderIDQuestionAnswerGetRes, error) {
	questionWithAnswer, err := h.questionAnswerService.Get(ctx, params.TenderID)
	if err != nil {
		if errors.Is(err, errstore.ErrTenderNotFound) {
			return nil, cerr.Wrap(err, cerr.CodeNotFound, "Тендер не найден", nil)
		}
		return nil, fmt.Errorf("get question-answer: %w", err)
	}

	return &api.V1TendersTenderIDQuestionAnswerGetOK{
		Data: convert.Slice[[]models.QuestionWithAnswer, []api.QuestionWithAnswer](questionWithAnswer, models.ConvertQuestionWithAnswerToAPI),
	}, nil
}
