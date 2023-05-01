package v1

import (
	"net/http"

	"github.com/gorilla/mux"
	"github.com/hanoy/messenger/internal/websocket"
)

func (h *Handler) InitChatRoutes(router *mux.Router) {
    router.HandleFunc("/chat", h.JoinChat).Methods(http.MethodGet)
}

func (h *Handler) JoinChat(w http.ResponseWriter, r *http.Request) {
    websocket.UpgradeConnection(w, r)
}
