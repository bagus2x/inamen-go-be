package middleware

import (
	"encoding/json"
	"errors"
	"net/http"

	"github.com/bagus2x/inamen-go-be/pkg/model"
)

type jsonResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

func success(w http.ResponseWriter, data interface{}, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(jsonResponse{
		Data:    data,
		Message: message,
		Success: true,
	})
}

func failure(w http.ResponseWriter, message string, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(jsonResponse{
		Data:    nil,
		Message: message,
		Success: false,
	})
}

func status(err error) int {
	switch errors.Unwrap(err) {
	case model.ErrBadRequest:
		return 400
	case model.ErrUnauthorized:
		return 401
	case model.ErrNotFound:
		return 404
	case model.ErrConflict:
		return 409
	case model.ErrInternalServer:
	}

	return 500
}
