package models

import (
	"time"
)

type UserRole int32

const (
	UserRoleUser  UserRole = 0
	UserRoleAdmin UserRole = 1
)

type User struct {
	ID            int
	Email         string
	Phone         string
	FirstName     string
	LastName      string
	MiddleName    string
	AvatarURL     string
	Verified      bool
	EmailVerified bool
	Role          UserRole
	IsContractor  bool
	IsBanned      bool
	Organization  Organization
	CreatedAt     time.Time
	UpdatedAt     time.Time
}
