package tender

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *TenderStore) Count(ctx context.Context, qe store.QueryExecutor, params store.TenderGetCountParams) (int, error) {
	builder := squirrel.Select("COUNT(*)").
		From("tenders AS t").
		PlaceholderFormat(squirrel.Dollar)

	if params.OrganizationID.Set {
		builder = builder.Where(squirrel.Eq{"t.organization_id": params.OrganizationID.Value})
	}

	if params.TenderIDs.Set {
		builder = builder.Where(squirrel.Eq{"t.id": params.TenderIDs.Value})
	}

	if !params.WithDrafts {
		builder = builder.Where(squirrel.Eq{"t.is_draft": false})
	}

	if params.VerifiedOnly {
		builder = builder.Where(squirrel.Eq{"t.verification_status": models.VerificationStatusApproved})
	}

	var count int
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&count); err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return count, nil
}
