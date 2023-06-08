package v1

import (
	"crypto/sha256"
	"crypto/subtle"
	"errors"
	"net/http"
	"strings"

	"github.com/hanoy/messenger/internal/domain"
	"github.com/hanoy/messenger/internal/domain/dto"
)

func (h *Handler) basicAuth(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		email, password, ok := r.BasicAuth()
		if ok {
			emailHash := sha256.Sum256([]byte(email))
			passwordHash := sha256.Sum256([]byte(password))

			user, err := h.services.Users.FindByEmail(r.Context(), dto.FindByEmailUserDTO{Email: email})
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

func extractToken(authHeader string) (string, error) {
	splitedAuthHeader := strings.Split(authHeader, " ")
	if len(splitedAuthHeader) != 2 {
		return "", errors.New("invalid authorization header")
	}

	return splitedAuthHeader[1], nil
}

func (h *Handler) JWTAuth(next http.Handler, role domain.Role) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		tokenString, err := extractToken(r.Header.Get("Authorization"))
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			writeError(w, http.StatusBadRequest, err.Error())
			return
		}

		tokenPayload, err := h.tokenProvider.VerifyToken(tokenString)
		if err != nil {
			w.Header().Set("Content-Type", "application/json")
			writeError(w, http.StatusUnauthorized, err.Error())
			return
		}

	    if tokenPayload.Role != role {
			w.Header().Set("Content-Type", "application/json")
            writeError(w, http.StatusForbidden, "the resource is forbidden")
            return
        }

		next.ServeHTTP(w, r)
	})
}
