package user

import (
	"context"

	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Update(ctx context.Context, params service.UserUpdateParams) error {
	return s.userStore.Update(ctx, s.psql.DB(), store.UserUpdateParams{
		UserID: params.UserID,
		Phone: params.Phone,
		FirstName: params.FirstName,
		LastName: params.LastName,
		MiddleName: params.MiddleName,
		AvatarURL: params.AvatarURL,
	})
}
