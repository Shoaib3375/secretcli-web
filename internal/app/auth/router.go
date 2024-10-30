package auth

import (
	"github.com/go-chi/chi"
	tmplrndr "github.com/mahinops/secretcli-web/internal/tmpl-rndr"
	"gorm.io/gorm"
)

// RegisterRoutes registers the auth-related routes with the router.
func RegisterAPIRoutes(router chi.Router, db *gorm.DB) {
	// Initialize repository, service, and handler for auth
	authRepo := NewSqlAuthRepository(db)
	authService := NewAuthService(authRepo)
	authHandler := NewAuthHandler(authService, nil)

	// Define auth-related routes
	router.Route("/auth/api", func(r chi.Router) {
		r.Post("/register", authHandler.RegisterUser)
		r.Post("/login", authHandler.LoginUser)
	})
}

// RegisterRoutes registers the auth-related routes with the router.
func RegisterWebRoutes(router chi.Router, db *gorm.DB, renderer *tmplrndr.Renderer) {
	// Initialize repository, service, and handler for auth
	authRepo := NewSqlAuthRepository(db)
	authService := NewAuthService(authRepo)
	authHandler := NewAuthHandler(authService, renderer)

	// Define auth-related routes
	router.Route("/auth/web", func(r chi.Router) {
		r.Get("/register", authHandler.RegisterUserForm)
		r.Get("/login", authHandler.LoginUserForm)
	})
}
