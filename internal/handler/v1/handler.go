package v1

import (
	"github.com/gorilla/mux"
	"github.com/hanoy/messenger/internal/service"
	"github.com/hanoy/messenger/pkg/auth"
)

type Handler struct {
    services *service.Services
    tokenProvider *auth.Provider
}

func NewHandler(services *service.Services, provider *auth.Provider) *Handler {
	return &Handler{services, provider}
}

func (h *Handler) InitRoutes(router *mux.Router) {
    h.InitUserRoutes(router)
    h.InitAdminRoutes(router)
    h.InitChatRoutes(router)
    h.InitMessageRoutes(router)
}
