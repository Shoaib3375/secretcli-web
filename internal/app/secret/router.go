package secret

import (
	"time"

	"github.com/go-chi/chi"
	tmplrndr "github.com/mahinops/secretcli-web/internal/tmpl-rndr"
	"github.com/mahinops/secretcli-web/internal/utils/common"
	"github.com/mahinops/secretcli-web/internal/utils/middleware"
	"gorm.io/gorm"
)

func RegisterAPIRoutes(router chi.Router, db *gorm.DB, commonConfig *common.CommonConfig) {
	secretRepo := NewSqlSecretRepository(db)
	secretService := NewSecretService(secretRepo)
	secretHandler := NewSecretHandler(secretService, commonConfig, nil)

	rateLimiter := middleware.NewRateLimiter(5 * time.Second)

	router.Route("/secret/api", func(r chi.Router) {
		r.Use(rateLimiter.LimitMiddleware)
		r.Post("/create", secretHandler.Create)
		r.Get("/list", secretHandler.List)
		r.Post("/generatepassword", secretHandler.GeneratePassword)
		r.Post("/secretdetail", secretHandler.SecretDetail)
	})
}

func RegisterWebRoutes(router chi.Router, db *gorm.DB, commonConfig *common.CommonConfig, renderer *tmplrndr.Renderer) {
	secretRepo := NewSqlSecretRepository(db)
	secretService := NewSecretService(secretRepo)
	secretHandler := NewSecretHandler(secretService, commonConfig, renderer)

	router.Route("/secret/web", func(r chi.Router) {
		r.Get("/list", secretHandler.SecretListTemplate)
		r.Get("/create", secretHandler.SecretCreateForm)
	})
}
