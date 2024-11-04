package auth

import (
	"time"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
	tmplrndr "github.com/mahinops/secretcli-web/internal/tmpl-rndr"
	"github.com/mahinops/secretcli-web/internal/utils/common"
	"github.com/mahinops/secretcli-web/internal/utils/middleware"
	"gorm.io/gorm"
)

func RegisterAPIRoutes(router chi.Router, db *gorm.DB, redisClient *redis.Client, commonConfig *common.CommonConfig) {
	// Initialize repository, service, and handler for auth
	authRepo := NewSqlAuthRepository(db)
	authService := NewAuthService(authRepo)
	authHandler := NewAuthHandler(authService, redisClient, commonConfig)

	rateLimiter := middleware.NewRateLimiter(1 * time.Second)

	router.Route("/auth/api", func(r chi.Router) {
		r.Use(rateLimiter.LimitMiddleware)
		r.Post("/register", authHandler.RegisterUser)
		r.Post("/login", authHandler.LoginUser)
	})
}

func RegisterWebRoutes(router chi.Router, renderer *tmplrndr.Renderer) {
	authWebHandler := NewAuthWebHandler(renderer)

	router.Route("/auth/web", func(r chi.Router) {
		r.Get("/register", authWebHandler.RegisterUserForm)
		r.Get("/login", authWebHandler.LoginUserForm)
	})
}
