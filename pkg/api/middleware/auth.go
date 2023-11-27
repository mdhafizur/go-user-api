// Package middleware internal/app/middleware/auth.go
package middleware

import (
	"net/http"
	"os"
)

// AuthMiddleware is a middleware for handling authentication.
func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Perform authentication logic here
		apiKey := r.Header.Get("Authorization")
		expectedAPIKey := os.Getenv("API_KEY")

		if apiKey != expectedAPIKey {
			http.Error(w, "Unauthorized", http.StatusUnauthorized)
			return
		}

		// Call the next handler in the chain
		next.ServeHTTP(w, r)
	})
}
