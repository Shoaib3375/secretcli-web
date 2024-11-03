package secret

import (
	"time"

	"github.com/go-chi/chi"
	tmplrndr "github.com/mahinops/secretcli-web/internal/tmpl-rndr"
	"github.com/mahinops/secretcli-web/internal/utils/database"
	"github.com/mahinops/secretcli-web/internal/utils/middleware"
	"gorm.io/gorm"
)

// RegisterRoutes registers the secret-related routes with the router.
func RegisterAPIRoutes(router chi.Router, db *gorm.DB, config *database.Config) {
	// Initialize repository, service, and handler for secrets
	secretRepo := NewSqlSecretRepository(db)
	secretService := NewSecretService(secretRepo)
	secretHandler := NewSecretHandler(secretService, config, nil)

	// Define the rate limiter with a 5-second limit per request
	rateLimiter := middleware.NewRateLimiter(5 * time.Second)

	// Define secret-related routes
	router.Route("/secret/api", func(r chi.Router) {
		r.Use(rateLimiter.LimitMiddleware)
		r.Post("/create", secretHandler.Create)
		r.Get("/list", secretHandler.List)
		r.Post("/generatepassword", secretHandler.GeneratePassword)
		r.Post("/secretdetail", secretHandler.SecretDetail)
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
