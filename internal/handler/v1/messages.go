package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hanoy/messenger/internal/domain"
)

func (h *Handler) InitMessageRoutes(router *mux.Router) {
	router.HandleFunc("/msg", h.AddMessage).Methods(http.MethodPut)
	router.HandleFunc("/msg", h.FindAllMessages).Methods(http.MethodGet)
	router.HandleFunc("/msg/{id}", h.FindMessagesByID).Methods(http.MethodGet)
	router.HandleFunc("/msg/{id}", h.DeleteMessage).Methods(http.MethodDelete)
}

// url: api/msg
// method: put
func (h *Handler) AddMessage(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "applicatoin/json")

	var msg domain.Message
	err := json.NewDecoder(r.Body).Decode(&msg)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	msg, err = h.services.Messages.Add(r.Context(), msg)
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
