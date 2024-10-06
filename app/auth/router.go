package auth

import (
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

// RegisterRoutes registers the auth-related routes with the router.
func RegisterRoutes(router chi.Router, db *gorm.DB) {
	// Initialize repository, service, and handler for auth
	authRepo := NewSqlAuthRepository(db)
	authService := NewAuthService(authRepo)
	authHandler := NewAuthHandler(authService)

	// Define auth-related routes
	router.Route("/auth", func(r chi.Router) {
		r.Post("/register", authHandler.RegisterUser)
	})
}
