package broker

import (
	"context"
	"log/slog"
	"time"

	"github.com/nats-io/nats.go/jetstream"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/broker"
	brokerJetstream "gitlab.ubrato.ru/ubrato/amo-sync/internal/broker/jetstream"
	eventsv1 "gitlab.ubrato.ru/ubrato/amo-sync/internal/gen/pb/events/v1"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/transport/broker/handler"
	"google.golang.org/protobuf/proto"
)

type Transport struct {
	logger  *slog.Logger
	broker  *brokerJetstream.Broker
	handler *handler.Handler

	amoConsumer jetstream.ConsumeContext
}

func New(
	logger *slog.Logger,
	broker *brokerJetstream.Broker,
	handler *handler.Handler,
) *Transport {
	return &Transport{
		logger:  logger,
		broker:  broker,
		handler: handler,
	}
}

func (t *Transport) Start() error {
	consumer, err := t.broker.NewConsumer()
	if err != nil {
		return err
	}

	t.amoConsumer, err = consumer.Consume(t.handleMsg)
	if err != nil {
		return err
	}

	return nil
}

func (t *Transport) Stop() {
	t.amoConsumer.Stop()
}

func (t *Transport) handleMsg(msg jetstream.Msg) {
	t.logger.Info("Handle msg", "subject", msg.Subject())
	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	var err error

	defer func() {
		if err != nil {
			if nakErr := msg.Nak(); nakErr != nil {
				t.logger.Error("Nak message", "error", err)
			}
			t.logger.Error("Handle", "error", err)
			return
		} else {
			if err = msg.Ack(); err != nil {
				t.logger.Error("Ack message", "error", err)
			}
		}
	}()

	switch msg.Subject() {
	case broker.UbratoUserRegisteredSubject:
		command := new(eventsv1.UserRegistered)
		err = proto.Unmarshal(msg.Data(), command)
		if err != nil {
			t.logger.Error("Unmarshal", "error", err)
			return
		}

		err = t.handler.UserRegistered(ctx, command)
	case broker.UbratoSurveySubmittedSubject:
		command := new(eventsv1.SurveySubmitted)
		err = proto.Unmarshal(msg.Data(), command)
		if err != nil {
			t.logger.Error("Unmarshal", "error", err)
			return
		}

		err = t.handler.SurveySubmitted(ctx, command)
	}
}
