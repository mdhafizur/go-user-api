// Package handler internal/app/handler/user_handler.go

// @title User API
// @version 1.0
// @description This is the API for managing users.
// @host localhost:8080
// @BasePath /api/v1
// @schemes http

package handler

import (
	"encoding/json"
	"log"
	"net/http"
	apphandler "workspace/pkg/handler"

	"workspace/internal/app/model"
	"workspace/internal/app/service"

	"github.com/gorilla/mux"
)

// UserHandler handles user-related HTTP requests.
type UserHandler struct {
	userService *service.UserService
}

// NewUserHandler creates a new UserHandler.
func NewUserHandler(userService *service.UserService) *UserHandler {
	return &UserHandler{
		userService: userService,
	}
}

// GetUsers
// @Summary Get a list of users
// @Description Get a list of users
// @ID GetUsers
// @Produce json
// @Success 200 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse "Bad Request"
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /api/v1/users [get]
// @Tags users
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetUsers()
	if err != nil {
		log.Printf("Error getting users: %v", err)
		apphandler.RespondWithError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	apphandler.RespondWithJSON(w, users, http.StatusOK, "Users retrieved successfully.")
}

// GetUserByUUID
// @Summary Get user by UUID
// @Description Get a user by UUID
// @ID GetUserByUUID
// @Produce json
// @Param uuid path string true "User UUID"
// @Success 200 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse "Bad Request"
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /api/v1/users/{uuid} [get]
// @Tags users
func (h *UserHandler) GetUserByUUID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userUUID := vars["uuid"]

	user, err := h.userService.GetUserByUUID(userUUID)
	if err != nil {
		log.Printf("Error getting user by UUID %s: %v", userUUID, err)
		apphandler.RespondWithError(w, "User not found", http.StatusNotFound)
		return
	}

	apphandler.RespondWithJSON(w, user, http.StatusOK, "User retrieved successfully.")
}

// CreateUser
// @Summary Create a new user
// @Description Create a new user
// @ID CreateUser
// @Accept json
// @Produce json
// @Success 201 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse "Bad Request"
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /api/v1/users [post]
// @Tags users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		apphandler.RespondWithError(w, "Bad Request", http.StatusBadRequest)
		return
	}

	//Validate the user data
	if validationErrs := user.Validate(); validationErrs != nil {
		// Handle the validation errors, e.g., return an error response to the user
		apphandler.RespondWithError(w, "Validation failed", http.StatusBadRequest, validationErrs)
		return
	}

	createdUser, err := h.userService.CreateUser(user)

	if err != nil {
		log.Printf("Error creating user: %v", err)
		apphandler.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	apphandler.RespondWithJSON(w, createdUser, http.StatusCreated, "User created successfully.")
}

// UpdateUser
// @Summary Update user by UUID
// @Description Update user by UUID
// @ID UpdateUser
// @Accept json
// @Produce json
// @Param uuid path string true "User UUID"
// @Param user body model.User true "Updated user object"
// @Success 200 {object} types.SuccessResponse
// @Failure 400 {object} types.ErrorResponse "Bad Request"
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /api/v1/users/{uuid} [put]
// @Tags users
func (h *UserHandler) UpdateUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userUUID := vars["uuid"]

	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		apphandler.RespondWithError(w, "Bad Request", http.StatusBadRequest)
		return
	}

	updatedUser, err := h.userService.UpdateUser(userUUID, user)
	if err != nil {
		log.Printf("Error updating user with ID %s: %v", userUUID, err)
		apphandler.RespondWithError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	apphandler.RespondWithJSON(w, updatedUser, http.StatusOK, "User updated successfully.")
}

// DeleteUser this handler handles deletion of a user
// @Summary Delete user by UUID
// @Description Delete user by UUID
// @ID DeleteUser
// @Produce json
// @Param uuid path string true "User UUID"
// @Success 204 "No Content"
// @Failure 400 {object} types.ErrorResponse "Bad Request"
// @Failure 500 {object} types.ErrorResponse "Internal Server Error"
// @Router /api/v1/users/{uuid} [delete]
// @Tags users
func (h *UserHandler) DeleteUser(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userUUID := vars["uuid"]

	err := h.userService.DeleteUser(userUUID)
	if err != nil {
		log.Printf("Error deleting user with ID %s: %v", userUUID, err)
		apphandler.RespondWithError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusNoContent)
}
