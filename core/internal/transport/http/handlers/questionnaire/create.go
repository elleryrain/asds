package questionnaire

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

func (h *Handler) V1QuestionnaireOrganizationIDPost(ctx context.Context, req *api.V1QuestionnaireOrganizationIDPostReq, params api.V1QuestionnaireOrganizationIDPostParams) (api.V1QuestionnaireOrganizationIDPostRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to submit questionnaire answers", nil)
	}

	if err := h.questionnaireService.Create(ctx, service.QuestionnaireCreateParams{
		OrganizationID: params.OrganizationID,
		Answers:        convert.Slice[[]api.QuestionnaireAnswer, []models.Answer](req.Answers, models.ConvertAPIToAnswer),
		IsCompleted:    req.GetIsCompleted(),
	}); err != nil {
		if errors.Is(err, errstore.ErrQuestionnaireExist) {
			return nil, cerr.Wrap(err, cerr.CodeConflict, "The questionnaire has been completed", nil)
		}
		return nil, fmt.Errorf("create questionnaire: %w", err)
	}

	return &api.V1QuestionnaireOrganizationIDPostOK{}, nil
}
