package portfolio

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/errstore"
)

func (s *PortfolioStore) Get(ctx context.Context, qe store.QueryExecutor, params store.PortfolioGetParams) ([]models.Portfolio, error) {
	builder := squirrel.Select(
		"p.id",
		"p.organization_id",
		"p.title",
		"p.description",
		"p.attachments",
		"p.created_at",
		"p.updated_at",
	).From("portfolios AS p").
		PlaceholderFormat(squirrel.Dollar)

	if params.PortfolioID.Set {
		builder = builder.Where(
			squirrel.Eq{"p.id": params.PortfolioID.Value})
	}

	if params.OrganizationID.Set {
		builder = builder.Where(
			squirrel.Eq{"p.organization_id": params.OrganizationID.Value})
	}

	rows, err := builder.RunWith(qe).QueryContext(ctx)
	if err != nil {
		return nil, fmt.Errorf("query row: %w", err)
	}
	defer rows.Close()

	portfolios := []models.Portfolio{}
	for rows.Next() {
		var (
			portfolio models.Portfolio
			updatedAt sql.NullTime
		)

		if err = rows.Scan(
			&portfolio.ID,
			&portfolio.OrganizationID,
			&portfolio.Title,
			&portfolio.Description,
			pq.Array(&portfolio.Attachments),
			&portfolio.CreatedAt,
			&updatedAt,
		); err != nil {
			return nil, fmt.Errorf("scan row: %w", err)
		}

		portfolio.UpdatedAt = models.Optional[time.Time]{Value: updatedAt.Time, Set: updatedAt.Valid}

		portfolios = append(portfolios, portfolio)
	}

	if rows.Err() != nil {
		return nil, fmt.Errorf("row iteration: %w", rows.Err())
	}

	return portfolios, nil
}

func (s *PortfolioStore) GetByID(ctx context.Context, qe store.QueryExecutor, portfolioID int) (models.Portfolio, error) {
	portfolios, err := s.Get(ctx, qe, store.PortfolioGetParams{
		PortfolioID: models.NewOptional(portfolioID)})
	if err != nil {
		return models.Portfolio{}, fmt.Errorf("get portfolios: %w", err)
	}

	if len(portfolios) == 0 {
		return models.Portfolio{}, errstore.ErrPortfolioNotFound
	}

	return portfolios[0], nil
}
