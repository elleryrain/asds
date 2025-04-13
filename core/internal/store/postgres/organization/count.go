package organization

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *OrganizationStore) Count(ctx context.Context, qe store.QueryExecutor, params store.OrganizationGetCountParams) (int, error) {
	builder := squirrel.Select("COUNT(*)").
		From("organizations AS o").
		PlaceholderFormat(squirrel.Dollar)

	if params.IsContractor.Set {
		builder = builder.Where(
			squirrel.Eq{
				"o.is_contractor":       params.IsContractor.Value,
				"o.verification_status": models.VerificationStatusApproved,
				"o.is_banned":           false,
			})
	}

	if len(params.OrganizationIDs) != 0 {
		builder = builder.Where(squirrel.Eq{"o.id": params.OrganizationIDs})
	}

	var count int
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&count); err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return count, nil
}
