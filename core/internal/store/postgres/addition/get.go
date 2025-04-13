package addition

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *AdditionStore) GetByID(ctx context.Context, qe store.QueryExecutor, id int) (models.Addition, error) {
	builder := squirrel.
		Select(
			"a.id",
			"a.tender_id",
			"a.title",
			"a.content",
			"a.attachments",
			"a.verification_status",
			"a.created_at",
		).
		From("additions AS a").
		Where(squirrel.Eq{"a.id": id}).
		PlaceholderFormat(squirrel.Dollar)

	var addition models.Addition
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&addition.ID,
		&addition.TenderID,
		&addition.Title,
		&addition.Content,
		pq.Array(&addition.Attachments),
		&addition.VerificationStatus,
		&addition.CreatedAt); err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return models.Addition{}, errors.New("addition not found")
		}
		return models.Addition{}, fmt.Errorf("query row: %w", err)
	}

	return addition, nil
}

func (s *AdditionStore) Get(ctx context.Context, qe store.QueryExecutor, params store.AdditionGetParams) ([]models.Addition, error) {
	builder := squirrel.
		Select(
			"a.id",
			"a.tender_id",
			"a.title",
			"a.content",
			"a.attachments",
			"a.verification_status",
			"a.created_at").
		From("additions AS a").
		PlaceholderFormat(squirrel.Dollar)

	if params.TenderID.Set {
		builder = builder.Where(squirrel.Eq{"a.tender_id": params.TenderID.Value})
	}

	if params.VerifiedOnly {
		builder = builder.Where(squirrel.Eq{"a.verification_status": models.VerificationStatusApproved})
	}

	if params.AdditionIDs.Set {
		builder = builder.Where(squirrel.Eq{"a.id": params.AdditionIDs.Value})
	}

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("exec query: %w", err)
	}
	defer rows.Close()

	var additions []models.Addition
	for rows.Next() {
		var addition models.Addition
		if err := rows.Scan(
			&addition.ID,
			&addition.TenderID,
			&addition.Title,
			&addition.Content,
			pq.Array(&addition.Attachments),
			&addition.VerificationStatus,
			&addition.CreatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		additions = append(additions, addition)
	}

	if err = rows.Err(); err != nil {
		return nil, fmt.Errorf("rows error: %w", err)
	}

	return additions, nil
}
