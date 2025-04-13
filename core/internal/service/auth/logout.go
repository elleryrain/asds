package auth

import "context"

func (s *Service) Logout(ctx context.Context, sessionToken string) error {
	return s.sessionStore.Delete(ctx, s.psql.DB(), sessionToken)
}
