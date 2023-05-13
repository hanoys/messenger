package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/websocket"
)

func (h *Handler) InitChatRoutes(router *mux.Router) {
	router.HandleFunc("/chat", h.JoinChat).Methods(http.MethodGet)

    router.HandleFunc("/chats", h.FindAllChats).Methods(http.MethodGet)
	router.HandleFunc("/chats/{id}", h.FindChatsByID).Methods(http.MethodGet)
	router.HandleFunc("/chats", h.CreateChat).Methods(http.MethodPut)
    router.HandleFunc("/chats/{id}", h.DeleteChat).Methods(http.MethodDelete)

}

func (h *Handler) JoinChat(w http.ResponseWriter, r *http.Request) {
	websocket.UpgradeConnection(w, r)
}

// url: /api/chats
// method: get
func (h *Handler) FindAllChats(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "applicatoin/json")

	chats, err := h.services.Chats.FindAll()
	if err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}

	json.NewEncoder(w).Encode(chats)
}

// url: api/chats/{id}
// method: get
func (h *Handler) FindChatsByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
    if err != nil {
        writeError(w, http.StatusBadRequest, err.Error())
        return
    }

    chat, err := h.services.Chats.FindByID(id)
    if err != nil {
        writeError(w, http.StatusConflict, err.Error())
        return 
    }

    json.NewEncoder(w).Encode(chat)
}

// url: api/chats
// method: put
func (h *Handler) CreateChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

    var chat domain.Chat
    err := json.NewDecoder(r.Body).Decode(&chat)
    if err != nil {
        writeError(w, http.StatusBadRequest, err.Error())
        return
    }

    chat, err = h.services.Chats.Create(chat)
    if err != nil {
        writeError(w, http.StatusConflict, err.Error())
        return
    }

    json.NewEncoder(w).Encode(chat)
}

// url: api/chats/{id}
// method: delete
func (h *Handler) DeleteChat(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
    
    id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	} 

    chat, err := h.services.Chats.Delete(id)
    if err != nil {
        writeError(w, http.StatusConflict, err.Error())
        return
    }

    json.NewEncoder(w).Encode(chat)
}
