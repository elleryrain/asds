package tender

import (
	"context"
	"fmt"

	"github.com/Masterminds/squirrel"
	"github.com/lib/pq"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *TenderStore) Update(ctx context.Context, qe store.QueryExecutor, params store.TenderUpdateParams) (int, error) {
	builder := squirrel.Update("tenders")

	if params.Name.Set {
		builder = builder.Set("name", params.Name.Value)
	}
	if params.ServiceIDs.Set {
		builder = builder.Set("services_ids", pq.Array(params.ServiceIDs.Value))
	}
	if params.ObjectIDs.Set {
		builder = builder.Set("objects_ids", pq.Array(params.ObjectIDs.Value))
	}
	if params.Price.Set {
		builder = builder.Set("price", params.Price.Value)
	}
	if params.IsContractPrice.Set {
		builder = builder.Set("is_contract_price", params.IsContractPrice.Value)
	}
	if params.IsNDSPrice.Set {
		builder = builder.Set("is_nds_price", params.IsNDSPrice.Value)
	}
	if params.IsDraft.Set {
		builder = builder.Set("is_draft", params.IsDraft.Value)

		if params.IsDraft.Value {
			builder = builder.Set("status", models.DraftStatus)
		}
	}
	if params.CityID.Set {
		builder = builder.Set("city_id", params.CityID.Value)
	}
	if params.FloorSpace.Set {
		builder = builder.Set("floor_space", params.FloorSpace.Value)
	}
	if params.Description.Set {
		builder = builder.Set("description", params.Description.Value)
	}
	if params.Wishes.Set {
		builder = builder.Set("wishes", params.Wishes.Value)
	}
	if params.Specification.Set {
		builder = builder.Set("specification", params.Specification.Value)
	}
	if params.Attachments.Set {
		builder = builder.Set("attachments", pq.Array(params.Attachments.Value))
	}
	if params.ReceptionStart.Set {
		builder = builder.Set("reception_start", params.ReceptionStart.Value)
	}
	if params.ReceptionEnd.Set {
		builder = builder.Set("reception_end", params.ReceptionEnd.Value)
	}
	if params.WorkStart.Set {
		builder = builder.Set("work_start", params.WorkStart.Value)
	}
	if params.WorkEnd.Set {
		builder = builder.Set("work_end", params.WorkEnd.Value)
	}

	builder = builder.
		Where(squirrel.Eq{"id": params.ID}).
		Suffix(`
			RETURNING id
		`).
		PlaceholderFormat(squirrel.Dollar)

	var id int

	err := builder.RunWith(qe).QueryRowContext(ctx).Scan(&id)
	if err != nil {
		return 0, fmt.Errorf("query row: %w", err)
	}

	return id, nil
}

func (s *TenderStore) UpdateVerificationStatus(ctx context.Context, qe store.QueryExecutor, params store.TenderUpdateVerifStatusParams) error {
	builder := squirrel.Update("tenders").
		Set("verification_status", params.VerificationStatus).
		Set("updated_at", squirrel.Expr("CURRENT_TIMESTAMP")).
		Set("status", params.Status).
		Where(squirrel.Eq{"id": params.TenderID}).
		PlaceholderFormat(squirrel.Dollar)

	result, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("exec row: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")
	}

	return nil
}

func (s *TenderStore) UpdateStatus(ctx context.Context, qe store.QueryExecutor, params store.TenderUpdateStatusParams) error {
	builder := squirrel.Update("tenders").
		Where(squirrel.Eq{"id": params.TenderID}).
		PlaceholderFormat(squirrel.Dollar)

	if params.Status.Set {
		builder = builder.Set("status", params.Status.Value)
	}

	result, err := builder.RunWith(qe).ExecContext(ctx)
	if err != nil {
		return fmt.Errorf("exec row: %w", err)
	}

	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return fmt.Errorf("get affected rows: %w", err)
	}

	if rowsAffected == 0 {
		return fmt.Errorf("no rows were updated")
	}

	return nil
}
