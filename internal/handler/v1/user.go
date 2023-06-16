package v1

import (
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/gorilla/mux"
	"github.com/hanoy/messenger/internal/domain"
)

func (h *Handler) InitUserRoutes(router *mux.Router) {
	userRouter := router.PathPrefix("/").Subrouter()

	userRouter.HandleFunc("/user/sign-up", h.SignUpUser).Methods(http.MethodPost)
	userRouter.HandleFunc("/user/log-in", h.LogInUser).Methods(http.MethodPost)

	authUserRouter := userRouter.PathPrefix("/").Subrouter()
	authUserRouter.Use(h.userJWTAuth)
	authUserRouter.HandleFunc("/user/log-out", h.LogOutUser).Methods(http.MethodPost)
	authUserRouter.HandleFunc("/user/refresh", h.RefreshTokenUser).Methods(http.MethodPost)

	authUserRouter.HandleFunc("/user/{id}", h.FindUserById).Methods(http.MethodGet)
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

func (h *Handler) userJWTAuth(next http.Handler) http.Handler {
    return h.JWTAuth(next, domain.UserRole)
}
