package organization

import (
	"context"
	"errors"
	"fmt"
	"net/url"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (h *Handler) V1OrganizationsOrganizationIDPortfolioPost(
	ctx context.Context,
	req *api.V1OrganizationsOrganizationIDPortfolioPostReq,
	params api.V1OrganizationsOrganizationIDPortfolioPostParams,
) (api.V1OrganizationsOrganizationIDPortfolioPostRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to create the portfolio", nil)
	}

	portfolio, err := h.portfolioService.Create(ctx, service.PortfolioCreateParams{
		OrganizationID: params.OrganizationID,
		Title:          string(req.Title),
		Description:    string(req.Description),
		Attachments: convert.Slice[[]url.URL, []string](
			req.GetAttachments(), func(u url.URL) string { return u.String() })})
	if err != nil {
		if errors.Is(err, errstore.ErrOrganizationNotFound) {
			return nil, cerr.Wrap(err, cerr.CodeNotFound, "Организация не существует", nil)
		}
		return nil, fmt.Errorf("comment: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDPortfolioPostOK{Data: models.ConvertPortfolioModelToApi(portfolio)}, nil

}

func (h *Handler) V1OrganizationsPortfolioPortfolioIDDelete(ctx context.Context, params api.V1OrganizationsPortfolioPortfolioIDDeleteParams) (api.V1OrganizationsPortfolioPortfolioIDDeleteRes, error) {
	if err := h.portfolioService.Delete(ctx, params.PortfolioID); err != nil {
		if errors.Is(err, errstore.ErrPortfolioNotFound) {
			return nil, cerr.Wrap(err, cerr.CodeNotFound, "Портфолио не существует", nil)
		}
		return nil, fmt.Errorf("portfilio delete: %w", err)
	}

	return &api.V1OrganizationsPortfolioPortfolioIDDeleteOK{}, nil
}

func (h *Handler) V1OrganizationsOrganizationIDPortfolioGet(ctx context.Context, params api.V1OrganizationsOrganizationIDPortfolioGetParams) (api.V1OrganizationsOrganizationIDPortfolioGetRes, error) {
	portfolio, err := h.portfolioService.Get(ctx, service.PortfolioGetParams{
		OrganizationID: params.OrganizationID,
	})
	if err != nil {
		return nil, fmt.Errorf("get porrtfolios: %w", err)
	}

	return &api.V1OrganizationsOrganizationIDPortfolioGetOK{
		Data: convert.Slice[[]models.Portfolio, []api.Portfolio](portfolio, models.ConvertPortfolioModelToApi),
	}, nil
}

func (h *Handler) V1OrganizationsPortfolioPortfolioIDPut(ctx context.Context, req *api.V1OrganizationsPortfolioPortfolioIDPutReq, params api.V1OrganizationsPortfolioPortfolioIDPutParams) (api.V1OrganizationsPortfolioPortfolioIDPutRes, error) {
	portfolio, err := h.portfolioService.Update(ctx, service.PortfolioUpdateParams{
		PortfolioID: params.PortfolioID,
		Title:       models.Optional[string]{Value: string(req.Title.Value), Set: req.Title.Set},
		Description: models.Optional[string]{Value: string(req.Description.Value), Set: req.Description.Set},
		Attachments: models.Optional[[]string]{Value: convert.Slice[[]url.URL, []string](
			req.GetAttachments(), func(u url.URL) string { return u.String() },
		), Set: req.GetAttachments() != nil}})
	switch {
	case errors.Is(err, errstore.ErrPortfolioNotFound):
		return nil, cerr.Wrap(err, cerr.CodeNotFound, "Портфолио не существует", nil)
	case err != nil:
		return nil, fmt.Errorf("update portfolio: %w", err)
	}

	return &api.V1OrganizationsPortfolioPortfolioIDPutOK{
		Data: models.ConvertPortfolioModelToApi(portfolio),
	}, nil
}
