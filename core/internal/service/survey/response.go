package survey

import (
	"context"
	"fmt"

	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	eventsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/events/v1"
	modelsv1 "gitlab.ubrato.ru/ubrato/core/internal/models/gen/proto/models/v1"
	"google.golang.org/protobuf/proto"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type ResponseParams struct {
	Name     string
	Type     string
	Phone    string
	Question string
}

func (s *Service) Response(ctx context.Context, params ResponseParams) error {
	b, err := proto.Marshal(&eventsv1.SurveySubmitted{
		Context: &modelsv1.EventContext{
			Timestamp: timestamppb.Now(),
		},
		Survey: &modelsv1.Survey{
			Name:     params.Name,
			Type:     modelsv1.SurveyType(modelsv1.SurveyType_value[params.Type]),
			Phone:    params.Phone,
			Question: params.Question,
		},
	})
	if err != nil {
		return fmt.Errorf("marshal proto: %w", err)
	}

	return s.broker.Publish(ctx, broker.UbratoSurveySubmittedSubject, b)
}
