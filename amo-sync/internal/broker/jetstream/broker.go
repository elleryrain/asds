package jetstream

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/broker"
)

type Broker struct {
	logger          *slog.Logger
	js              jetstream.JetStream
	ubratoStream    jetstream.Stream
	ubratoDLQStream jetstream.Stream
}

func New(logger *slog.Logger, addr string) (*Broker, error) {
	nc, err := nats.Connect(addr)
	if err != nil {
		return nil, fmt.Errorf("nats connect: %w", err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		return nil, fmt.Errorf("create jetstream: %w", err)
	}

	ubratoStream, err := js.CreateOrUpdateStream(context.Background(), jetstream.StreamConfig{
		Name:        "ubrato",
		Description: "Ubrato stream",
		Subjects: []string{
			broker.UbratoUserRegisteredSubject,
			broker.UbratoSurveySubmittedSubject,
		},
		MaxMsgs: 100000,
	})
	if err != nil {
		return nil, fmt.Errorf("create stream: %w", err)
	}

	ubratoDLQStream, err := js.CreateOrUpdateStream(context.Background(), jetstream.StreamConfig{
		Name:        "ubrato_dlq",
		Description: "Dead letter queue for jobs stream",
		Subjects: []string{
			"$JS.EVENT.ADVISORY.CONSUMER.MAX_DELIVERIES.ubrato.*",
		},
		MaxMsgs: 50000,
	})
	if err != nil {
		return nil, fmt.Errorf("create stream: %w", err)
	}

	return &Broker{
		logger:          logger,
		js:              js,
		ubratoStream:    ubratoStream,
		ubratoDLQStream: ubratoDLQStream,
	}, nil
}

func (b *Broker) Publish(ctx context.Context, subject string, data []byte) error {
	_, err := b.js.Publish(ctx, subject, data)
	return err
}

func (b *Broker) NewConsumer() (jetstream.Consumer, error) {
	consumer, err := b.ubratoStream.CreateOrUpdateConsumer(context.Background(), jetstream.ConsumerConfig{
		Name:        "amo-sync",
		Durable:     "amo-sync",
		Description: "amo-sync worker",
		AckWait:     5 * time.Minute,
		MaxDeliver:  1,
	})
	if err != nil {
		return nil, err
	}

	return consumer, nil
}
