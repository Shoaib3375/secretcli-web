package cmd

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq" // Import Postgres driver
	"github.com/mahinops/secretcli-web/internal/app/auth"
	"github.com/mahinops/secretcli-web/internal/app/secret"
	tmplrndr "github.com/mahinops/secretcli-web/internal/tmpl-rndr"
	"github.com/mahinops/secretcli-web/internal/utils/database"
	"github.com/mahinops/secretcli-web/internal/utils/web/health"
	"github.com/prometheus/client_golang/prometheus/promhttp"
	"gorm.io/gorm"
)

// App holds the dependencies for the application
type App struct {
	Router *chi.Mux
	db     *gorm.DB
	config *database.Config
}

// NewApp initializes a new App instance
func NewApp(configFile, mode string) (*App, error) {
	cfg := loadConfig(configFile)
	db := connectDatabase(cfg)

	// Create a new router
	router := chi.NewRouter()
	// Initialize the template renderer with the path to your templates
	renderer := tmplrndr.NewRenderer("templates/**/*.tmpl")

	if mode == "api" {
		// Register API-specific routes
		registerAPIRoutes(router, db, cfg)
	} else if mode == "web" {
		// Register Web-specific routes
		registerWebRoutes(router, renderer)
	}

	return &App{Router: router, db: db, config: cfg}, nil
}

// registerAPIRoutes registers only API routes
func registerAPIRoutes(router *chi.Mux, db *gorm.DB, config *database.Config) {
	router.Handle("/metrics", promhttp.Handler())
	router.Get("/health", health.Handler)
	auth.RegisterRoutes(router, db)
	secret.RegisterRoutes(router, db, config)
}

// registerWebRoutes registers only Web routes, including template rendering
func registerWebRoutes(router *chi.Mux, renderer *tmplrndr.Renderer) {
	router.Get("/", func(w http.ResponseWriter, r *http.Request) {
		renderer.Render(w, "index", nil) // Use the renderer instance
	})
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
func (a *App) CloseDatabase() {
	sqlDB, err := a.db.DB()
	if err != nil {
		log.Fatal(err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatal("Error closing database: ", err)
	}
}

// Start the HTTP API server
func (a *App) StartAPIServer(port string) {
	log.Println("Server is running on port " + port + "...")
	if err := http.ListenAndServe(port, a.Router); err != nil {
		log.Fatal(err)
	}
}

// Start the HTTP Web server
func (a *App) StartWebServer(port string) {
	log.Println("Server is running on port " + port + "...")
	if err := http.ListenAndServe(port, a.Router); err != nil {
		log.Fatal(err)
	}
}
