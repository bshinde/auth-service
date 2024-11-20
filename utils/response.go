package utils

import (
	"encoding/json"
	"net/http"
)

// RespondWithJSON sends a JSON response to the client.
func RespondWithJSON(w http.ResponseWriter, statusCode int, payload interface{}) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(payload)
}

// RespondWithError sends a JSON error response to the client.
func RespondWithError(w http.ResponseWriter, statusCode int, errorMessage string) {
	RespondWithJSON(w, statusCode, map[string]string{"error": errorMessage})
}
