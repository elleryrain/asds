package jetstream

import (
	"context"
	"fmt"
	"log/slog"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"gitlab.ubrato.ru/ubrato/core/internal/broker"
)

type Broker struct {
	logger *slog.Logger
	js     jetstream.JetStream
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

	return &Broker{
		logger: logger,
		js:     js,
	}, nil
}

func (b *Broker) Publish(ctx context.Context, subject broker.Topic, data []byte) error {
	_, err := b.js.Publish(ctx, string(subject), data)
	return err
}
