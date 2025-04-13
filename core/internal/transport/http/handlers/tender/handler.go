package tender

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

type Handler struct {
	logger                *slog.Logger
	tenderService         TenderService
	questionAnswerService QuestionAnswerService
	respondService        RespondService
	winnersStore          WinnersStore
}

type TenderService interface {
	Create(ctx context.Context, params service.TenderCreateParams) (models.Tender, error)
	Update(ctx context.Context, params service.TenderUpdateParams) (models.Tender, error)
	GetByID(ctx context.Context, tenderID int) (models.Tender, error)
	List(ctx context.Context, params service.TenderListParams) (models.TendersPagination, error)

	CreateAddition(ctx context.Context, params service.AdditionCreateParams) error
	GetAdditions(ctx context.Context, params service.GetAdditionParams) ([]models.Addition, error)
}

type QuestionAnswerService interface {
	Create(ctx context.Context, params service.CreateQuestionAnswerParams) (models.QuestionAnswer, error)
	Get(ctx context.Context, tenderID int) ([]models.QuestionWithAnswer, error)
}

type RespondService interface {
	Create(ctx context.Context, params service.RespondCreateParams) error
	Get(ctx context.Context, params service.RespondGetParams) (models.RespondPagination, error)
}

type WinnersStore interface {
	Create(ctx context.Context, params service.WinnersCreateParams) (models.Winners, error)
	Get(ctx context.Context, tenderID int) ([]models.Winners, error)
	UpdateStatus(ctx context.Context, params service.WinnerUpdateParams) error
}

func New(
	logger *slog.Logger,
	tenderService TenderService,
	questionAnswerService QuestionAnswerService,
	respondService RespondService,
	winnersStore WinnersStore,
) *Handler {
	return &Handler{
		logger:                logger,
		tenderService:         tenderService,
		questionAnswerService: questionAnswerService,
		respondService:        respondService,
		winnersStore:          winnersStore,
	}
}
