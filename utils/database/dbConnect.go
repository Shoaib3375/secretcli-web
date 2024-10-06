package database

import (
	"database/sql"
	"log"
	"strconv"

	_ "github.com/lib/pq" // Import Postgres driver
)

// Connect establishes a connection to the database using the given configuration
func Connect(cfg *Config) (*sql.DB, error) {
	// Create the connection string
	connStr := "host=" + cfg.DBHost +
		" port=" + strconv.Itoa(cfg.DBPort) +
		" user=" + cfg.DBUser +
		" password=" + cfg.DBPassword +
		" dbname=" + cfg.DBName +
		" sslmode=disable"

	// Database connection
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, err
	}

	// Test the connection
	if err := db.Ping(); err != nil {
		return nil, err
	}

	log.Println("Database connection successful!")
	return db, nil
}
