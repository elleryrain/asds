package suggest

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
)

func (h *Handler) V1SuggestCompanyGet(ctx context.Context, params api.V1SuggestCompanyGetParams) (api.V1SuggestCompanyGetRes, error) {
	name, err := h.svc.Company(ctx, string(params.Inn))
	if err != nil {
		return nil, fmt.Errorf("sugget company: %w", err)
	}

	return &api.V1SuggestCompanyGetOK{
		Data: api.V1SuggestCompanyGetOKData{
			Name: name,
		},
	}, nil
}
