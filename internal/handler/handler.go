package handler

import (
	"github.com/gorilla/mux"
	v1 "github.com/hanoy/messenger/internal/handler/v1"
	"github.com/hanoy/messenger/internal/service"
)

type Handler struct {
    userService *service.UsersService
}

func NewHandler(service *service.UsersService) *Handler {
    return &Handler{service} 
}

func (h *Handler) Init() *mux.Router {
    router := mux.NewRouter() 
    handler := v1.NewHandler(h.userService)
    handler.InitRoutes(router)
    return router
}
