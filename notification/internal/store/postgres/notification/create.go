package notificationStore

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/notification/internal/models"
	"gitlab.ubrato.ru/ubrato/notification/internal/store"
)

func (s *NotificationStore) Create(ctx context.Context, psql store.Querier, params store.NotifictionCreateParams) (models.Notification, error) {
	builder := squirrel.Insert("notifications").
		Columns(
			"user_id",
			"title",
			"comment",
			"action_button_text",
			"action_button_url",
			"action_button_styled",
			"status",
			"status_text").
		Values(
			params.UserID,
			params.Title,
			sql.NullString{String: params.Comment.Value, Valid: params.Comment.Set},
			sql.NullString{String: params.ActionButton.Value.Text, Valid: params.ActionButton.Set},
			sql.NullString{String: params.ActionButton.Value.Url, Valid: params.ActionButton.Set},
			sql.NullBool{Bool: params.ActionButton.Value.Styled, Valid: params.ActionButton.Set},
			sql.NullInt16{Int16: int16(params.StatusBlock.Value.Status), Valid: params.StatusBlock.Set},
			sql.NullString{String: params.StatusBlock.Value.Text, Valid: params.StatusBlock.Set},
		).
		Suffix(`
            RETURNING 
                id, 
                user_id, 
                title, 
                comment, 
                action_button_text, 
                action_button_url, 
                action_button_styled, 
                status, 
                status_text, 
                is_read
        `).
		PlaceholderFormat(squirrel.Dollar)

	query, args, err := builder.ToSql()
	if err != nil {
		return models.Notification{}, fmt.Errorf("builder to query: %w", err)
	}

	var (
		notification       models.Notification
		comment            sql.NullString
		actionButtonURL    sql.NullString
		actionButtonText   sql.NullString
		actionButtonStyled sql.NullBool
		statusBlockStatus  sql.NullInt32
		statusBlockText    sql.NullString
	)

	err = psql.QueryRow(ctx, query, args...).Scan(
		&notification.ID,
		&notification.UserID,
		&notification.Title,
		&comment,
		&actionButtonText,
		&actionButtonURL,
		&actionButtonStyled,
		&statusBlockStatus,
		&statusBlockText,
		&notification.IsRead)
	if err != nil {
		return models.Notification{}, fmt.Errorf("query row: %w", err)
	}

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

	return notification, nil
}
