// Package apphandler Package handler Package handler response.go
package apphandler

import (
	"encoding/json"
	"net/http"
)

// RespondWithError formats an error response with the given message and status code.
func RespondWithError(w http.ResponseWriter, message string, statusCode int) {
	http.Error(w, message, statusCode)
}

// RespondWithJSON formats a success response with JSON data and the given status code.
func RespondWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
