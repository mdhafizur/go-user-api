// Package test internal/app/test/user_repository_test.go
package test

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"workspace/internal/app/model"
	"workspace/internal/app/repository"
	"workspace/pkg/database"
)

func TestUserRepository_GetUserByID(t *testing.T) {
	setup()
	defer teardown()

	// Arrange
	userRepo := repository.NewUserRepository()
	userRepo.DeleteUser("1")

	// Insert a user into the database for testing
	userToInsert := model.User{UUID: "1", Username: "JohnDoe", Email: "john@example.com"}
	insertedUser, err := userRepo.CreateUser(userToInsert)
	assert.NoError(t, err)

	// Act
	retrievedUser, err := userRepo.GetUserByID(insertedUser.UUID)

	// Assert
	assert.NoError(t, err)
	assert.NotNil(t, retrievedUser)
	assert.Equal(t, insertedUser.UUID, retrievedUser.UUID)
	assert.Equal(t, insertedUser.Username, retrievedUser.Username)
	assert.Equal(t, insertedUser.Email, retrievedUser.Email)

}

// Similar tests for GetUsers, CreateUser, UpdateUser, and DeleteUser

func setup() {
	// Setup your test environment (e.g., initialize a test database connection)
	db.Initialize("mongodb+srv://admin:admin@testing.lzovael.mongodb.net/")
}

func teardown() {
	// Teardown your test environment (e.g., close database connections, clean up resources)
	//db.Disconnect()
}
