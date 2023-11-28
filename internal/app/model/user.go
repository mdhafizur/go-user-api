// Package model internal/app/model/user.go
package model

import (
	"fmt"
	"github.com/go-playground/validator/v10"
	"workspace/pkg/types"
)

var validate = validator.New()

// User represents a user in the system.
type User struct {
	UUID     string `json:"uuid,omitempty"`
	Username string `json:"username" validate:"required"`
	Email    string `json:"email" validate:"required,email"`
}

type validationErrors types.ValidationErrors

// Validate checks if the User struct is valid.
func (u *User) Validate() types.ValidationErrors {
	if err := validate.Struct(u); err != nil {
		// Convert the validation errors into a custom format
		var validationErrors types.ValidationErrors

		for _, err := range err.(validator.ValidationErrors) {
			validationErrors = append(validationErrors, types.ValidationError{
				Field:   err.Field(),
				Message: fmt.Sprintf("%s is %s", err.Field(), err.Tag()),
			})
		}

		return validationErrors
	}

	return nil
}
