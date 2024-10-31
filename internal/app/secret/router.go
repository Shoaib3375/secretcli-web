package secret

import (
	"github.com/go-chi/chi"
	tmplrndr "github.com/mahinops/secretcli-web/internal/tmpl-rndr"
	"github.com/mahinops/secretcli-web/internal/utils/database"
	"gorm.io/gorm"
)

// RegisterRoutes registers the secret-related routes with the router.
func RegisterAPIRoutes(router chi.Router, db *gorm.DB, config *database.Config) {
	// Initialize repository, service, and handler for secrets
	secretRepo := NewSqlSecretRepository(db)
	secretService := NewSecretService(secretRepo)
	secretHandler := NewSecretHandler(secretService, config, nil)

	// Define secret-related routes
	router.Route("/secret/api", func(r chi.Router) {
		r.Post("/create", secretHandler.Create)
		r.Get("/list", secretHandler.List)
	})
}

func RegisterWebRoutes(router chi.Router, db *gorm.DB, config *database.Config, renderer *tmplrndr.Renderer) {
	// Initialize repository, service, and handler for secrets
	secretRepo := NewSqlSecretRepository(db)
	secretService := NewSecretService(secretRepo)
	secretHandler := NewSecretHandler(secretService, config, renderer)

	// Define auth-related routes
	router.Route("/secret/web", func(r chi.Router) {
		r.Get("/list", secretHandler.SecretListTemplate)
		r.Get("/create", secretHandler.SecretCreateForm)
	})
}
