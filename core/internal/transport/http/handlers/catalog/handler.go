package catalog

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	catalogService "gitlab.ubrato.ru/ubrato/core/internal/service/catalog"
)

type Handler struct {
	logger *slog.Logger
	svc    Service
}

type Service interface {
	GetObjects(ctx context.Context) (models.Objects, error)
	GetServices(ctx context.Context) (models.Services, error)
	CreateCity(ctx context.Context, params catalogService.CreateCityParams) (models.City, error)
	CreateRegion(ctx context.Context, params catalogService.CreateRegionParams) (models.Region, error)
	CreateObject(ctx context.Context, params catalogService.CreateObjectParams) (models.Object, error)
	CreateService(ctx context.Context, params catalogService.CreateServiceParams) (models.Service, error)
	GetMeasurements(ctx context.Context) ([]models.Measure, error)
}

func New(logger *slog.Logger, svc Service) *Handler {
	return &Handler{
		logger: logger,
		svc:    svc,
	}
}
