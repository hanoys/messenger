package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/domain/dto"
)

func (h *Handler) InitMessageRoutes(router *mux.Router) {
    msgRouter := router.PathPrefix("/").Subrouter()
	msgRouter.HandleFunc("/msg", h.AddMessage).Methods(http.MethodPut)
	msgRouter.HandleFunc("/msg", h.FindAllMessages).Methods(http.MethodGet)
	msgRouter.HandleFunc("/msg/{id}", h.FindMessagesByID).Methods(http.MethodGet)
	msgRouter.HandleFunc("/msg/{id}", h.DeleteMessage).Methods(http.MethodDelete)
}

// url: api/msg
// method: put
func (h *Handler) AddMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "applicatoin/json")

	var msgDTO dto.AddMessageDTO
	err := json.NewDecoder(r.Body).Decode(&msgDTO)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

    var msg domain.Message
	msg, err = h.services.Messages.Add(r.Context(), msgDTO)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(msg)
}

// url: api/msg
// method: get
func (h *Handler) FindAllMessages(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "applicatoin/json")

	messages, err := h.services.Messages.FindAll(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(messages)
}

// url: api/msg/{id}
// method: get
func (h *Handler) FindMessagesByID(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "applicatoin/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	messages, err := h.services.Messages.FindByID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(messages)
}

// url: api/msg/{id}
// method: delete
func (h *Handler) DeleteMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "applicatoin/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	msg, err := h.services.Messages.Delete(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(msg)
}
