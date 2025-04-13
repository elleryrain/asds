package user

import (
	"context"
	"errors"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/crypto"
	modelsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/models/v1"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"google.golang.org/protobuf/proto"
)

func (s *Service) ReqResetPassword(ctx context.Context, email string) error {
	users, err := s.userStore.Get(ctx, s.psql.DB(), store.UserGetParams{Email: email})
	if err != nil {
		return fmt.Errorf("store get user: %w", err)
	}

	if len(users) == 0 {
		cerr.Wrap(
			errors.New("user not found"),
			cerr.CodeNotFound,
			fmt.Sprintf("user with %s email not found", email),
			nil,
		)
	}

	user := users[0]

	code, err := crypto.GenerateTOTPCode(user.TOTPSalt)
	if err != nil {
		return fmt.Errorf("generate topt: %w", err)
	}

	resetPb, err := proto.Marshal(&modelsv1.PasswordRecovery{
		Email: user.Email,
		Salt:  code,
		Name:  user.FirstName,
	})
	if err != nil {
		return fmt.Errorf("marshal proto: %w", err)
	}

	return s.broker.Publish(ctx, broker.UbratoUserEmailResetPass, resetPb)
}

type ResetPasswordParams struct {
	UserID   int
	Code     string
	Password string
}

func (s *Service) ConfirmResetPassword(ctx context.Context, params ResetPasswordParams) error {
	users, err := s.userStore.Get(ctx, s.psql.DB(), store.UserGetParams{ID: params.UserID})
	if err != nil {
		return fmt.Errorf("store get user: %w", err)
	}

	if len(users) == 0 {
		cerr.Wrap(
			errors.New("user not found"),
			cerr.CodeNotFound,
			fmt.Sprintf("user with %d id not found", params.UserID),
			nil,
		)
	}

	user := users[0]

	if err := crypto.ValidateTOTP(params.Code, user.TOTPSalt); err != nil {
		return fmt.Errorf("validate totp: %w", err)
	}

	hashedPassword, err := crypto.Password(params.Password)
	if err != nil {
		return fmt.Errorf("hash password: %w", err)
	}

	if err := s.userStore.ResetPassword(ctx, s.psql.DB(), store.ResetPasswordParams{
		UserID:       params.UserID,
		PasswordHash: hashedPassword,
	}); err != nil {
		return fmt.Errorf("reset password: %w", err)
	}

	return nil
}
