package main

import (
	"fmt"
	"log/slog"
	"os"
	"os/signal"
	"syscall"

	"gitlab.ubrato.ru/ubrato/core/internal/broker/jetstream"
	"gitlab.ubrato.ru/ubrato/core/internal/config"
	dadataGateway "gitlab.ubrato.ru/ubrato/core/internal/gateway/dadata"
	"gitlab.ubrato.ru/ubrato/core/internal/lib/token"
	authService "gitlab.ubrato.ru/ubrato/core/internal/service/auth"
	catalogService "gitlab.ubrato.ru/ubrato/core/internal/service/catalog"
	favouriteService "gitlab.ubrato.ru/ubrato/core/internal/service/favourite"
	organizationService "gitlab.ubrato.ru/ubrato/core/internal/service/organization"
	portfolioService "gitlab.ubrato.ru/ubrato/core/internal/service/portfolio"
	questionAnswerService "gitlab.ubrato.ru/ubrato/core/internal/service/question_answer"
	questionnaireService "gitlab.ubrato.ru/ubrato/core/internal/service/questionnaire"
	respondService "gitlab.ubrato.ru/ubrato/core/internal/service/respond"
	suggestService "gitlab.ubrato.ru/ubrato/core/internal/service/suggest"
	surveyService "gitlab.ubrato.ru/ubrato/core/internal/service/survey"
	tenderService "gitlab.ubrato.ru/ubrato/core/internal/service/tender"
	userService "gitlab.ubrato.ru/ubrato/core/internal/service/user"
	verificationService "gitlab.ubrato.ru/ubrato/core/internal/service/verification"
	winnersService "gitlab.ubrato.ru/ubrato/core/internal/service/winners"
	"gitlab.ubrato.ru/ubrato/core/internal/store"
	"gitlab.ubrato.ru/ubrato/core/internal/store/postgres"
	additionStore "gitlab.ubrato.ru/ubrato/core/internal/store/postgres/addition"
	catalogStore "gitlab.ubrato.ru/ubrato/core/internal/store/postgres/catalog"
	favouriteStore "gitlab.ubrato.ru/ubrato/core/internal/store/postgres/favourite"
	organizationStore "gitlab.ubrato.ru/ubrato/core/internal/store/postgres/organization"
	portfolioStore "gitlab.ubrato.ru/ubrato/core/internal/store/postgres/portfolio"
	questionAnswerStore "gitlab.ubrato.ru/ubrato/core/internal/store/postgres/question_answer"
	questionnaireStore "gitlab.ubrato.ru/ubrato/core/internal/store/postgres/questionnaire"
	respondStore "gitlab.ubrato.ru/ubrato/core/internal/store/postgres/respond"
	sessionStore "gitlab.ubrato.ru/ubrato/core/internal/store/postgres/session"
	tenderStore "gitlab.ubrato.ru/ubrato/core/internal/store/postgres/tender"
	userStore "gitlab.ubrato.ru/ubrato/core/internal/store/postgres/user"
	verificationStore "gitlab.ubrato.ru/ubrato/core/internal/store/postgres/verification"
	winnersStore "gitlab.ubrato.ru/ubrato/core/internal/store/postgres/winners"
	"gitlab.ubrato.ru/ubrato/core/internal/transport/http"
	authHandler "gitlab.ubrato.ru/ubrato/core/internal/transport/http/handlers/auth"
	catalogHandler "gitlab.ubrato.ru/ubrato/core/internal/transport/http/handlers/catalog"
	employeeHandler "gitlab.ubrato.ru/ubrato/core/internal/transport/http/handlers/employee"
	errorHandler "gitlab.ubrato.ru/ubrato/core/internal/transport/http/handlers/error"
	organizationHandler "gitlab.ubrato.ru/ubrato/core/internal/transport/http/handlers/organization"
	questionnaireHandler "gitlab.ubrato.ru/ubrato/core/internal/transport/http/handlers/questionnaire"
	suggestHandler "gitlab.ubrato.ru/ubrato/core/internal/transport/http/handlers/suggest"
	surveyHandler "gitlab.ubrato.ru/ubrato/core/internal/transport/http/handlers/survey"
	tenderHandler "gitlab.ubrato.ru/ubrato/core/internal/transport/http/handlers/tender"
	userHandler "gitlab.ubrato.ru/ubrato/core/internal/transport/http/handlers/user"
	verificationHandler "gitlab.ubrato.ru/ubrato/core/internal/transport/http/handlers/verification"
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

	psqlDB, err := postgres.New(cfg.Store.Postgres)
	if err != nil {
		return fmt.Errorf("create postgres: %w", err)
	}
	psql := store.New(psqlDB)

	userStore := userStore.NewUserStore()
	sessionStore := sessionStore.NewSessionStore()
	catalogStore := catalogStore.NewCatalogStore()
	organizationStore := organizationStore.NewOrganizationStore(catalogStore)
	tenderStore := tenderStore.NewTenderStore(catalogStore)
	verificationStore := verificationStore.NewVerificationRequestStore()
	additionStore := additionStore.NewAdditionStore()
	questionnaireStore := questionnaireStore.NewQuestionnaireStore()
	questionAnswerStore := questionAnswerStore.NewQuestionAnswerStore()
	portfolioStore := portfolioStore.NewPortfolioStore()
	respondStore := respondStore.NewPesponsStore()
	favouriteStore := favouriteStore.NewFavouriteStore()
	winnersStore := winnersStore.NewWinnersStore()

	dadataGateway := dadataGateway.NewClient(cfg.Gateway.Dadata.APIKey)

	tokenAuthorizer, err := token.NewTokenAuthorizer(cfg.Auth.JWT)
	if err != nil {
		return fmt.Errorf("init token authorizer: %w", err)
	}

	jetStream, err := jetstream.New(logger, cfg.Broker.Nats.Address)
	if err != nil {
		return fmt.Errorf("init jetstream: %w", err)
	}

	authService := authService.New(
		psql,
		userStore,
		organizationStore,
		sessionStore,
		dadataGateway,
		tokenAuthorizer,
		jetStream,
	)

	respondService := respondService.New(
		psql,
		respondStore,
	)

	tenderService := tenderService.New(
		psql,
		tenderStore,
		additionStore,
		verificationStore,
		jetStream,
	)

	catalogService := catalogService.New(
		psql,
		catalogStore,
	)

	userService := userService.New(
		psql,
		userStore,
		jetStream,
	)

	surveyService := surveyService.New(
		psql,
		jetStream,
	)

	organizationService := organizationService.New(
		psql,
		organizationStore,
	)

	suggestService := suggestService.New(
		psql,
		dadataGateway,
		catalogStore,
	)

	verificationService := verificationService.New(
		psql,
		verificationStore,
		tenderStore,
		additionStore,
		organizationStore,
		questionAnswerStore,
		jetStream,
		userStore,
	)

	questionnaireService := questionnaireService.New(
		psql,
		questionnaireStore,
		organizationStore,
	)

	questionAnswerService := questionAnswerService.New(
		psql,
		questionAnswerStore,
		tenderStore,
		verificationStore,
		jetStream,
	)

	portfolioService := portfolioService.New(
		psql,
		portfolioStore,
	)

	favouriteService := favouriteService.New(
		psql,
		favouriteStore,
		tenderStore,
		organizationStore,
	)

	winnersService := winnersService.New(
		psql,
		jetStream,
		winnersStore,
		tenderStore,
		respondStore,
		userStore,
	)

	router := http.NewRouter(http.RouterParams{
		Error:         errorHandler.New(logger),
		Auth:          authHandler.New(logger, authService, userService),
		Tenders:       tenderHandler.New(logger, tenderService, questionAnswerService, respondService, winnersService),
		Catalog:       catalogHandler.New(logger, catalogService),
		Users:         userHandler.New(logger, userService),
		Survey:        surveyHandler.New(logger, surveyService),
		Organization:  organizationHandler.New(logger, organizationService, portfolioService, favouriteService),
		Suggest:       suggestHandler.New(logger, suggestService),
		Verification:  verificationHandler.New(logger, verificationService),
		Employee:      employeeHandler.New(logger, userService),
		Questionnaire: questionnaireHandler.New(logger, questionnaireService),
	})

	server, err := http.NewServer(logger, cfg.Transport.HTTP, router)
	if err != nil {
		return fmt.Errorf("create http server: %w", err)
	}

	go func() {
		<-sig
		logger.Info("Received termination signal, cleaning up")
		err := server.Stop()
		if err != nil {
			logger.Error("Stop http server", "error", err)
		}
	}()

	err = server.Start()
	if err != nil {
		return fmt.Errorf("serve http: %w", err)
	}

	return nil
}
