package jetstream

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/nats-io/nats.go/jetstream"
	"github.com/rs/zerolog/log"
	"gitlab.ubrato.ru/ubrato/notification/internal/models"
	eventsv1 "gitlab.ubrato.ru/ubrato/notification/internal/models/gen/proto/events/v1"
	"gitlab.ubrato.ru/ubrato/notification/internal/service"
	"gitlab.ubrato.ru/ubrato/notification/internal/transport/broker"
	"google.golang.org/protobuf/proto"
)

type Broker struct {
	js                  jetstream.JetStream
	consumer            jetstream.ConsumeContext
	notificationStream  jetstream.Stream
	notificationService NotificationService
	// UserNotifications хранит уведомления для пользователей по их userID, используется для отправки в реальном времени
	UserNotifications map[int]chan models.Notification
	mu                sync.RWMutex
}

type NotificationService interface {
	Create(ctx context.Context, params service.NotifictionCreateParams) (models.Notification, error)
}

type BrokerParams struct {
	Addr                string
	NotificationService NotificationService
}

func New(params BrokerParams) (*Broker, error) {
	nc, err := nats.Connect(params.Addr)
	if err != nil {
		return nil, fmt.Errorf("nats connect: %w", err)
	}

	js, err := jetstream.New(nc)
	if err != nil {
		return nil, fmt.Errorf("create jetstream: %w", err)
	}

	notificationStream, err := js.CreateOrUpdateStream(context.Background(), jetstream.StreamConfig{
		Name: "notification-stream",
		Subjects: []string{
			broker.NotifyUserEmailConfirmation,
			broker.NotifyUserEmailConfirmed,
			broker.NotifyOrganizationVerification,
			broker.NotifyTenderVerification,
			broker.NotifyTenderAdditionVerification,
			broker.NotifyTenderQA,
			broker.NotifyTenderQuestionAnswerVerification,
			broker.NotifyTenderWinners,
		}})
	if err != nil {
		return nil, fmt.Errorf("create stream: %w", err)
	}
	log.Debug().Str("stream_name", "notification-stream").Msg("JetStream stream created")

	return &Broker{
		js:                  js,
		notificationStream:  notificationStream,
		notificationService: params.NotificationService,
		UserNotifications:   make(map[int]chan models.Notification),
		mu:                  sync.RWMutex{},
	}, nil
}

func (b *Broker) RunConsumer() error {
	consumer, err := b.notificationStream.CreateOrUpdateConsumer(context.Background(), jetstream.ConsumerConfig{
		Name:          "notification",
		AckWait:       5 * time.Minute,
		MaxDeliver:    1,
		DeliverPolicy: jetstream.DeliverNewPolicy,
	})
	if err != nil {
		return fmt.Errorf("create consumer: %w", err)
	}

	b.consumer, err = consumer.Consume(b.handleMsg)
	if err != nil {
		return fmt.Errorf("consume: %w", err)
	}
	log.Debug().Str("consumer_name", consumer.CachedInfo().Name).Msg("Consumer started successfully")

	return nil
}

func (b *Broker) Stop() {
	b.consumer.Stop()
	log.Debug().Msg("Consumer stopped")
}

func (b *Broker) handleMsg(msg jetstream.Msg) {
	var err error

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	defer func() {
		if err != nil {
			if nakErr := msg.Nak(); nakErr != nil {
				log.Err(err).Msg("Nak message")
			}
			return
		}

		if err = msg.Ack(); err != nil {
			log.Err(err).Msg("Ack message")
		}
	}()

	event := &eventsv1.SentNotification{}
	err = proto.Unmarshal(msg.Data(), event)
	if err != nil {
		log.Err(err).Msg("Unmarshal proto notification")
		return
	}
	eventNotify := event.Notification

	var notification models.Notification
	switch msg.Subject() {
	case broker.NotifyUserEmailConfirmation:
		notification = models.Notification{
			UserID:  int(eventNotify.User.GetId()),
			Title:   "Подтверждение email",
			Comment: "Для завершения регистрации подтвердите email",
			ActionButton: &models.ActionButton{
				Text:   "Подтвердить",
				Url:    "/profile/settings",
				Styled: true,
			},
		}

	case broker.NotifyUserEmailConfirmed:
		notification = models.Notification{
			UserID:  int(eventNotify.User.GetId()),
			Title:   "Верификация компании",
			Comment: "Вы успешно зарегистрировались. Необходимо пройти верификацию компании.",
			ActionButton: &models.ActionButton{
				Text:   "Перейти",
				Url:    "/profile/documents",
				Styled: true,
			},
		}

	case broker.NotifyOrganizationVerification:
		notification = models.MustOrganizationVerificationNotify(eventNotify)

	case broker.NotifyTenderVerification:
		notification = models.MustTenderVerificationNotify(eventNotify)

	case broker.NotifyTenderAdditionVerification:
		notification = models.MustAdditionVerificationNotify(eventNotify)

	case broker.NotifyTenderQA:
		notification = models.MustQuestionAnswerNotify(eventNotify)

	case broker.NotifyTenderQuestionAnswerVerification:
		notification = models.MustQuestionAnswerVerificationNotify(eventNotify)

	case broker.NotifyTenderWinners:
		notification = models.MustWinnerNotify(eventNotify)

	default:
		err = fmt.Errorf("invalid topic: %v", msg.Subject())
		return
	}

	err = b.SaveNotification(ctx, notification)
	if err != nil {
		log.Err(err).Msg("Save notification")
	}
}

func (b *Broker) CreateUserChan(userID int) chan models.Notification {
	b.mu.Lock()
	defer b.mu.Unlock()

	if _, ok := b.UserNotifications[userID]; !ok {
		log.Debug().Int("user_id", userID).Msg("Creating new user channel")
		b.UserNotifications[userID] = make(chan models.Notification, 10)
	}

	return b.UserNotifications[userID]
}

func (b *Broker) DeleteUserChan(userID int) {
	b.mu.Lock()
	defer b.mu.Unlock()

	if userChan, ok := b.UserNotifications[userID]; ok {
		close(userChan)
		delete(b.UserNotifications, userID)
		log.Debug().Int("user_id", userID).Msg("Closing and deleting user channel")
	}
}
