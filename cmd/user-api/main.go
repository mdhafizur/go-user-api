// Package main cmd/user-api/main.go
package main

import (
	"fmt"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"

	"workspace/internal/app/handler"
	"workspace/internal/app/repository"
	"workspace/internal/app/service"
	_ "workspace/pkg/api/middleware"
	"workspace/pkg/api/router"
	db "workspace/pkg/database"

	_ "workspace/docs" // Import the generated Swagger docs
)

// @title Go User API
// @version 1.0
// @description This is a sample go user api server.
// @termsOfService http://swagger.io/terms/

// @contact.name Md Hafizur Rahman
// @contact.url <a href="mailto:hafizur.upm@gmail.com">Reach out</a>
// @contact.email hafizur.upm@gmail.com

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host
// @BasePath /
func main() {
	// Load environment variables
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	// Initialize the database connection
	db.Initialize()

	// Initialize repositories and services
	userRepo := repository.NewUserRepository()
	userService := service.NewUserService(userRepo)

	r := mux.NewRouter()

	// Initialize router and middleware
	//r.Use(middleware.AuthMiddleware)

	// Initialize API handlers
	userHandler := handler.NewUserHandler(userService)

	// Define API routes
	router.SetupUserRoutes(r, userHandler)

	// Start the HTTP server
	port := os.Getenv("GO_API_PORT")
	if port == "" {
		port = "8080"
	}

	serverAddr := fmt.Sprintf(":%s", port)
	fmt.Printf("Server listening on %s...\n", serverAddr)
	log.Fatal(http.ListenAndServe(serverAddr, r))
}
