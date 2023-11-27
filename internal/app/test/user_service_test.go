// Package test internal/app/test/user_service_test.go
package test

import (
	_ "errors"
	"reflect"
	"testing"
	"workspace/internal/app/model"
	"workspace/internal/app/repository"
	"workspace/internal/app/service"
)

func TestGetUserByID(t *testing.T) {
	// Arrange

	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)

	expectedUser := &model.User{
		UUID: "1", Username: "JohnDoe", Email: "john@example.com",
	}

	// Act
	actualUser, err := userService.GetUserByID("1")

	// Assert
	if err != nil {
		t.Errorf("Unexpected error: %v", err)
	}

	if !reflect.DeepEqual(actualUser, expectedUser) {
		t.Errorf("Expected user %v, but got %v", expectedUser, actualUser)
	}
}

// Add similar tests for GetUsers, CreateUser, UpdateUser, and DeleteUser
