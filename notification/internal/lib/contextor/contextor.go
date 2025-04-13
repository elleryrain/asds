package contextor

import (
	"context"
)

type ctxKey int

const (
	UserIDKey ctxKey = iota
	OrganizationIDKey
)

func GetUserID(ctx context.Context) int {
	v, ok := ctx.Value(UserIDKey).(int)
	if !ok {
		return 0
	}

	return v
}

func GetOrganizationID(ctx context.Context) int {
	v, ok := ctx.Value(OrganizationIDKey).(int)
	if !ok {
		return 0
	}

	return v
}
