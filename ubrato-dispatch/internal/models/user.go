package models

import "time"

type User struct {
	ID           int
	Email        string
	Phone        string
	PasswordHash string
	TotpSalt     string
	FirstName    string
	MiddleName   string
	LastName     string
	Avatar       string
	Verified     bool
	Role         int16
	IsContractor bool
	CreatedAt    time.Time
}
