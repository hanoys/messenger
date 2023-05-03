package v1

import (
	"github.com/gorilla/mux"
	"github.com/hanoy/messenger/internal/service"
)

type Handler struct {
    services *service.Services
}

func NewHandler(services *service.Services) *Handler {
	return &Handler{services}
}

func (h *Handler) InitRoutes(router *mux.Router) {
    h.InitUserRoutes(router)
    h.InitChatRoutes(router)
}
