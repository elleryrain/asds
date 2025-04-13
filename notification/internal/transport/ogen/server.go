package ogen

import (
	"context"
	"errors"
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/go-chi/jwtauth/v5"
	"github.com/rs/zerolog/log"
	api "gitlab.ubrato.ru/ubrato/notification/api/gen"
	"gitlab.ubrato.ru/ubrato/notification/internal/config"
	"gitlab.ubrato.ru/ubrato/notification/internal/lib/token"
)

type Server struct {
	srv *http.Server
}

func NewServer(cfg config.HTTP, router *Router) (*Server, error) {
	apiServer, err := api.NewServer(
		router,
		router.Auth,
		api.WithErrorHandler(router.HandleError),
	)
	if err != nil {
		return nil, fmt.Errorf("init http server: %w", err)
	}

	r := chi.NewRouter()
	registerMiddleware(r)
	swaggerUI(r, cfg.SwaggerPath)

	r.Mount("/", apiServer)

	// Заменяет заглушку из ogen, тк не поддерживается sse
	r.Group(func(r chi.Router) {
		r.Use(jwtauth.Verifier(token.TokenAuth)) // Проверка токена
		r.Use(jwtauth.Authenticator(token.TokenAuth))
		r.Get("/v1/notifications/{userID}/stream", router.V1GetUserNotificationsBySSE())
	})

	server := &Server{
		srv: &http.Server{
			Addr:    fmt.Sprintf(":%d", cfg.Port),
			Handler: r,
		},
	}

	return server, nil
}

func newCORShandler() func(http.Handler) http.Handler {
	return cors.Handler(cors.Options{
		AllowedOrigins: []string{
			"http://ubrato.ru",
			"https://ubrato.ru",
			"http://dev.ubrato.ru",
			"https://dev.ubrato.ru",
			"http://admin.ubrato.ru",
			"https://admin.ubrato.ru",
			"http://localhost",
			"http://localhost:5174",
			"http://localhost:5173",
			"http://localhost:3000",
		},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		AllowCredentials: true,
	})
}

func (s *Server) Start() error {
	log.Info().Str("addr", s.srv.Addr).Msg("Starting http server")
	err := s.srv.ListenAndServe()
	if err != nil && !errors.Is(err, http.ErrServerClosed) {
		return err
	}

	return nil
}

func (s *Server) Stop() error {
	log.Info().Msg("Shutting down http server")
	return s.srv.Shutdown(context.Background())
}
