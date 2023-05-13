package handler

import (
	"github.com/gorilla/mux"
	v1 "github.com/hanoy/messenger/internal/handler/v1"
	"github.com/hanoy/messenger/internal/service"
)

type Handler struct {
    services *service.Services
}

func NewHandler(services *service.Services) *Handler {
    return &Handler{services} 
}

func (h *Handler) Init() *mux.Router {
    router := mux.NewRouter() 
    apiRouter := router.PathPrefix("/api").Subrouter()
    handler := v1.NewHandler(h.services)
    handler.InitRoutes(apiRouter)
    return router
}
