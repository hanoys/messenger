package v1

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Code    int
	Message string
}

func writeError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errorResponse{Code: code, Message: message})
}
