package v1

import (
	"encoding/json"
	"net/http"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/domain/dto"
	"github.com/hanoy/messenger/pkg/auth"
)

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

	tokenPayload, err := auth.NewPayload(user.ID, string(domain.UserRole))
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

    tokenPayload, err := auth.NewPayload(user.ID, domain.UserRole)
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
		Token: tokenString,
	}

	json.NewEncoder(w).Encode(token)

}

func (h *Handler) LogOutUser(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) RefreshTokenUser(w http.ResponseWriter, r *http.Request) {

}
