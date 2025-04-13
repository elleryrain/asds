package handler

import (
	"context"
	"log/slog"
	"os"
	"testing"

	"gitlab.ubrato.ru/ubrato/amo-sync/internal/broker"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/broker/jetstream"
	eventsv1 "gitlab.ubrato.ru/ubrato/amo-sync/internal/gen/pb/events/v1"
	v1 "gitlab.ubrato.ru/ubrato/amo-sync/internal/gen/pb/models/v1"
	"google.golang.org/protobuf/proto"
)

var js *jetstream.Broker

func TestMain(m *testing.M) {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	var err error
	js, err = jetstream.New(logger, ":4222")
	if err != nil {
		logger.Error("Init jetstream", "error", err)
		return
	}

	m.Run()
}

func TestSurveySubmitted(t *testing.T) {
	event := &eventsv1.SurveySubmitted{
		Survey: &v1.Survey{
			Name:     "Тестовый опросник",
			Type:     v1.SurveyType_SURVEY_TYPE_FEEDBACK,
			Phone:    "1234567890",
			Question: "Вопрос",
		},
	}

	data, err := proto.Marshal(event)
	if err != nil {
		t.Fatalf("Failed to marshal event: %v", err)
	}

	err = js.Publish(context.Background(), broker.UbratoSurveySubmittedSubject, data)
	if err != nil {
		t.Fatalf("Failed to publish event: %v", err)
	}
}

func TestUserRegistered(t *testing.T) {
	event := &eventsv1.UserRegistered{
		User: &v1.User{
			Id:         1,
			FirstName:  "first name",
			LastName:   "last name",
			MiddleName: "middle name",
			Email:      "test@mail.ru",
			Phone:      "1234567890",
			Organization: &v1.Organization{
				Id:        1,
				ShortName: "short name",
				FullName:  "full name",
				BrandName: "brand name",
				Inn:       "INN",
				Okpo:      "OKPO",
				Ogrn:      "OGRN",
				Kpp:       "KPP",
				TaxCode:   "tax code",
			},
		},
	}

	data, err := proto.Marshal(event)
	if err != nil {
		t.Fatalf("Failed to marshal event: %v", err)
	}

	err = js.Publish(context.Background(), broker.UbratoUserRegisteredSubject, data)
	if err != nil {
		t.Fatalf("Failed to publish event: %v", err)
	}
}
