package utils

import (
	"encoding/json"
	"net/http"
)

func SetResponse(w http.ResponseWriter, statusCode int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	if err := json.NewEncoder(w).Encode(body); err != nil {
		http.Error(w, "Failed to encode response", http.StatusInternalServerError)
	}
}
