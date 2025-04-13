package catalog

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	catalogService "gitlab.ubrato.ru/ubrato/core/internal/service/catalog"
)

func (h *Handler) V1CatalogObjectsGet(ctx context.Context) (api.V1CatalogObjectsGetRes, error) {
	objects, err := h.svc.GetObjects(ctx)
	if err != nil {
		return nil, fmt.Errorf("get objects catalog: %w", err)
	}

	return &api.V1CatalogObjectsGetOK{
		Data: convert.Slice[models.Objects, api.Objects](objects, models.ConvertModelObjectToApi),
	}, nil
}

func (h *Handler) V1CatalogObjectsPost(ctx context.Context, req *api.V1CatalogObjectsPostReq) (api.V1CatalogObjectsPostRes, error) {
	object, err := h.svc.CreateObject(ctx, catalogService.CreateObjectParams{
		Name:     req.GetName(),
		ParentID: req.ParentID.Value,
	})
	if err != nil {
		return nil, fmt.Errorf("create catalog object: %w", err)
	}

	return &api.V1CatalogObjectsPostCreated{
		Data: models.ConvertModelObjectToApi(object),
	}, nil
}
