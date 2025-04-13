package suggest

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/lib/cerr"
)

func (s *Service) Company(ctx context.Context, INN string) (string, error) {
	resp, err := s.dadataGateway.FindByINN(ctx, INN)
	if err != nil {
		return "", err
	}

	if len(resp.Suggestions) == 0 {
		return "", cerr.Wrap(fmt.Errorf("company not found"), cerr.CodeNotFound, "No suggestions for provided INN", nil)
	}

	return resp.Suggestions[0].Data.Name.Short, nil
}
