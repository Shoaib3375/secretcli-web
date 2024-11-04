package cmd

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-redis/redis/v8"
	"github.com/mahinops/secretcli-web/internal/app/auth"
	"github.com/mahinops/secretcli-web/internal/app/secret"
	"github.com/mahinops/secretcli-web/internal/utils/common"
	"github.com/mahinops/secretcli-web/internal/utils/health"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	httpSwagger "github.com/swaggo/http-swagger"
	"gorm.io/gorm"
)

func RegisterAPIRoutes(router *chi.Mux, db *gorm.DB, commonConfig *common.CommonConfig, redisClient *redis.Client) {
	router.Get("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("/swagger/doc.json"), // The url pointing to API definition
		httpSwagger.DeepLinking(true),
		httpSwagger.DocExpansion("none"),
		httpSwagger.DomID("swagger-ui"),
	))
	router.Handle("/metrics", promhttp.Handler())
	router.Get("/health", health.HealthCheck)
	auth.RegisterAPIRoutes(router, db, redisClient, commonConfig)
	secret.RegisterAPIRoutes(router, db, commonConfig)
}

func (a *App) StartAPIServer(port string) {
	log.Println("Server is running on port " + port + "...")
	if err := http.ListenAndServe(port, a.Router); err != nil {
		log.Fatal(err)
	}
}
