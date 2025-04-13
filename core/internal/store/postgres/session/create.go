package session

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *SessionStore) Create(ctx context.Context, qe store.QueryExecutor, params store.SessionCreateParams) (models.Session, error) {
	builder := squirrel.
		Insert("sessions").
		Columns(
			"id",
			"user_id",
			"ip_address",
			"expires_at",
		).
		Values(
			params.ID,
			params.UserID,
			params.IPAddress,
			params.ExpiresAt,
		).
		Suffix(`
			RETURNING
				id,
				user_id,
				ip_address,
				created_at,
				expires_at
		`).
		PlaceholderFormat(squirrel.Dollar)

	var createdSession models.Session

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&createdSession.ID,
		&createdSession.UserID,
		&createdSession.IPAddress,
		&createdSession.CreatedAt,
		&createdSession.ExpiresAt,
	)
	if err != nil {
		return models.Session{}, fmt.Errorf("query row: %w", err)
	}

	return createdSession, nil
}
