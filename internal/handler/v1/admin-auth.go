package v1

import (
	"encoding/json"
	"net/http"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/domain/dto"
	"github.com/hanoy/messenger/pkg/auth"
)

func (h *Handler) LogInAdmin(w http.ResponseWriter, r *http.Request) {
	w.Header().Add("Content-Type", "application/json")

	var adminDTO dto.LogInAdminDTO
	err := json.NewDecoder(r.Body).Decode(&adminDTO)
	if err != nil {
		writeError(w, http.StatusBadRequest, err.Error())
		return
	}

	admin, err := h.services.Admins.FindByCredentials(r.Context(), adminDTO)
	if err != nil {
		writeError(w, http.StatusBadRequest, "email or password is invalid")
		return
	}

	tokenPayload, err := auth.NewPayload(admin.ID, domain.AdminRole)
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

func (h *Handler) LogOutAdmin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) RefreshTokenAdmin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) adminJWTAuth(next http.Handler) http.Handler {
    return h.JWTAuth(next, domain.AdminRole)
}
