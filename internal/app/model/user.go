// Package model internal/models/user.go
package model

// User represents a user in the system.
type User struct {
	UUID     string `json:"uuid,omitempty"`
	Username string `json:"username"`
	Email    string `json:"email"`
	// Add other fields as needed
}
