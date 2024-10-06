package main

import (
	"log"
	"net/http"

	"github.com/go-chi/chi"
	_ "github.com/lib/pq" // Import Postgres driver
	"github.com/mahinops/secretcli-web/app/auth"
	"github.com/mahinops/secretcli-web/utils/database"
	"gorm.io/gorm"
)

func main() {
	cfg := loadConfig()
	db := connectDatabase(cfg)
	defer closeDatabase(db)

	// Create a new router
	router := chi.NewRouter()

	// Register routes
	registerRoutes(router, db)

	// Start the server
	startServer(router)
}

// Load configuration
func loadConfig() *database.Config {
	cfg, err := database.LoadConfig()
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
func closeDatabase(db *gorm.DB) {
	sqlDB, err := db.DB()
	if err != nil {
		log.Fatal(err)
	}
	if err := sqlDB.Close(); err != nil {
		log.Fatal("Error closing database: ", err)
	}
}

// Start the HTTP server
func startServer(router http.Handler) {
	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}

// Register application routes
func registerRoutes(router *chi.Mux, db *gorm.DB) {
	// Register auth-related routes from the auth package
	auth.RegisterRoutes(router, db)
}
