package notificationStore

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/notification/internal/models"
	"gitlab.ubrato.ru/ubrato/notification/internal/store"
	"gitlab.ubrato.ru/ubrato/notification/internal/store/errstore"
)

func (s *NotificationStore) Get(ctx context.Context, psql store.Querier, params store.NotifictionGetParams) ([]models.Notification, error) {
	builder := squirrel.Select(
		"id",
		"user_id",
		"title",
		"comment",
		"action_button_text",
		"action_button_url",
		"action_button_styled",
		"status",
		"status_text",
		"is_read",
	).
		From("notifications").
		OrderBy("created_at DESC").
		PlaceholderFormat(squirrel.Dollar)

	if params.NotificationID.Set {
		builder = builder.Where(squirrel.Eq{"id": params.NotificationID.Value})
	}

	if params.UserID.Set {
		builder = builder.Where(squirrel.Eq{"user_id": params.UserID.Value})
	}

	query, args, err := builder.ToSql()
	if err != nil {
		return nil, fmt.Errorf("builder to query: %w", err)
	}

	rows, err := psql.Query(ctx, query, args...)
	if err != nil {
		return nil, fmt.Errorf("query: %w", err)
	}
	defer rows.Close()

	var notifications []models.Notification
	for rows.Next() {
		var (
			notification       models.Notification
			comment            sql.NullString
			actionButtonText   sql.NullString
			actionButtonURL    sql.NullString
			actionButtonStyled sql.NullBool
			statusBlockStatus  sql.NullInt32
			statusBlockText    sql.NullString
		)

		if err := rows.Scan(
			&notification.ID,
			&notification.UserID,
			&notification.Title,
			&comment,
			&actionButtonText,
			&actionButtonURL,
			&actionButtonStyled,
			&statusBlockStatus,
			&statusBlockText,
			&notification.IsRead,
		); err != nil {
			return nil, fmt.Errorf("scan: %w", err)
		}

		notification.Comment = comment.String

		if actionButtonText.Valid {
			notification.ActionButton = &models.ActionButton{
				Text:   actionButtonText.String,
				Url:    actionButtonURL.String,
				Styled: actionButtonStyled.Bool,
			}
		}

		if statusBlockStatus.Valid {
			notification.StatusBlock = &models.StatusBlock{
				Status: models.Status(statusBlockStatus.Int32),
				Text:   statusBlockText.String,
			}
		}

		notifications = append(notifications, notification)
	}

	if err := rows.Err(); err != nil {
		return nil, fmt.Errorf("rows iteration: %w", err)
	}

	return notifications, nil
}

func (s *NotificationStore) GetByID(ctx context.Context, qe store.Querier, notificationID int) (models.Notification, error) {
	notifications, err := s.Get(ctx, qe, store.NotifictionGetParams{
		NotificationID: models.NewOptional(notificationID)})
	if err != nil {
		return models.Notification{}, fmt.Errorf("get notifications: %w", err)
	}

	if len(notifications) == 0 {
		return models.Notification{}, errstore.ErrNotificationNotFound
	}

	return notifications[0], nil
}
