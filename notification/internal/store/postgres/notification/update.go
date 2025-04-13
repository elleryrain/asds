package notificationStore

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/notification/internal/store"
	"gitlab.ubrato.ru/ubrato/notification/internal/store/errstore"
)

func (s *NotificationStore) Update(ctx context.Context, psql store.Querier, params store.NotifictionUpdateParams) error {
	builder := squirrel.Update("notifications").
		Set("is_read", true).
		Where(squirrel.Eq{"id": params.NotificationID}).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return fmt.Errorf("builder to query: %w", err)
	}

	result, err := psql.Exec(ctx, query, args...)
	if err != nil {
		return fmt.Errorf("exec query: %w", err)
	}

	if result.RowsAffected() == 0 {
		return errstore.ErrNotificationNotFound
	}

	return nil
}
