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
	authHandler := NewAuthHandler(authService, nil, redisClient, commonConfig)

	rateLimiter := middleware.NewRateLimiter(5 * time.Second)

	router.Route("/auth/api", func(r chi.Router) {
		r.Use(rateLimiter.LimitMiddleware)
		r.Post("/register", authHandler.RegisterUser)
		r.Post("/login", authHandler.LoginUser)
	})
}

func RegisterWebRoutes(router chi.Router, db *gorm.DB, renderer *tmplrndr.Renderer) {
	authRepo := NewSqlAuthRepository(db)
	authService := NewAuthService(authRepo)
	authHandler := NewAuthHandler(authService, renderer, nil, nil)

	router.Route("/auth/web", func(r chi.Router) {
		r.Get("/register", authHandler.RegisterUserForm)
		r.Get("/login", authHandler.LoginUserForm)
	})
}
