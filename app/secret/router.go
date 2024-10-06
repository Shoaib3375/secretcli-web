package secret

import (
	"github.com/go-chi/chi"
	"gorm.io/gorm"
)

// RegisterRoutes registers the secret-related routes with the router.
func RegisterRoutes(router chi.Router, db *gorm.DB) {
	// Initialize repository, service, and handler for secrets
	secretRepo := NewSqlSecretRepository(db)
	secretService := NewSecretService(secretRepo)
	secretHandler := NewSecretHandler(secretService)

	// Define secret-related routes
	router.Route("/secret", func(r chi.Router) {
		r.Post("/create", secretHandler.Create) // Route to create a secret
	})
}
