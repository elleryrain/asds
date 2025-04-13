package mail

import (
	"context"

	"git.ubrato.ru/ubrato/dispatch-service/internal/broker"
	"git.ubrato.ru/ubrato/dispatch-service/internal/smtp"
)

type Mail interface {
	Start() error
	Stop()
}

type Service struct {
	broker     broker.Broker
	smtpClient smtp.SMTPClient
	ctx        context.Context

	emailQueueSub broker.Subscription
}

func NewService(ctx context.Context, broker broker.Broker, smtpClient smtp.SMTPClient) *Service {
	return &Service{
		ctx:        ctx,
		broker:     broker,
		smtpClient: smtpClient,
	}
}

func (s *Service) Start() (err error) {
	s.emailQueueSub, err = s.broker.Subscribe(s.ctx, "email_queue", s.emailQueueHandler)

	if err != nil {
		return err
	}

	return nil
}

func (s *Service) Stop() {
	s.emailQueueSub.Stop()
}
