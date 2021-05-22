package middleware

import (
	"encoding/json"
	"net/http"
)

type jsonResponse struct {
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
	Success bool        `json:"success"`
}

func success(w http.ResponseWriter, data interface{}, message string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(jsonResponse{
		Data:    data,
		Message: message,
		Success: true,
	})
}

func failure(w http.ResponseWriter, message string, status int) {
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(jsonResponse{
		Data:    nil,
		Message: message,
		Success: false,
	})
}
