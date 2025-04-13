package store

import "time"

type SessionCreateParams struct {
	ID        string
	UserID    int
	IPAddress string
	ExpiresAt time.Time
}

type SessionGetParams struct {
	ID string
}

type SessionUpdateParams struct {
	ID        string
	ExpiresAt time.Time
}
