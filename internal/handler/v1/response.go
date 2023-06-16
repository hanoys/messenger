package v1

import (
	"encoding/json"
	"net/http"
)

type errorResponse struct {
	Code    int
	Message string
}

type successResponse struct {
    Message string `json:"data"`
}

func writeError(w http.ResponseWriter, code int, message string) {
	w.WriteHeader(code)
	json.NewEncoder(w).Encode(errorResponse{Code: code, Message: message})
}

func writeSuccess(w http.ResponseWriter, message string) {
    w.WriteHeader(http.StatusOK)
    json.NewEncoder(w).Encode(successResponse{Message: message})
}
