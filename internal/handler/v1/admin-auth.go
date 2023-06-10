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

    session, err := h.tokenProvider.NewSession(r.Context(), tokenPayload)
	if err != nil {
		writeError(w, http.StatusInternalServerError, err.Error())
		return
	}

    json.NewEncoder(w).Encode(session.Tokens)
}

func (h *Handler) LogOutAdmin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) RefreshTokenAdmin(w http.ResponseWriter, r *http.Request) {

}

func (h *Handler) adminJWTAuth(next http.Handler) http.Handler {
    return h.JWTAuth(next, domain.AdminRole)
}
