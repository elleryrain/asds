package employee

import (
	"context"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
)

func (h *Handler) V1EmployeePost(ctx context.Context, req *api.V1EmployeePostReq) (api.V1EmployeePostRes, error) {
	createdEmployee, err := h.svc.CreateEmployee(ctx, service.UserCreateEmployeeParams{
		Email:      string(req.GetEmail()),
		Phone:      string(req.GetPhone()),
		Password:   string(req.GetPassword()),
		FirstName:  string(req.GetFirstName()),
		LastName:   string(req.GetLastName()),
		MiddleName: models.Optional[string]{Value: string(req.MiddleName.Value), Set: req.MiddleName.Set},
		Role:       models.APIRoleToModel(req.GetRole()),
		Position:   req.GetPosition(),
	})
	if err != nil {
		return nil, fmt.Errorf("create employee user: %w", err)
	}

	return &api.V1EmployeePostCreated{
		Data: models.ConvertEmployeeUserModelToApi(createdEmployee),
	}, nil
}
