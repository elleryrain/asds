package survey

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	surveyService "gitlab.ubrato.ru/ubrato/core/internal/service/survey"
)

func (h *Handler) V1SurveyPost(ctx context.Context, req *api.V1SurveyPostReq) (api.V1SurveyPostRes, error) {
	err := h.svc.Response(ctx, surveyService.ResponseParams{
		Name:     req.GetName(),
		Type:     string(req.GetType()),
		Phone:    string(req.GetPhone()),
		Question: req.GetQuestion(),
	})
	if err != nil {
		return nil, fmt.Errorf("send survey response: %w", err)
	}
	return &api.V1SurveyPostOK{}, nil
}
