// Package apphandler pkg/handler/response.go
package apphandler

import (
	"encoding/json"
	"net/http"
	"workspace/pkg/types"
)

// RespondWithError formats an error response with the given message, status code, and optional validation details.
func RespondWithError(w http.ResponseWriter, message string, statusCode int, validationErrors ...types.ValidationErrors) {
	var details types.ValidationErrors
	if len(validationErrors) > 0 {
		details = validationErrors[0]
	}

	errorResponse := types.ErrorResponse{
		Error: struct {
			Code    int                    `json:"code"`
			Message string                 `json:"message"`
			Details types.ValidationErrors `json:"details,omitempty"`
		}{
			Code:    statusCode,
			Message: message,
			Details: details,
		},
	}

	respondWithJSON(w, errorResponse, statusCode)
}

// RespondWithJSON sends a JSON response with the provided data, status code, and message.
func RespondWithJSON(w http.ResponseWriter, data interface{}, statusCode int, message string) {
	successResponse := types.SuccessResponse{
		Status:  "success",
		Code:    statusCode,
		Data:    data,
		Message: message,
	}

	respondWithJSON(w, successResponse, statusCode)
}

func respondWithJSON(w http.ResponseWriter, data interface{}, statusCode int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCode)
	json.NewEncoder(w).Encode(data)
}
