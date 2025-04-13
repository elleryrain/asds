package tender

import (
	"context"
	"fmt"
	"net/url"
	"time"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1TendersTenderIDPut(ctx context.Context, req *api.V1TendersTenderIDPutReq, params api.V1TendersTenderIDPutParams) (api.V1TendersTenderIDPutRes, error) {
	tender, err := h.tenderService.Update(ctx, service.TenderUpdateParams{
		ID:              params.TenderID,
		OrganizationID:  contextor.GetOrganizationID(ctx),
		Name:            models.Optional[string]{Value: req.GetName().Value, Set: req.GetName().Set},
		Price:           models.Optional[int]{Value: int(req.GetPrice().Value * 100), Set: req.GetPrice().Set},
		IsContractPrice: models.Optional[bool]{Value: req.GetIsContractPrice().Value, Set: req.GetIsContractPrice().Set},
		IsNDSPrice:      models.Optional[bool]{Value: req.GetIsNdsPrice().Value, Set: req.GetIsNdsPrice().Set},
		IsDraft:         models.Optional[bool]{Value: req.GetIsDraft().Value, Set: req.GetIsDraft().Set},
		CityID:          models.Optional[int]{Value: req.GetCity().Value, Set: req.GetCity().Set},
		FloorSpace:      models.Optional[int]{Value: req.GetFloorSpace().Value, Set: req.GetFloorSpace().Set},
		Description:     models.Optional[string]{Value: req.GetDescription().Value, Set: req.GetDescription().Set},
		Wishes:          models.Optional[string]{Value: req.GetWishes().Value, Set: req.GetWishes().Set},
		Specification:   models.Optional[string]{Value: string(req.Specification.Value.String()), Set: req.GetSpecification().Set},
		Attachments: models.Optional[[]string]{Value: convert.Slice[[]url.URL, []string](
			req.GetAttachments(), func(u url.URL) string { return u.String() },
		), Set: req.GetAttachments() != nil},
		ServiceIDs:     models.Optional[[]int]{Value: req.GetServices(), Set: len(req.GetServices()) > 0},
		ObjectIDs:      models.Optional[[]int]{Value: req.GetObjects(), Set: len(req.GetObjects()) > 0},
		ReceptionStart: models.Optional[time.Time]{Value: req.GetReceptionStart().Value, Set: req.GetReceptionStart().Set},
		ReceptionEnd:   models.Optional[time.Time]{Value: req.GetReceptionEnd().Value, Set: req.GetReceptionEnd().Set},
		WorkStart:      models.Optional[time.Time]{Value: req.GetWorkStart().Value, Set: req.GetWorkStart().Set},
		WorkEnd:        models.Optional[time.Time]{Value: req.GetWorkEnd().Value, Set: req.GetWorkEnd().Set},
	})
	if err != nil {
		return nil, fmt.Errorf("update tender: %w", err)
	}

	return &api.V1TendersTenderIDPutOK{
		Data: models.ConvertTenderModelToApi(tender),
	}, nil
}
