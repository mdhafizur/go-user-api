// Package router pkg/api/router/router.go
package router

import (
	"fmt"
	"net/http"
	"workspace/internal/app/handler"

	"github.com/gorilla/mux"
	httpSwagger "github.com/swaggo/http-swagger"
)

// SetupUserRoutes sets up the user-related routes.
func SetupUserRoutes(r *mux.Router, userHandler *handler.UserHandler) {
	apiRouter := r.PathPrefix("/api").Subrouter()

	// Hello route without the /api prefix
	r.HandleFunc("/", helloHandler).Methods("GET").Name("Hello")

	// User-related routes with the /api prefix
	apiRouter.HandleFunc("/v1/users", userHandler.GetUsers).Methods("GET").Name("GetUsers")
	apiRouter.HandleFunc("/v1/users/{uuid}", userHandler.GetUserByUUID).Methods("GET").Name("GetUserByID")
	apiRouter.HandleFunc("/v1/users", userHandler.CreateUser).Methods("POST").Name("CreateUser")
	apiRouter.HandleFunc("/v1/users/{uuid}", userHandler.UpdateUser).Methods("PUT").Name("UpdateUser")
	apiRouter.HandleFunc("/v1/users/{uuid}", userHandler.DeleteUser).Methods("DELETE").Name("DeleteUser")

	// Serve Swagger documentation
	r.PathPrefix("/swagger/").Handler(httpSwagger.Handler(
		httpSwagger.URL("http://localhost:8080/swagger/doc.json"), // The URL pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	)).Methods(http.MethodGet)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "API is running!")
}
