package user

import (
	"context"
	"errors"
	"fmt"

	api "gitlab.ubrato.ru/ubrato/core/api/gen"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (h *Handler) V1UsersUserIDPut(ctx context.Context, req *api.V1UsersUserIDPutReq, params api.V1UsersUserIDPutParams) (api.V1UsersUserIDPutRes, error) {
	if params.UserID != contextor.GetUserID(ctx) {
		return nil, cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to edit the user", nil)
	}

	if err := h.svc.Update(ctx, service.UserUpdateParams{
		UserID:     params.UserID,
		Phone:      models.Optional[string]{Value: string(req.Phone.Value), Set: req.Phone.Set},
		FirstName:  models.Optional[string]{Value: string(req.FirstName.Value), Set: req.FirstName.Set},
		LastName:   models.Optional[string]{Value: string(req.LastName.Value), Set: req.LastName.Set},
		MiddleName: models.Optional[string]{Value: string(req.MiddleName.Value), Set: req.MiddleName.Set},
		AvatarURL:  models.Optional[string]{Value: string(req.AvatarURL.Value.String()), Set: req.AvatarURL.Set},
	}); err != nil {
		if errors.Is(err, errstore.ErrUserNotFound) {
			return nil, cerr.Wrap(err, cerr.CodeNotFound, "Пользователь не найден", map[string]interface{}{
				"user_id": params.UserID,
			})
		}

		return nil, fmt.Errorf("update user: %w", err)
	}

	return &api.V1UsersUserIDPutOK{}, nil
}
