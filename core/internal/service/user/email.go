package user

import (
	"context"
	"errors"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/crypto"
	eventsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/events/v1"
	modelsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/models/v1"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"google.golang.org/protobuf/proto"
)

func (s *Service) ReqEmailVerification(ctx context.Context, email string) error {
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

	confirmPb, err := proto.Marshal(&modelsv1.EmailConfirmation{
		Email: user.Email,
		Salt:  code,
	})
	if err != nil {
		return fmt.Errorf("marshal proto: %w", err)
	}

	return s.broker.Publish(ctx, broker.UbratoUserConfirmEmail, confirmPb)
}

type ConfirmEmailParams struct {
	UserID int
	Code   string
}

func (s *Service) ConfirmEmail(ctx context.Context, params ConfirmEmailParams) error {
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

	if err := s.userStore.SetEmailVerified(ctx, s.psql.DB(), user.ID); err != nil {
		return fmt.Errorf("set email verifed: %w", err)
	}

	// уведомление
	notify, err := proto.Marshal(&eventsv1.SentNotification{
		Notification: &modelsv1.Notification{
			User: &modelsv1.NotifiedUser{
				Id: *proto.Int32(int32(user.ID)),
			},
		}})
	if err != nil {
		return fmt.Errorf("marhal notification proto: %w", err)
	}

	err = s.broker.Publish(ctx, broker.NotifyUserEmailConfirmed, notify)
	if err != nil {
		return fmt.Errorf("notification: %w", err)
	}

	return nil
}
