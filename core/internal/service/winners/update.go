package winners

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) UpdateStatus(ctx context.Context, params service.WinnerUpdateParams) error {
	organizationID, err := s.winnersStore.GetOrganizationIDByWinnerID(ctx, s.psql.DB(), params.WinnerID)
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return cerr.Wrap(err, cerr.CodeNotFound, "winner not found", nil)
		}
		return err
	}

	if organizationID != contextor.GetOrganizationID(ctx) {
		return cerr.Wrap(cerr.ErrPermission, cerr.CodeNotPermitted, "not enough permissions to update the winner accepted status", nil)
	}

	return s.psql.WithTransaction(ctx, func(qe store.QueryExecutor) error {
		var status models.Optional[models.TenderStatus]

		if params.Accepted == models.AcceptedStatusApproved {
			status.Value = models.ContractorSelectedStatus
			status.Set = true
		} else {
			status.Set = false
		}

		tenderID, err := s.winnersStore.GetTenderIDByWinnerID(ctx, s.psql.DB(), params.WinnerID)
		if err != nil {
			return fmt.Errorf("get tender by winner ID: %w", err)
		}

		err = s.tenderStore.UpdateStatus(ctx, s.psql.DB(), store.TenderUpdateStatusParams{
			TenderID: tenderID,
			Status:   status,
		})
		if err != nil {
			return fmt.Errorf("update tender status: %w", err)
		}

		err = s.winnersStore.UpdateStatus(ctx, s.psql.DB(), store.WinnerUpdateParams{
			WinnerID: params.WinnerID,
			Accepted: params.Accepted,
		})
		if err != nil {
			return fmt.Errorf("update winner status: %w", err)
		}

		return nil
	})
}
