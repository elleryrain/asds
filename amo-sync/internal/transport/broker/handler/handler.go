package handler

import (
	"context"
	"log/slog"

	"gitlab.ubrato.ru/ubrato/amo-sync/internal/service/amoproxy"
)

type Handler struct {
	logger      *slog.Logger
	amoProxySvc AmoProxyService
}

type AmoProxyService interface {
	CreateCompany(ctx context.Context, params amoproxy.CreateCompanyParams) (int, error)
	CreateContact(ctx context.Context, params amoproxy.CreateContactParams) (int, error)
	CreateLead(ctx context.Context, params amoproxy.CreateLeadParams) (int, error)
	CreateNotes(ctx context.Context, params amoproxy.CreateNotesParams) error
	CreateLink(ctx context.Context, params amoproxy.CreateLink) error
}

func New(logger *slog.Logger, amoProxySvc AmoProxyService) *Handler {
	return &Handler{
		logger:      logger,
		amoProxySvc: amoProxySvc,
	}
}
