package handler

import (
	"github.com/gorilla/mux"
	v1 "github.com/hanoy/messenger/internal/handler/v1"
	"github.com/hanoy/messenger/internal/service"
	"github.com/hanoy/messenger/pkg/auth"
)

type Handler struct {
	services      *service.Services
	tokenProvider *auth.Provider
}

func NewHandler(services *service.Services, provider *auth.Provider) *Handler {
	return &Handler{services, provider}
}

func (h *Handler) Init() *mux.Router {
	router := mux.NewRouter()
	apiRouter := router.PathPrefix("/api").Subrouter()
	handler := v1.NewHandler(h.services, h.tokenProvider)
	handler.InitRoutes(apiRouter)
	return router
}
