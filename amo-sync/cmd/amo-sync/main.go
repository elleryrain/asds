package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"gitlab.ubrato.ru/ubrato/amo-sync/internal/broker/jetstream"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/config"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/gateway/amocrm"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/service/amoproxy"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/store"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/store/postgres"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/transport/broker"
	"gitlab.ubrato.ru/ubrato/amo-sync/internal/transport/broker/handler"
)

func main() {
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))

	cfg, err := config.ReadConfig()
	if err != nil {
		logger.Error("Error parsing default config from env", "error", err)
		os.Exit(1)
	}

	if cfg.Debug {
		logger = slog.New(slog.NewTextHandler(os.Stdout, &slog.HandlerOptions{
			Level: slog.LevelDebug,
		}))

		logger.Warn("Debug enabled")
	}

	if err := run(cfg, logger); err != nil {
		logger.Error("Error initializing service", "error", err)
		os.Exit(1)
	}

}

func run(cfg config.Default, logger *slog.Logger) error {
	sig := make(chan os.Signal, 1)
	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM)

	postgresConn, err := postgres.New(cfg.Store.Postgres)
	if err != nil {
		return fmt.Errorf("postgres connection: %w", err)
	}

	jetstream, err := jetstream.New(logger, cfg.Broker.JetStream.Addr)
	if err != nil {
		return err
	}

	psql := store.New(postgresConn)

	amoStore := postgres.NewAmoStore()

	amoCRMGateway := amocrm.NewClient(cfg.Gateway.AmoCRM.Token)

	amoProxySvc := amoproxy.New(amoCRMGateway, amoStore, psql)

	handler := handler.New(logger, amoProxySvc)

	broker := broker.New(logger, jetstream, handler)

	err = broker.Start()
	if err != nil {
		return err
	}

	<-sig
	logger.Info("Received termination signal, cleaning up")
	broker.Stop()
	postgresConn.Close()

	return nil
}
