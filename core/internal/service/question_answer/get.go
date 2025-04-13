package questionanswer

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/contextor"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Get(ctx context.Context, tenderID int) ([]models.QuestionWithAnswer, error) {
	tender, err := s.tenderStore.GetByID(ctx, s.psql.DB(), tenderID)
	if err != nil {
		return nil, fmt.Errorf("get tender by id: %w", err)
	}

	organizationID := contextor.GetOrganizationID(ctx)
	return s.questionAnswerStore.GetWithAccess(ctx, s.psql.DB(), store.QuestionAnswerGetWithAccessParams{
		TenderID:        tenderID,
		VerifedOnly:     organizationID == 0,
		OrganizationID:  models.Optional[int]{Value: organizationID, Set: organizationID != 0},
		IsTenderCreator: tender.Organization.ID == organizationID,
	})
}
