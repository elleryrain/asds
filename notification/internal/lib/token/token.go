package token

import (
	"errors"
	"fmt"

	"github.com/go-chi/jwtauth/v5"
	"github.com/golang-jwt/jwt/v5"
	"gitlab.ubrato.ru/ubrato/notification/internal/config"
)

// TokenAuth нужен для chi jwt middleware
var TokenAuth *jwtauth.JWTAuth

var (
	errEmptySecret = errors.New("authorizer empty JWT secret not allowed")
)

type Claims struct {
	jwt.RegisteredClaims
	Payload
}

type TokenAuthorizer struct {
	cfg config.JWT
}

func NewTokenAuthorizer(settings config.JWT) (*TokenAuthorizer, error) {
	if settings.Secret == "" {
		return nil, errEmptySecret
	}

	TokenAuth = jwtauth.New("HS256", []byte(settings.Secret), nil)

	return &TokenAuthorizer{cfg: settings}, nil
}

func (ta *TokenAuthorizer) ValidateToken(rawToken string) (Claims, error) {
	var claims Claims
	_, err := jwt.ParseWithClaims(rawToken, &claims, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(ta.cfg.Secret), nil
	})
	if err != nil {
		return Claims{}, fmt.Errorf("validate token: %w", err)
	}

	return claims, nil
}
