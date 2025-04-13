package organization

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

type Handler struct {
	logger              *slog.Logger
	organizationService OrganizationService
	portfolioService    PortfolioService
	favouriteService    FavouriteService
}

type OrganizationService interface {
	Get(ctx context.Context, params service.OrganizationGetParams) (models.OrganizationsPagination, error)
	GetByID(ctx context.Context, id int) (models.Organization, error)
	GetCustomer(ctx context.Context, organizationId int) (models.Organization, error)
	GetContractorByID(ctx context.Context, organizationId int) (models.Organization, error)
	GetContractors(ctx context.Context, params service.OrganizationContractorsGetParams) (models.OrganizationsPagination, error)
	UpdateBrand(ctx context.Context, params service.OrganizationUpdateBrandParams) error
	UpdateContacts(ctx context.Context, params service.OrganizationUpdateContactsParams) error
	UpdateCustomer(ctx context.Context, params service.OrganizationUpdateCustomerParams) (models.Organization, error)
	UpdateContractor(ctx context.Context, params service.OrganizationUpdateContractorParams) (models.Organization, error)
}

type PortfolioService interface {
	Create(ctx context.Context, params service.PortfolioCreateParams) (models.Portfolio, error)
	Delete(ctx context.Context, id int) error
	Get(ctx context.Context, params service.PortfolioGetParams) ([]models.Portfolio, error)
	Update(ctx context.Context, params service.PortfolioUpdateParams) (models.Portfolio, error)
}

type FavouriteService interface {
	Create(ctx context.Context, params service.FavouriteCreateParams) (int64, error)
	Get(ctx context.Context, params service.FavouriteGetParams) (models.FavouritePagination[models.FavouriteObject], error)
	Delete(ctx context.Context, favouriteID int) error
}

func New(
	logger *slog.Logger,
	organizationService OrganizationService,
	portfolioService PortfolioService,
	favouriteService FavouriteService) *Handler {
	return &Handler{
		logger:              logger,
		organizationService: organizationService,
		portfolioService:    portfolioService,
		favouriteService:    favouriteService,
	}
}
