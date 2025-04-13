package auth

import (
	"context"
	"time"

	"gitlab.ubrato.ru/ubrato/core/internal/broker"
	"gitlab.ubrato.ru/ubrato/core/internal/gateway/dadata"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/token"
	"gitlab.ubrato.ru/ubrato/core/internal/models"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"golang.org/x/exp/rand"
)

type Service struct {
	psql              DBTX
	userStore         UserStore
	organizationStore OrganizationStore
	sessionStore      SessionStore
	dadataGateway     DadataGateway
	tokenAuthorizer   TokenAuthorizer
	broker            Broker
}

const (
	sessionLength = 32
)

var sessionRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ!_-")

type DBTX interface {
	DB() store.QueryExecutor
	TX(ctx context.Context) (store.QueryExecutorTx, error)
	WithTransaction(ctx context.Context, fn store.ExecFn) (err error)
}

type UserStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.UserCreateParams) (models.User, error)
	GetWithOrganiztion(ctx context.Context, qe store.QueryExecutor, params store.UserGetParams) ([]models.RegularUser, error)
	GetWithEmployee(ctx context.Context, qe store.QueryExecutor, params store.UserGetParams) ([]models.EmployeeUser, error)
	Get(ctx context.Context, qe store.QueryExecutor, params store.UserGetParams) ([]models.FullUser, error)
}

type OrganizationStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, organization store.OrganizationCreateParams) (models.Organization, error)
	AddUser(ctx context.Context, qe store.QueryExecutor, params store.OrganizationAddUserParams) error
}

type SessionStore interface {
	Create(ctx context.Context, qe store.QueryExecutor, params store.SessionCreateParams) (models.Session, error)
	Get(ctx context.Context, qe store.QueryExecutor, params store.SessionGetParams) (models.Session, error)
	Update(ctx context.Context, qe store.QueryExecutor, params store.SessionUpdateParams) (models.Session, error)
	Delete(ctx context.Context, qe store.QueryExecutor, sessionID string) error
}

type DadataGateway interface {
	FindByINN(ctx context.Context, INN string) (dadata.FindByInnResponse, error)
}

type TokenAuthorizer interface {
	GenerateToken(payload token.Payload) (string, error)
	GetRefreshTokenDurationLifetime() time.Duration
	ValidateToken(rawToken string) (token.Claims, error)
}

type Broker interface {
	Publish(ctx context.Context, subject broker.Topic, data []byte) error
}

func New(
	psql DBTX,
	userStore UserStore,
	organizationStore OrganizationStore,
	sessionStore SessionStore,
	dadataGateway DadataGateway,
	tokenAuthorizer TokenAuthorizer,
	broker Broker,

) *Service {
	return &Service{
		psql:              psql,
		userStore:         userStore,
		organizationStore: organizationStore,
		sessionStore:      sessionStore,
		dadataGateway:     dadataGateway,
		tokenAuthorizer:   tokenAuthorizer,
		broker:            broker,
	}
}

func randSessionID(n int) string {
	rand.Seed(uint64(time.Now().Unix()))
	b := make([]rune, n)
	for i := range b {
		b[i] = sessionRunes[rand.Intn(len(sessionRunes))]
	}
	return string(b)
}
