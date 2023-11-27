// main.go
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
	"workspace/pkg/api/router"
	"workspace/pkg/database"
)

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

	// Initialize router and middleware
	r := mux.NewRouter()
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
