package mail

import (
	"errors"

	"git.ubrato.ru/ubrato/dispatch-service/internal/broker"
)

func (s *Service) emailQueueHandler(msg *broker.Message) (err error) {
	switch msg.Topic {
	case broker.EmailResetPassTopic:
		err = s.resetPassHandler(msg)
	case broker.ConfirmEmailTopic:
		err = s.emailConfirmHandler(msg)
	default:
		err = errors.New("topic not found")
	}
	if err != nil {
		return err
	}
	return nil
}
