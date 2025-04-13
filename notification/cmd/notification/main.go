package main

import (
	"flag"
	"os"
	"os/signal"
	"syscall"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	"gitlab.ubrato.ru/ubrato/notification/internal/config"
	"gitlab.ubrato.ru/ubrato/notification/internal/lib/token"
	authSrv "gitlab.ubrato.ru/ubrato/notification/internal/service/auth"
	notificationSrv "gitlab.ubrato.ru/ubrato/notification/internal/service/notification"
	"gitlab.ubrato.ru/ubrato/notification/internal/store"
	"gitlab.ubrato.ru/ubrato/notification/internal/store/postgres"
	notificationStore "gitlab.ubrato.ru/ubrato/notification/internal/store/postgres/notification"
	"gitlab.ubrato.ru/ubrato/notification/internal/transport/broker/jetstream"
	"gitlab.ubrato.ru/ubrato/notification/internal/transport/ogen"
	authHandler "gitlab.ubrato.ru/ubrato/notification/internal/transport/ogen/auth"
	errorHandler "gitlab.ubrato.ru/ubrato/notification/internal/transport/ogen/error"
	notificationHandler "gitlab.ubrato.ru/ubrato/notification/internal/transport/ogen/notification"
)

func init() {
	logLevel := flag.String("loglevel", "info", "log level: debug, info, warn, error")
	flag.Parse()

	setupLogger(*logLevel)
}

func main() {
	cfg, err := config.ReadConfig()
	if err != nil {
		log.Fatal().Err(err).Msg("Error parsing default config from env")
	}

	// DB
	pool, err := postgres.NewPG(cfg.Store.Postgres.DSN())
	if err != nil {
		log.Fatal().Err(err).Msg("Unable to connect to pool")
	}
	defer pool.Close()
	dbtx := store.NewDBTX(pool)

	notificationStore := notificationStore.NewNotificationStore()

	// SERVICE
	tokenAuthorizer, err := token.NewTokenAuthorizer(cfg.Auth.JWT)
	if err != nil {
		log.Fatal().Err(err).Msg("Init token authorizer")
	}

	notificationSrv := notificationSrv.New(
		dbtx,
		notificationStore)

	authSrv := authSrv.New(tokenAuthorizer)

	// JETSTREAM
	jetStream, err := jetstream.New(jetstream.BrokerParams{
		Addr:                cfg.Broker.Nats.Address,
		NotificationService: notificationSrv,
	})
	if err != nil {
		log.Fatal().Err(err).Msg("Init Jetstream")
	}

	if err := jetStream.RunConsumer(); err != nil {
		log.Fatal().Err(err).Msg("Run comsumer Jetstream")
	}
	defer jetStream.Stop()

	// SERVER
	router := ogen.NewRouter(ogen.RouterParams{
		Auth:         authHandler.New(authSrv),
		Error:        errorHandler.New(),
		Notification: notificationHandler.New(notificationSrv, authSrv, jetStream),
	})

	server, err := ogen.NewServer(cfg.Transport.HTTP, router)
	if err != nil {
		log.Fatal().Err(err)
	}

	go func() {
		log.Info().Msgf("The server is starting on %v", cfg.Transport.HTTP.Port)

		if err := server.Start(); err != nil {
			log.Fatal().Err(err).Msg("Error occured while running http server")
		}
	}()

	done := make(chan os.Signal, 1)
	signal.Notify(done, syscall.SIGTERM, syscall.SIGINT)
	<-done

	if err := server.Stop(); err != nil {
		log.Err(err).Msg("Error occured on server shutting down")
	}
}

func setupLogger(level string) {
	zlevel, err := zerolog.ParseLevel(level)
	if err != nil {
		zlevel = zerolog.InfoLevel
	}

	zerolog.SetGlobalLevel(zlevel)
	zerolog.TimeFieldFormat = "02.01.2006 15:04:05"

	switch zlevel {
	case zerolog.DebugLevel:
		log.Logger = zerolog.New(zerolog.ConsoleWriter{
			Out:        os.Stderr,
			TimeFormat: "15:04:05"}).
			With().
			Timestamp().
			Caller().
			Logger()
	default:
		log.Logger = zerolog.New(os.Stdout).
			With().
			Timestamp().
			Caller().
			Logger()
	}
}
