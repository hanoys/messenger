package v1

import (
	"encoding/json"
	"fmt"
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

    writeSuccess(w, fmt.Sprintf("user %v registered", user.Login))
}

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

    session, err := h.tokenProvider.NewSession(r.Context(), tokenPayload)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

    json.NewEncoder(w).Encode(session.Tokens)
}

func (h *Handler) LogOutUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    tokenString, err := h.extractToken(r.Header.Get("Authorization"))
    if err != nil {
        writeError(w, http.StatusBadRequest, err.Error())
        return
    }

    err = h.tokenProvider.CloseSession(r.Context(), tokenString)
    if err != nil {
        writeError(w, http.StatusInternalServerError, err.Error())
        return
    }

    writeSuccess(w, "user loged out")
}

func (h *Handler) RefreshTokenUser(w http.ResponseWriter, r *http.Request) {
    w.Header().Set("Content-Type", "application/json")
    tokenString, err := h.extractToken(r.Header.Get("Authorization"))
    if err != nil {
        writeError(w, http.StatusBadRequest, err.Error())
        return
    }

    session, err := h.tokenProvider.RefreshSession(r.Context(), tokenString)
    if err != nil {
        writeError(w, http.StatusInternalServerError, err.Error())
        return
    }

    json.NewEncoder(w).Encode(session.Tokens)
}
