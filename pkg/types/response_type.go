// Package types pkg/types/response_type.go
package types

// ValidationError represents a validation error.
type ValidationError struct {
	Field   string `json:"field"`
	Message string `json:"message"`
}

// ValidationErrors is a collection of validation errors.
type ValidationErrors []ValidationError

// ErrorResponse represents a standard error response.
type ErrorResponse struct {
	Error struct {
		Code    int              `json:"code"`
		Message string           `json:"message"`
		Details ValidationErrors `json:"details,omitempty"`
	} `json:"error"`
}

// SuccessResponse represents the structure of a success response.
type SuccessResponse struct {
	Status  string      `json:"status"`
	Code    int         `json:"code"`
	Data    interface{} `json:"data"`
	Message string      `json:"message"`
}
