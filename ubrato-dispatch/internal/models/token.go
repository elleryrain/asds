package models

type Token struct {
	AccessToken string
	TokenType   string
	ExpiresIn   int
}
