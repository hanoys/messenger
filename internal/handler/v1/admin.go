package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/domain/dto"
)

func (h *Handler) InitAdminRoutes(router *mux.Router) {
	adminRouter := router.PathPrefix("/").Subrouter()

	adminRouter.HandleFunc("/admin/log-in", h.LogInAdmin).Methods(http.MethodPost)

	authAdminRouter := adminRouter.PathPrefix("/").Subrouter()
	authAdminRouter.Use(h.adminJWTAuth)
	authAdminRouter.HandleFunc("/admin/log-out", h.LogOutAdmin)
	authAdminRouter.HandleFunc("/admin/refresh", h.RefreshTokenAdmin)
	{
		authAdminRouter.HandleFunc("/user", h.FindAllUsers).Methods(http.MethodGet)
		authAdminRouter.HandleFunc("/user", h.CreateUser).Methods(http.MethodPut)
		authAdminRouter.HandleFunc("/user/{id}", h.DeleteUser).Methods(http.MethodDelete)
	}
	{
		authAdminRouter.HandleFunc("/chats", h.FindAllChats).Methods(http.MethodGet)
	}
	{
		authAdminRouter.HandleFunc("/msg", h.FindAllMessages).Methods(http.MethodGet)
	}
}

// url: api/users
// method: get
func (h *Handler) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	users, err := h.services.Users.FindAll(r.Context())
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(users)
}

// url: /api/chats
// method: get
func (h *Handler) FindAllChats(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "applicatoin/json")

	chats, err := h.services.Chats.FindAll(r.Context())
	if err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}

	json.NewEncoder(w).Encode(chats)
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

// url: api/users/
// method: put
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var userDTO dto.CreateUserDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	var user domain.User
	user, err = h.services.Users.Create(r.Context(), userDTO)
	if err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}

	json.NewEncoder(w).Encode(user)
}

// url: api/users/{id}
// method: delete
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.Users.Delete(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}

	json.NewEncoder(w).Encode(user)

}
