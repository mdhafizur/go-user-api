// Package handler internal/app/handler/user_apphandler.go
package handler

import (
	"encoding/json"
	"log"
	"net/http"
	"workspace/pkg/handler"

	"github.com/gorilla/mux"
	"workspace/internal/app/model"
	"workspace/internal/app/service"
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

// @Summary Get a list of users
// @Description Get a list of users
// @ID GetUsers
// @Produce json
// @Success 200 {object} YourResponseType
// @Router /users [get]
// @Tags users
func (h *UserHandler) GetUsers(w http.ResponseWriter, r *http.Request) {
	users, err := h.userService.GetUsers()
	if err != nil {
		log.Printf("Error getting users: %v", err)
		apphandler.RespondWithError(w, "Internal Server Error", http.StatusInternalServerError)
		return
	}

	apphandler.RespondWithJSON(w, users, http.StatusOK)
}

// @Summary Get user by ID
// @Description Get a user by ID
// @ID get-user
// @Produce json
// @Param id path int true "User ID"
// @Success 200 {object} User
// @Router /users/{id} [get]
// @Tags users
func (h *UserHandler) GetUserByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	userUUID := vars["uuid"]

	user, err := h.userService.GetUserByID(userUUID)
	if err != nil {
		log.Printf("Error getting user by UUID %s: %v", userUUID, err)
		apphandler.RespondWithError(w, "User not found", http.StatusNotFound)
		return
	}

	apphandler.RespondWithJSON(w, user, http.StatusOK)
}

// @Summary Create a new user
// @Description Create a new user
// @ID CreateUser
// @Accept json
// @Produce json
// @Param user body model.User true "User object"
// @Success 201 {object} User
// @Router /users [post]
// @Tags users
func (h *UserHandler) CreateUser(w http.ResponseWriter, r *http.Request) {
	var user model.User
	err := json.NewDecoder(r.Body).Decode(&user)
	if err != nil {
		log.Printf("Error decoding request body: %v", err)
		apphandler.RespondWithError(w, "Bad Request", http.StatusBadRequest)
		return
	}

	createdUser, err := h.userService.CreateUser(user)
	if err != nil {
		log.Printf("Error creating user: %v", err)
		apphandler.RespondWithError(w, err.Error(), http.StatusInternalServerError)
		return
	}

	apphandler.RespondWithJSON(w, createdUser, http.StatusCreated)
}

// @Summary Update user by ID
// @Description Update user by ID
// @ID UpdateUser
// @Accept json
// @Produce json
// @Param id path int true "User ID"
// @Param user body model.User true "Updated user object"
// @Success 200 {object} User
// @Router /users/{id} [put]
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

	apphandler.RespondWithJSON(w, updatedUser, http.StatusOK)
}

// @Summary Delete user by ID
// @Description Delete user by ID
// @ID DeleteUser
// @Produce json
// @Param id path int true "User ID"
// @Success 204 "No Content"
// @Router /users/{id} [delete]
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
