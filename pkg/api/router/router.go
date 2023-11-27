// Package router internal/app/router/router.go
package router

import (
	"fmt"
	"net/http"
	_ "workspace/cmd/user-api/docs" // Import the generated Swagger docs
	"workspace/internal/app/handler"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// SetupUserRoutes sets up the user-related routes.
// @title User API
// @version 1.0
// @description This is a sample User API.
// @host localhost:8080
// @BasePath /app/v1
func SetupUserRoutes(r *mux.Router, userHandler *handler.UserHandler) {
	r.HandleFunc("/", helloHandler).Methods("GET").Name("Hello")
	r.HandleFunc("/users", userHandler.GetUsers).Methods("GET").Name("GetUsers")
	r.HandleFunc("/users/{uuid}", userHandler.GetUserByID).Methods("GET").Name("GetUserByID")
	r.HandleFunc("/users", userHandler.CreateUser).Methods("POST").Name("CreateUser")
	r.HandleFunc("/users/{uuid}", userHandler.UpdateUser).Methods("PUT").Name("UpdateUser")
	r.HandleFunc("/users/{uuid}", userHandler.DeleteUser).Methods("DELETE").Name("DeleteUser")

	// Serve Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	))

	// Serve Swagger UI
	r.Handle("/swagger", SwaggerUIHandler())
}

// SwaggerUIHandler returns a handler for serving the Swagger UI.
// @Summary Swagger UI
// @Description Swagger UI for API documentation
// @ID SwaggerUI
// @Produce html
// @Router /swagger [get]
func SwaggerUIHandler() http.Handler {
	return httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"),
	)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}
