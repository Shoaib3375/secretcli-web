package secret

import (
	"github.com/go-chi/chi"
	"github.com/mahinops/secretcli-web/utils/database"
	"gorm.io/gorm"
)

// RegisterRoutes registers the secret-related routes with the router.
func RegisterRoutes(router chi.Router, db *gorm.DB, config *database.Config) {
	// Initialize repository, service, and handler for secrets
	secretRepo := NewSqlSecretRepository(db)
	secretService := NewSecretService(secretRepo)
	secretHandler := NewSecretHandler(secretService, config)

	// Define secret-related routes
	router.Route("/secret", func(r chi.Router) {
		r.Post("/create", secretHandler.Create) // Route to create a secret
	})
}
