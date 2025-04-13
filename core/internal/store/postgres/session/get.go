package session

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *SessionStore) Get(ctx context.Context, qe store.QueryExecutor, params store.SessionGetParams) (models.Session, error) {
	builder := squirrel.
		Select(
			"id",
			"user_id",
			"ip_address",
			"created_at",
			"expires_at",
		).
		From("sessions").
		Where(squirrel.Eq{"id": params.ID}).
		PlaceholderFormat(squirrel.Dollar)

	var session models.Session

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&session.ID,
		&session.UserID,
		&session.IPAddress,
		&session.CreatedAt,
		&session.ExpiresAt,
	)
	if err != nil {
		return models.Session{}, fmt.Errorf("query row: %w", err)
	}

	return session, nil
}
