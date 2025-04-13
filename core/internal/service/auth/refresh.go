package auth

import (
	"context"
	"fmt"
	"time"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/token"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Refresh(ctx context.Context, sessionToken string) (SignInResult, error) {
	session, err := s.sessionStore.Get(ctx, s.psql.DB(), store.SessionGetParams{ID: sessionToken})
	if err != nil {
		return SignInResult{}, fmt.Errorf("get session: %w", err)
	}

	users, err := s.userStore.Get(ctx, s.psql.DB(), store.UserGetParams{ID: session.UserID})
	if err != nil {
		return SignInResult{}, fmt.Errorf("get user: %w", err)
	}

	if len(users) == 0 {
		return SignInResult{}, cerr.Wrap(
			fmt.Errorf("user not found"),
			cerr.CodeNotFound,
			fmt.Sprintf("user with %d user_id not found", session.UserID),
			nil,
		)
	}

	user := users[0]

	rawToken, err := s.tokenAuthorizer.GenerateToken(token.Payload{
		UserID:         user.ID,
		OrganizationID: user.Organization.ID,
		Role:           int(user.Role),
	})
	if err != nil {
		return SignInResult{}, fmt.Errorf("generate access token: %w", err)
	}

	session, err = s.sessionStore.Update(ctx, s.psql.DB(), store.SessionUpdateParams{
		ID:        session.ID,
		ExpiresAt: time.Now().Add(s.tokenAuthorizer.GetRefreshTokenDurationLifetime()),
	})
	if err != nil {
		return SignInResult{}, fmt.Errorf("update session: %w", err)
	}

	return SignInResult{
		User:        user.RegularUser,
		Session:     session,
		AccessToken: rawToken,
	}, nil
}
