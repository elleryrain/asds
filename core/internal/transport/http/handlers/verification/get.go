package verification

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/convert"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/pagination"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1VerificationsRequestIDGet(ctx context.Context, params api.V1VerificationsRequestIDGetParams) (api.V1VerificationsRequestIDGetRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	request, err := h.verificationService.GetByID(ctx, params.RequestID)
	if err != nil {
		return nil, fmt.Errorf("get by id: %w", err)
	}

	return &api.V1VerificationsRequestIDGetOK{
		Data: models.VerificationRequestModelToApi(request),
	}, nil
}

func (h *Handler) V1VerificationsOrganizationsOrganizationIDGet(ctx context.Context, params api.V1VerificationsOrganizationsOrganizationIDGetParams) (api.V1VerificationsOrganizationsOrganizationIDGetRes, error) {
	if params.OrganizationID != contextor.GetOrganizationID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to get verification requests of the organization", nil)
	}

	requests, err := h.verificationService.Get(ctx, service.VerificationRequestsObjectGetParams{
		ObjectType: models.ObjectTypeOrganization,
		ObjectID:   models.NewOptional(params.OrganizationID)})
	if err != nil {
		return nil, fmt.Errorf("get verif requests: %w", err)
	}

	return &api.V1VerificationsOrganizationsOrganizationIDGetOK{
		Data: convert.Slice[[]models.VerificationRequest[models.VerificationObject], []api.VerificationRequest](requests.VerificationRequests, models.VerificationRequestModelToApi),
		// В будущем может нужно будет
		// Pagination: pagination.ConvertPaginationToAPI(request.Pagination),
	}, nil
}

func (h *Handler) V1VerificationsOrganizationsGet(ctx context.Context, params api.V1VerificationsOrganizationsGetParams) (api.V1VerificationsOrganizationsGetRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	requests, err := h.verificationService.Get(ctx, service.VerificationRequestsObjectGetParams{
		ObjectType: models.ObjectTypeOrganization,
		Status:     convert.Slice[[]api.VerificationStatus, []models.VerificationStatus](params.Status, models.APIToVerificationStatus),
		Page:       uint64(params.Page.Or(pagination.Page)),
		PerPage:    uint64(params.PerPage.Or(pagination.PerPage))})
	if err != nil {
		return nil, fmt.Errorf("get organization verif req: %w", err)
	}

	return &api.V1VerificationsOrganizationsGetOK{
		Data:       convert.Slice[[]models.VerificationRequest[models.VerificationObject], []api.VerificationRequest](requests.VerificationRequests, models.VerificationRequestModelToApi),
		Pagination: pagination.ConvertPaginationToAPI(requests.Pagination),
	}, nil
}

func (h *Handler) V1VerificationsQuestionAnswerGet(ctx context.Context, params api.V1VerificationsQuestionAnswerGetParams) (api.V1VerificationsQuestionAnswerGetRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	requests, err := h.verificationService.Get(ctx, service.VerificationRequestsObjectGetParams{
		ObjectType: models.ObjectTypeQuestionAnswer,
		Status:     convert.Slice[[]api.VerificationStatus, []models.VerificationStatus](params.Status, models.APIToVerificationStatus),
		Page:       uint64(params.Page.Or(pagination.Page)),
		PerPage:    uint64(params.PerPage.Or(pagination.PerPage))})
	if err != nil {
		return nil, fmt.Errorf("get organization verif req: %w", err)
	}

	return &api.V1VerificationsQuestionAnswerGetOK{
		Data:       convert.Slice[[]models.VerificationRequest[models.VerificationObject], []api.VerificationRequest](requests.VerificationRequests, models.VerificationRequestModelToApi),
		Pagination: pagination.ConvertPaginationToAPI(requests.Pagination),
	}, nil
}

func (h *Handler) V1VerificationsAdditionsGet(ctx context.Context, params api.V1VerificationsAdditionsGetParams) (api.V1VerificationsAdditionsGetRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	requests, err := h.verificationService.Get(ctx, service.VerificationRequestsObjectGetParams{
		ObjectType: models.ObjectTypeAddition,
		Status:     convert.Slice[[]api.VerificationStatus, []models.VerificationStatus](params.Status, models.APIToVerificationStatus),
		Page:       uint64(params.Page.Or(pagination.Page)),
		PerPage:    uint64(params.PerPage.Or(pagination.PerPage))})
	if err != nil {
		return nil, fmt.Errorf("get organization verif req: %w", err)
	}

	return &api.V1VerificationsAdditionsGetOK{
		Data:       convert.Slice[[]models.VerificationRequest[models.VerificationObject], []api.VerificationRequest](requests.VerificationRequests, models.VerificationRequestModelToApi),
		Pagination: pagination.ConvertPaginationToAPI(requests.Pagination),
	}, nil
}

func (h *Handler) V1VerificationsTendersGet(ctx context.Context, params api.V1VerificationsTendersGetParams) (api.V1VerificationsTendersGetRes, error) {
	if contextor.GetRole(ctx) < models.UserRoleEmployee {
		return nil, cerr.ErrPermission
	}

	requests, err := h.verificationService.Get(ctx, service.VerificationRequestsObjectGetParams{
		ObjectType: models.ObjectTypeTender,
		Status:     convert.Slice[[]api.VerificationStatus, []models.VerificationStatus](params.Status, models.APIToVerificationStatus),
		Page:       uint64(params.Page.Or(pagination.Page)),
		PerPage:    uint64(params.PerPage.Or(pagination.PerPage))})
	if err != nil {
		return nil, fmt.Errorf("get organization verif req: %w", err)
	}

	return &api.V1VerificationsTendersGetOK{
		Data:       convert.Slice[[]models.VerificationRequest[models.VerificationObject], []api.VerificationRequest](requests.VerificationRequests, models.VerificationRequestModelToApi),
		Pagination: pagination.ConvertPaginationToAPI(requests.Pagination),
	}, nil
}
