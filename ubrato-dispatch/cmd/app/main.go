package main

import (
	"context"
	"log/slog"
	"os"
	"sync"

	"git.ubrato.ru/ubrato/dispatch-service/internal/broker/jetstream"
	"git.ubrato.ru/ubrato/dispatch-service/internal/config"
	"git.ubrato.ru/ubrato/dispatch-service/internal/service"
	"git.ubrato.ru/ubrato/dispatch-service/internal/service/mail"
	"git.ubrato.ru/ubrato/dispatch-service/internal/smtp"
)

func main() {
	sig := make(chan os.Signal, 1)
	ctx := context.Background()
	logger := slog.New(slog.NewTextHandler(os.Stdout, nil))
	config := config.Load()

	jetstream, err := jetstream.New(config.Nats.Addr, logger)
	if err != nil {
		logger.Error("Error opening nats connection", "error", err)
		os.Exit(1)
	}

	smtpClient, err := smtp.NewSMTPClient(
		config.Smtp.Login,
		config.Smtp.Pass,
		config.Smtp.Host,
		config.Smtp.From,
	)

	if err != nil {
		logger.Error("Error init smtp client", "error", err)
		os.Exit(1)
	}

	mailService := mail.NewService(ctx, jetstream, smtpClient)

	services := service.Services{
		Mail: mailService,
	}

	wg := sync.WaitGroup{}

	err = services.Mail.Start()
	if err != nil {
		logger.Error("Error setup jetstream", "error", err)
		os.Exit(1)
	}

	wg.Add(1)
	go func() {
		<-sig
		logger.Info("Received kill signal")
		defer services.Mail.Stop()
		wg.Done()
	}()

	wg.Wait()
}
