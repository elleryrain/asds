package catalog

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	catalogService "gitlab.ubrato.ru/ubrato/core/internal/service/catalog"
)

func (h *Handler) V1CatalogServicesGet(ctx context.Context) (api.V1CatalogServicesGetRes, error) {
	services, err := h.svc.GetServices(ctx)
	if err != nil {
		return nil, fmt.Errorf("get objects catalog: %w", err)
	}

	return &api.V1CatalogServicesGetOK{
		Data: convert.Slice[models.Services, api.Services](services, models.ConvertModelServiceToApi),
	}, nil
}

func (h *Handler) V1CatalogServicesPost(ctx context.Context, req *api.V1CatalogServicesPostReq) (api.V1CatalogServicesPostRes, error) {
	service, err := h.svc.CreateService(ctx, catalogService.CreateServiceParams{
		Name:     req.GetName(),
		ParentID: req.ParentID.Value,
	})
	if err != nil {
		return nil, fmt.Errorf("create catalog object: %w", err)
	}

	return &api.V1CatalogServicesPostCreated{
		Data: models.ConvertModelServiceToApi(service),
	}, nil
}
