package postgres

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/models"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/store"
)

type AmoStore struct {
}

func NewAmoStore() *AmoStore {
	return &AmoStore{}
}

func (s *AmoStore) Create(ctx context.Context, qe store.QueryExecutor, externalID int, amoID int, entity models.Entity) error {
	builder := squirrel.
		Insert("external_id_amo_id").
		Columns(
			"external_id",
			"amo_id",
			"entity",
		).
		Values(
			externalID,
			amoID,
			entity,
		).
		PlaceholderFormat(squirrel.Dollar)

	_, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("amostore create: %w", err)
	}

	return nil
}

func (s *AmoStore) Get(ctx context.Context, qe store.QueryExecutor, externalID int, entity models.Entity) (int, error) {
	var amoID int

	builder := squirrel.
		Select("amo_id").
		From("external_id_amo_id").
		Where(squirrel.Eq{
			"external_id": externalID,
			"entity":      entity,
		}).
		PlaceholderFormat(squirrel.Dollar)

	err := builder.RunWith(qe).QueryRow().Scan(&amoID)
	if err != nil {
		return 0, fmt.Errorf("amostore get: %w", err)
	}

	return amoID, nil
}
