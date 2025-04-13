package authSrv

import "gitlab.ubrato.ru/ubrato/notification/internal/lib/token"

type Service struct {
	tokenAuthorizer TokenAuthorizer
}

type TokenAuthorizer interface {
	ValidateToken(rawToken string) (token.Claims, error)
}

func New(tokenAuthorizer TokenAuthorizer) *Service {
	return &Service{
		tokenAuthorizer: tokenAuthorizer,
	}
}
