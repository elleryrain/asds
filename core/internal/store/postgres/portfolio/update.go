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

func (s *PortfolioStore) Update(ctx context.Context, qe store.QueryExecutor, params store.PortfolioUpdateParams) (models.Portfolio, error) {
	builder := squirrel.Update("portfolios").
		Set("updated_at", squirrel.Expr("CURRENT_TIMESTAMP")).
		Where(squirrel.Eq{"id": params.PortfolioID}).
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

	if params.Title.Set {
		builder = builder.Set("title", params.Title.Value)
	}

	if params.Description.Set {
		builder = builder.Set("description", params.Description.Value)
	}

	if params.Attachments.Set {
		builder = builder.Set("attachments", pq.Array(params.Attachments.Value))
	}

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

	switch {
	case errors.Is(err, sql.ErrNoRows):
		return models.Portfolio{}, errstore.ErrPortfolioNotFound
	case err != nil:
		return models.Portfolio{}, fmt.Errorf("query row: %w", err)
	}

	portfolio.UpdatedAt = models.Optional[time.Time]{Value: updatedAt.Time, Set: updatedAt.Valid}

	return portfolio, nil
}
