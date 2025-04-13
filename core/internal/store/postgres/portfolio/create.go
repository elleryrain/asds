package portfolio

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (s *PortfolioStore) Create(ctx context.Context, qe store.QueryExecutor, params store.PortfolioCreateParams) (models.Portfolio, error) {
	builder := squirrel.Insert("portfolios").Columns(
		"organization_id",
		"title",
		"description",
		"attachments",
	).Values(
		params.OrganizationID,
		params.Title,
		params.Description,
		pq.Array(params.Attachments),
	).
		Suffix(`
			RETURNING 
				id, 
				organization_id, 
				title, 
				description, 
				attachments,
				created_at,
				updated_at
		`).
		PlaceholderFormat(squirrel.Dollar)

	var (
		portfolio models.Portfolio
		updatedAt sql.NullTime
	)

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(
		&portfolio.ID,
		&portfolio.OrganizationID,
		&portfolio.Title,
		&portfolio.Description,
		pq.Array(&portfolio.Attachments),
		&portfolio.CreatedAt,
		&updatedAt,
	)
	if err != nil {
		var pqErr *pq.Error
		if errors.As(err, &pqErr) {
			if pqErr.Code == errstore.FKViolation {
				return models.Portfolio{}, errstore.ErrOrganizationNotFound
			}
		}
		return models.Portfolio{}, fmt.Errorf("query row: %w", err)
	}

	portfolio.UpdatedAt = models.Optional[time.Time]{Value: updatedAt.Time, Set: updatedAt.Valid}

	return portfolio, nil
}
