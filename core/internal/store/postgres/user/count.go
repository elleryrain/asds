package user

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *UserStore) CountUsers(ctx context.Context, qe store.QueryExecutor) (int, error) {
	builder := squirrel.Select("COUNT(*)").
		From("users").
		PlaceholderFormat(squirrel.Dollar)

	var count int
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&count); err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return count, nil
}

func (s *UserStore) CountEmployee(ctx context.Context, qe store.QueryExecutor, role models.Optional[[]models.UserRole]) (int, error) {
	builder := squirrel.Select("COUNT(*)").
		From("employee AS e").
		PlaceholderFormat(squirrel.Dollar)

	if role.Set {
		builder = builder.Where(squirrel.Eq{"e.role": role.Value})
	}

	var count int
	if err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&count); err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return count, nil
}
