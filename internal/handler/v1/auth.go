package v1

import (
	"crypto/sha256"
	"crypto/subtle"
	"net/http"
)

func (h *Handler) basicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email, password, ok := r.BasicAuth()
		if ok {
			emailHash := sha256.Sum256([]byte(email))
			passwordHash := sha256.Sum256([]byte(password))

			user, err := h.services.Users.FindByEmail(r.Context(), email)
			if err == nil {
				expectedEmailHash := sha256.Sum256([]byte(user.Email))
				expectedPasswordHash := sha256.Sum256([]byte(user.Password))

				emailMatch := subtle.ConstantTimeCompare(emailHash[:],
					expectedEmailHash[:]) == 1
				passwordMatch := subtle.ConstantTimeCompare(passwordHash[:],
					expectedPasswordHash[:]) == 1

				if emailMatch && passwordMatch {
					next.ServeHTTP(w, r)
					return
				}
			}
		}

		w.Header().Set("WWW-Authenticate", `Basic realm="restricted"`)
		w.Header().Set("Content-Type", "application/json")
		writeError(w, http.StatusUnauthorized, "Unauthorized")
	})
}

