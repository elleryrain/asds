package questionnaire

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
	"gitlab.ubrato.ru/ubrato/core/internal/service"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
)

func (s *Service) Create(ctx context.Context, params service.QuestionnaireCreateParams) error {
	org, err := s.organizationStore.GetByID(ctx, s.psql.DB(), params.OrganizationID)
	if err != nil {
		return fmt.Errorf("get organization: %w", err)
	}

	if !org.IsContractor {
		return cerr.Wrap(
			fmt.Errorf("user is not a contractor and cannot access this questionnaire"),
			cerr.CodeNotPermitted,
			"User is not a contractor",
			map[string]interface{}{"organization_id": params.OrganizationID},
		)
	}

	return s.questionnaireStore.Create(ctx, s.psql.DB(), store.QuestionnaireCreateParams{
		OrganizationID: params.OrganizationID,
		Answers:        params.Answers,
		IsCompleted:    params.IsCompleted,
	})
}
