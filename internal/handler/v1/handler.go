package v1

import (
	"github.com/gorilla/mux"
	"github.com/hanoy/messenger/internal/service"
)

type Handler struct {
	userService *service.UsersService
}

func NewHandler(userService *service.UsersService) *Handler {
	return &Handler{userService}
}

func (h *Handler) InitRoutes(router *mux.Router) {
    h.InitUserRoutes(router)
    h.InitChatRoutes(router)
}
