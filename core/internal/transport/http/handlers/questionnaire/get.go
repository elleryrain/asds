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

func (h *Handler) V1QuestionnaireGet(ctx context.Context, params api.V1QuestionnaireGetParams) (api.V1QuestionnaireGetRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	questionnaires, err := h.questionnaireService.Get(ctx, service.QuestionnaireGetParams{
		Offset: uint64(params.Offset.Or(0)),
		Limit:  uint64(params.Limit.Or(100)),
	})
	if err != nil {
		return nil, fmt.Errorf("get questionnaires: %w", err)
	}

	return &api.V1QuestionnaireGetOK{
		Data: convert.Slice[[]models.Questionnaire, []api.Questionnaire](questionnaires, models.ConvertQuestionnaireToAPI),
	}, nil
}

func (h *Handler) V1QuestionnaireOrganizationIDStatusGet(ctx context.Context, params api.V1QuestionnaireOrganizationIDStatusGetParams) (api.V1QuestionnaireOrganizationIDStatusGetRes, error) {
	status, err := h.questionnaireService.GetStatus(ctx, params.OrganizationID)
	if err != nil {
		if errors.Is(err, errstore.ErrQuestionnaireNotFound) {
			return &api.V1QuestionnaireOrganizationIDStatusGetOK{
				Data: api.V1QuestionnaireOrganizationIDStatusGetOKData{IsCompleted: false}}, nil
		}
		return nil, fmt.Errorf("get questionnaire status: %w", err)
	}

	return &api.V1QuestionnaireOrganizationIDStatusGetOK{
		Data: api.V1QuestionnaireOrganizationIDStatusGetOKData{
			IsCompleted: status,
		},
	}, nil
}
