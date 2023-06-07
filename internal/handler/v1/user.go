package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/domain/dto"
	"github.com/hanoy/messenger/pkg/auth"
)

func (h *Handler) InitUserRoutes(router *mux.Router) {
	userRouter := router.PathPrefix("/").Subrouter()

	userRouter.HandleFunc("/user/sign-up", h.SignUpUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/user/log-in", h.LogInUser).Methods(http.MethodPost)

	authenticatedUserRouter := userRouter.PathPrefix("/").Subrouter()
	authenticatedUserRouter.Use(h.jwtAuth)
	authenticatedUserRouter.HandleFunc("/user/log-out", h.LogOutUser).Methods(http.MethodPost)
	authenticatedUserRouter.HandleFunc("/user/jwt-refresh", h.RefreshToken).Methods(http.MethodPost)

	authenticatedUserRouter.HandleFunc("/user", h.FindAllUsers).Methods(http.MethodGet)
	authenticatedUserRouter.HandleFunc("/user/{id}", h.FindUserById).Methods(http.MethodGet)
	authenticatedUserRouter.HandleFunc("/user", h.CreateUser).Methods(http.MethodPut)
	authenticatedUserRouter.HandleFunc("/user/{id}", h.DeleteUser).Methods(http.MethodDelete)
}

// url:    /api/user/sign-up
// method: post
func (h *Handler) SignUpUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var userDTO dto.SignUpUserDTO
	err := json.NewDecoder(r.Body).Decode(&userDTO)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	_, err = h.services.Users.FindByEmail(r.Context(), dto.FindByEmailUserDTO{Email: userDTO.Email})
	if err == nil {
		writeError(w, http.StatusBadRequest, "user with such email already exists")
		return
	}

	user, err := h.services.Users.Create(r.Context(), dto.CreateUserDTO{
		FirstName: userDTO.FirstName,
		LastName:  userDTO.LastName,
		Email:     userDTO.Email,
		Login:     userDTO.Login,
		Password:  userDTO.Password,
	})
    if err != nil {
        writeError(w, http.StatusInternalServerError, err.Error())
        return
    }

	tokenPayload, err := auth.NewPayload(user.ID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tokenString, err := h.tokenProvider.CreateToken(tokenPayload)
    if err != nil {
        writeError(w, http.StatusInternalServerError, err.Error())
        return
    }
	token := struct {
		Token string `json:"token"`
	}{
		Token: tokenString}

	json.NewEncoder(w).Encode(token)
}

// url:    /api/user/log-in
// method: post
func (h *Handler) LogInUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")
    
    var userDTO dto.LogInUserDTO
    err := json.NewDecoder(r.Body).Decode(&userDTO)
    if err != nil {
        writeError(w, http.StatusBadRequest, err.Error())
        return
    }

    user, err := h.services.Users.FindByCredentials(r.Context(), userDTO)
    if err != nil {
        writeError(w, http.StatusBadRequest, "email or password is invalid")
        return
    }

    tokenPayload, err := auth.NewPayload(user.ID)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	tokenString, err := h.tokenProvider.CreateToken(tokenPayload)
    if err != nil {
        writeError(w, http.StatusInternalServerError, err.Error())
        return
    }
	token := struct {
		Token string `json:"token"`
	}{
		Token: tokenString}

	json.NewEncoder(w).Encode(token)

}

func (h *Handler) LogOutUser(w http.ResponseWriter, r *http.Request) {
    
}

func (h *Handler) RefreshToken(w http.ResponseWriter, r *http.Request) {

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

// url: api/users/{id}
// method: get
func (h *Handler) FindUserById(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	id, err := strconv.Atoi(mux.Vars(r)["id"])
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	user, err := h.services.Users.FindByID(r.Context(), id)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

	json.NewEncoder(w).Encode(user)
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
