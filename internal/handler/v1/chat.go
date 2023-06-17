package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/domain/dto"
)

func (h *Handler) InitChatRoutes(router *mux.Router) {
	chatRouter := router.PathPrefix("/").Subrouter()
	chatRouter.Use(h.userJWTAuth)
	chatRouter.HandleFunc("/chats/{id}", h.FindChatsByID).Methods(http.MethodGet)
	chatRouter.HandleFunc("/chats", h.CreateChat).Methods(http.MethodPut)
	chatRouter.HandleFunc("/chats/{id}", h.DeleteChat).Methods(http.MethodDelete)

}

// url: api/chats/{id}
// method: get
func (h *Handler) FindChatsByID(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	chat, err := h.services.Chats.FindByID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}
	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chat)
}

// url: api/chats
// method: put
func (h *Handler) CreateChat(w http.ResponseWriter, r *http.Request) {
	var chatDTO dto.CreateChatDTO
	err := json.NewDecoder(r.Body).Decode(&chatDTO)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	var chat domain.Chat
	chat, err = h.services.Chats.Create(r.Context(), chatDTO)
	if err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chat)
}

// url: api/chats/{id}
// method: delete
func (h *Handler) DeleteChat(w http.ResponseWriter, r *http.Request) {
	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	chat, err := h.services.Chats.Delete(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}

	w.Header().Add("Content-Type", "application/json")
	json.NewEncoder(w).Encode(chat)
}
