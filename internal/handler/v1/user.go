package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hanoy/messenger/internal/domain"
)

func (h *Handler) InitUserRoutes(router *mux.Router) {
	router.HandleFunc("/users", h.FindAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{id}", h.FindUserById).Methods(http.MethodGet)
	router.HandleFunc("/users", h.CreateUser).Methods(http.MethodPut)
}

// url: users/
// method: get
func (h *Handler) FindAllUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	users, err := h.services.Users.FindAll()
	if err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}

	json.NewEncoder(w).Encode(users)
}

// url: users/{id}
// method: get
func (h *Handler) FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.Users.FindByID(id)
	if err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}

	json.NewEncoder(w).Encode(user)
}

// url: users/
// method: post
func (h *Handler) CreateUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var user domain.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err = h.services.Users.Create(user)
	if err != nil {
		writeError(w, http.StatusConflict, err.Error())
		return
	}

	json.NewEncoder(w).Encode(user)
}

// url: users/{id}
// method: delete
func (h *Handler) DeleteUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Add("Content-Type", "application/json")
    
    id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	} 

    user, err := h.services.Chats.Delete(id)
    if err != nil {
        writeError(w, http.StatusConflict, err.Error())
        return
    }

    json.NewEncoder(w).Encode(user)

}
