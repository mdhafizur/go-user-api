// Package service internal/app/service/user_service.go
package service

import (
	"workspace/internal/app/model"
	"workspace/internal/app/repository"
)

// UserService provides business logic related to users.
type UserService struct {
	userRepository *repository.UserRepository
}

// NewUserService creates a new UserService.
func NewUserService(userRepository *repository.UserRepository) *UserService {
	return &UserService{
		userRepository: userRepository,
	}
}

// GetUserByID retrieves a user by ID.
func (s *UserService) GetUserByID(userUUID string) (*model.User, error) {
	return s.userRepository.GetUserByID(userUUID)
}

// GetUsers retrieves all users.
func (s *UserService) GetUsers() ([]model.User, error) {
	return s.userRepository.GetUsers()
}

// CreateUser creates a new user.
func (s *UserService) CreateUser(user model.User) (*model.User, error) {
	return s.userRepository.CreateUser(user)
}

// UpdateUser updates a user.
func (s *UserService) UpdateUser(userUUID string, user model.User) (*model.User, error) {
	return s.userRepository.UpdateUser(userUUID, user)
}

// DeleteUser deletes a user.
func (s *UserService) DeleteUser(userUUID string) error {
	return s.userRepository.DeleteUser(userUUID)
}
