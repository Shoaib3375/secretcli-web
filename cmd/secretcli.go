package main

import (
	"log"
	"net/http"

	_ "github.com/lib/pq" // Import Postgres driver
	"github.com/mahinops/secretcli-web/app/auth"
	"github.com/mahinops/secretcli-web/utils/database"
)

func main() {
	// Load configuration
	cfg, err := database.LoadConfig()
	if err != nil {
		log.Fatal("Error loading configuration: ", err)
	}

	// Establish database connection
	db, err := database.Connect(cfg)
	if err != nil {
		log.Fatal("Error opening database: ", err)
	}
	defer db.Close()

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
