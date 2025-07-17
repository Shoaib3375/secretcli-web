package cmd

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
	_ "github.com/lib/pq" // Import Postgres driver
	_ "github.com/mahinops/secretcli-web/docs"
	tmplrndr "github.com/mahinops/secretcli-web/internal/tmpl-rndr"
	"github.com/mahinops/secretcli-web/internal/utils/common"
	"github.com/mahinops/secretcli-web/internal/utils/database"
	"github.com/mahinops/secretcli-web/internal/utils/redisconn"
	"gorm.io/gorm"
)

type App struct {
	Router       *chi.Mux
	db           *gorm.DB
	commonConfig *common.CommonConfig
}

func NewApp(configFile, mode string) (*App, error) {
	cfg := loadDBConfig(configFile)
	db := connectDatabase(cfg)
	redisCfg := loadRedisConfig(configFile)
	redisClient := redisconn.ConnectRedis(redisCfg)
	commonCfg := loadCommonConfig(configFile)

	router := chi.NewRouter()

	// CORS configuration
	corsOptions := cors.New(cors.Options{
		AllowedOrigins:   []string{"http://localhost:8081", "http://localhost:5173"}, // Allow your frontend origin
		AllowedMethods:   []string{"GET", "POST", "OPTIONS"},
		AllowedHeaders:   []string{"Content-Type", "Authorization"}, // Make sure to include any custom headers
		ExposedHeaders:   []string{"Authorization"},
		AllowCredentials: true,
	})

	// Use CORS middleware
	router.Use(corsOptions.Handler)

	if mode == "api" {
		RegisterAPIRoutes(router, db, commonCfg, redisClient)
	} else if mode == "web" {
		renderer := tmplrndr.NewRenderer("templates/**/*.tmpl")
		router.Handle("/static/*", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))
		RegisterWebRoutes(router, renderer)
	}
	return &App{Router: router, db: db, commonConfig: commonCfg}, nil
}

func loadDBConfig(configFile string) *database.Config {
	cfg, err := database.LoadConfig(configFile)
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}
	return cfg
}

func loadRedisConfig(configFile string) *redisconn.RedisConfig {
	redisCfg, err := redisconn.LoadRedisConfig(configFile)
	if err != nil {
		log.Fatal("Error loading redis configuration: ", err)
	}
	return redisCfg
}

func loadCommonConfig(configFile string) *common.CommonConfig {
	commonCfg, err := common.LoadCommonConfig(configFile)
	if err != nil {
		log.Fatal("Error loading common configuration: ", err)
	}
	return commonCfg
}

func connectDatabase(cfg *database.Config) *gorm.DB {
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	return db
}

func (a *App) CloseDatabase() {
	sqlDB, err := a.db.DB()
	if err != nil {
		log.Fatal(err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatal("Error closing database: ", err)
	}
}
