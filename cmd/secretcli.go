package main

import (
	"database/sql"
	"log"
	"net/http"
	"strconv"

	_ "github.com/lib/pq" // Import Postgres driver
	"github.com/mahinops/secretcli-web/app/auth"
	"github.com/mahinops/secretcli-web/config"
)

func main() {

	// Load configuration
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	// Create the connection string using the cfg variable
	connStr := "host=" + cfg.DBHost +
		" port=" + strconv.Itoa(cfg.DBPort) +
		" user=" + cfg.DBUser +
		" password=" + cfg.DBPassword +
		" dbname=" + cfg.DBName +
		" sslmode=disable"

	// Database connection (assuming you have Postgres running)
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

	// Test the connection
	if err := db.Ping(); err != nil {
		log.Fatal("Database connection failed: ", err)
	}

	log.Println("Database connection successful!")

	// Initialize repository, service, and handler
	authRepo := auth.NewSqlAuthRepository(db)
	authService := auth.NewAuthService(authRepo)
	authHandler := auth.NewAuthHandler(authService)

	// Define routes
	http.HandleFunc("/register", authHandler.RegisterUser)

	// Start the server
	log.Println("Server is running on port 8080...")
	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}
}
