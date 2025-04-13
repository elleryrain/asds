package models

import (
	"time"
)

type Session struct {
	ID        string
	UserID    int
	IPAddress string
	CreatedAt time.Time
	ExpiresAt time.Time
}
