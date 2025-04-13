package verification

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *VerificationStore) GetByIDWithEmptyObject(ctx context.Context, qe store.QueryExecutor, requestID int) (models.VerificationRequest[models.VerificationObject], error) {
	requests, err := s.GetWithEmptyObject(ctx, qe, store.VerificationRequestsObjectGetParams{
		VerificationID: models.NewOptional(requestID),
	})
	if err != nil {
		return models.VerificationRequest[models.VerificationObject]{}, fmt.Errorf("get objects: %w", err)
	}

	if len(requests) == 0 {
		return models.VerificationRequest[models.VerificationObject]{}, errors.New("object not found")
	}

	return requests[0], nil
}

func (s *VerificationStore) GetWithEmptyObject(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestsObjectGetParams) ([]models.VerificationRequest[models.VerificationObject], error) {
	builder := squirrel.Select(
		"vr.id",
		"vr.reviewer_user_id",
		"u.email",
		"u.phone",
		"u.first_name",
		"u.last_name",
		"u.middle_name",
		"u.avatar_url",
		"u.email_verified",
		"u.is_banned",
		"u.created_at AS user_created_at",
		"u.updated_at AS user_updated_at",
		"e.position",
		"e.role",
		"vr.object_type",
		"vr.object_id",
		"vr.content",
		"vr.attachments",
		"vr.status",
		"vr.review_comment",
		"vr.created_at AS vr_created_at",
		"vr.reviewed_at AS vr_reviewed_at").
		From("verification_requests vr").
		LeftJoin("users u ON vr.reviewer_user_id = u.id").
		LeftJoin("employee e ON e.user_id = u.id").
		OrderBy("vr.created_at DESC").
		PlaceholderFormat(squirrel.Dollar)

	if params.Limit.Set {
		builder = builder.Limit(params.Limit.Value)
	}

	if params.Offset.Set {
		builder = builder.Offset(params.Offset.Value)
	}

	if params.ObjectType.Set {
		builder = builder.Where(squirrel.Eq{"vr.object_type": params.ObjectType.Value})
	}

	if params.VerificationID.Set {
		builder = builder.Where(squirrel.Eq{"vr.id": params.VerificationID.Value})
	}

	if params.ObjectID.Set {
		builder = builder.Where(squirrel.Eq{"vr.object_id": params.ObjectID.Value})
	}

	if len(params.Status) != 0 {
		builder = builder.Where(squirrel.Eq{"vr.status": params.Status})
	}

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query context: %w", err)
	}
	defer rows.Close()

	requests := []models.VerificationRequest[models.VerificationObject]{}
	for rows.Next() {
		var (
			request models.VerificationRequest[models.VerificationObject]

			reviewerID            sql.NullInt64
			reviewerEmail         sql.NullString
			reviewerPhone         sql.NullString
			reviewerFirstName     sql.NullString
			reviewerLastName      sql.NullString
			reviewerMiddleName    sql.NullString
			reviewerAvatarURL     sql.NullString
			reviewerEmailVerified sql.NullBool
			reviewerIsBanned      sql.NullBool
			reviewerCreatedAt     sql.NullTime
			reviewerUpdatedAt     sql.NullTime
			reviewerPosition      sql.NullString
			reviewerRole          sql.NullInt16

			requestContent         sql.NullString
			requestReviewerComment sql.NullString
			requestReviewedAt      sql.NullTime
		)

		if err := rows.Scan(
			&request.ID,
			&reviewerID,
			&reviewerEmail,
			&reviewerPhone,
			&reviewerFirstName,
			&reviewerLastName,
			&reviewerMiddleName,
			&reviewerAvatarURL,
			&reviewerEmailVerified,
			&reviewerIsBanned,
			&reviewerCreatedAt,
			&reviewerUpdatedAt,
			&reviewerPosition,
			&reviewerRole,
			&request.ObjectType,
			&request.ObjectID,
			&requestContent,
			&request.Attachments,
			&request.Status,
			&requestReviewerComment,
			&request.CreatedAt,
			&requestReviewedAt); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		request.Content = requestContent.String
		request.ReviewComment = requestReviewerComment.String
		request.ReviewedAt = requestReviewedAt.Time

		if reviewerID.Valid {
			request.Reviewer = models.EmployeeUser{
				User: models.User{
					ID:            int(reviewerID.Int64),
					Email:         reviewerEmail.String,
					Phone:         reviewerPhone.String,
					FirstName:     reviewerFirstName.String,
					LastName:      reviewerLastName.String,
					MiddleName:    models.Optional[string]{Value: reviewerMiddleName.String, Set: reviewerMiddleName.Valid},
					AvatarURL:     reviewerAvatarURL.String,
					EmailVerified: reviewerEmailVerified.Bool,
					IsBanned:      reviewerIsBanned.Bool,
					CreatedAt:     reviewerCreatedAt.Time,
					UpdatedAt:     reviewerUpdatedAt.Time,
				},
				Position: reviewerPosition.String,
				Role:     models.UserRole(reviewerRole.Int16),
			}
		}

		requests = append(requests, request)
	}

	return requests, nil
}
