package verification

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *VerificationStore) Count(ctx context.Context, qe store.QueryExecutor, params store.VerificationRequestsObjectGetCountParams) (int, error) {
	builder := squirrel.Select("COUNT(*)").
		From("verification_requests AS vr").
		PlaceholderFormat(squirrel.Dollar)

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

	var count int
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&count); err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return count, nil
}
