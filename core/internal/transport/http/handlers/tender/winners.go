package tender

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1TendersTenderIDWinnersPost(ctx context.Context,
	params api.V1TendersTenderIDWinnersPostParams) (
	api.V1TendersTenderIDWinnersPostRes, error) {
	winner, err := h.winnersStore.Create(ctx, service.WinnersCreateParams{
		TenderID:       params.TenderID,
		OrganizationID: params.OrganizationID,
	})
	if err != nil {
		return nil, fmt.Errorf("winners: %w", err)
	}

	return &api.V1TendersTenderIDWinnersPostCreated{Data: models.ConvertWinnerModelToApi(winner)}, nil
}

func (h *Handler) V1TendersTenderIDWinnersGet(
	ctx context.Context,
	params api.V1TendersTenderIDWinnersGetParams) (
	api.V1TendersTenderIDWinnersGetRes, error) {
	winners, err := h.winnersStore.Get(ctx, params.TenderID)
	if err != nil {
		return nil, fmt.Errorf("get winners: %w", err)
	}

	return &api.V1TendersTenderIDWinnersGetOK{
		Data: convert.Slice[[]models.Winners, []api.Winners](winners, models.ConvertWinnerModelToApi),
	}, nil
}

func (h *Handler) V1TendersWinnersWinnerIDAprovePost(
	ctx context.Context,
	params api.V1TendersWinnersWinnerIDAprovePostParams) (
	api.V1TendersWinnersWinnerIDAprovePostRes, error) {
	if err := h.winnersStore.UpdateStatus(ctx, service.WinnerUpdateParams{
		WinnerID: params.WinnerID,
		Accepted: models.AcceptedStatusApproved,
	}); err != nil {
		return nil, err
	}

	return &api.V1TendersWinnersWinnerIDAprovePostOK{}, nil
}

func (h *Handler) V1TendersWinnersWinnerIDDenyPost(
	ctx context.Context,
	params api.V1TendersWinnersWinnerIDDenyPostParams) (
	api.V1TendersWinnersWinnerIDDenyPostRes, error) {
	if err := h.winnersStore.UpdateStatus(ctx, service.WinnerUpdateParams{
		WinnerID: params.WinnerID,
		Accepted: models.AcceptedStatusDeclined,
	}); err != nil {
		return nil, err
	}

	return &api.V1TendersWinnersWinnerIDDenyPostOK{}, nil
}
