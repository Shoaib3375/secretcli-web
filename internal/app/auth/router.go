package auth

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
	tmplrndr "github.com/mahinops/secretcli-web/internal/tmpl-rndr"
	"github.com/mahinops/secretcli-web/internal/utils/middleware"
	"gorm.io/gorm"
)

// RegisterRoutes registers the auth-related routes with the router.
func RegisterAPIRoutes(router chi.Router, db *gorm.DB, redisClient *redis.Client) {
	// Initialize repository, service, and handler for auth
	authRepo := NewSqlAuthRepository(db)
	authService := NewAuthService(authRepo)
	authHandler := NewAuthHandler(authService, nil, redisClient)

	// Define the rate limiter with a 5-second limit per request
	rateLimiter := middleware.NewRateLimiter(5 * time.Second)

	// Define auth-related routes
	router.Route("/auth/api", func(r chi.Router) {
		r.Use(rateLimiter.LimitMiddleware)
		r.Post("/register", authHandler.RegisterUser)
		r.Post("/login", authHandler.LoginUser)
	})
}

// RegisterRoutes registers the auth-related routes with the router.
func RegisterWebRoutes(router chi.Router, db *gorm.DB, renderer *tmplrndr.Renderer) {
	// Initialize repository, service, and handler for auth
	authRepo := NewSqlAuthRepository(db)
	authService := NewAuthService(authRepo)
	authHandler := NewAuthHandler(authService, renderer, nil)

	// Define auth-related routes
	router.Route("/auth/web", func(r chi.Router) {
		r.Get("/register", authHandler.RegisterUserForm)
		r.Get("/login", authHandler.LoginUserForm)
	})
}
