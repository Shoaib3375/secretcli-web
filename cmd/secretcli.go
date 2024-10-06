package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq" // Import Postgres driver
	"github.com/mahinops/secretcli-web/app/auth"
	"github.com/mahinops/secretcli-web/utils/database"
	"gorm.io/gorm"
)

func main() {
	cfg := loadConfig()
	db := connectDatabase(cfg)
	defer closeDatabase(db)

	// Initialize repository, service, and handler
	authRepo := auth.NewSqlAuthRepository(db)
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	// Define routes
	http.HandleFunc("/register", authHandler.RegisterUser)

	// Start the server
	startServer()
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
func startServer() {
	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
