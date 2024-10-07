package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq" // Import Postgres driver
	"github.com/mahinops/secretcli-web/app/auth"
	"github.com/mahinops/secretcli-web/app/secret"
	"github.com/mahinops/secretcli-web/utils/database"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"
)

// App holds the dependencies for the application
type App struct {
	router *chi.Mux
	db     *gorm.DB
	config *database.Config
}

// NewApp initializes a new App instance
func NewApp(configFile string) (*App, error) {
	cfg := loadConfig(configFile)
	db := connectDatabase(cfg)

	// Create a new router
	router := chi.NewRouter()

	// Register Prometheus metrics endpoint (default system metrics)
	router.Handle("/metrics", promhttp.Handler())

	// Register routes
	registerRoutes(router, db, cfg)

	return &App{router: router, db: db, config: cfg}, nil
}

// Load configuration
func loadConfig(configFile string) *database.Config {
	cfg, err := database.LoadConfig(configFile)
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}
	return cfg
}

// Connect to the database
func connectDatabase(cfg *database.Config) *gorm.DB {
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	return db
}

// Close the database connection
func (a *App) closeDatabase() {
	sqlDB, err := a.db.DB()
	if err != nil {
		log.Fatal(err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatal("Error closing database: ", err)
	}
}

// Start the HTTP server
func (a *App) startServer(port string) {
	log.Println("Server is running on port " + port + "...")
	if err := http.ListenAndServe(port, a.router); err != nil {
		log.Fatal(err)
	}
}

// Register application routes
func registerRoutes(router *chi.Mux, db *gorm.DB, config *database.Config) {
	// Register auth-related routes from the auth package
	auth.RegisterRoutes(router, db)
	secret.RegisterRoutes(router, db, config)
}

// Main entry point
func main() {
	app, err := NewApp(".env")
	if err != nil {
		log.Fatal(err)
	}
	defer app.closeDatabase()

	// Start the server
	app.startServer(":8080")
}
