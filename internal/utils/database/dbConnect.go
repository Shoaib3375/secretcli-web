package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mahinops/secretcli-web/model" // Adjust the import path based on your project structure
)

// Connect establishes a connection to the database
func Connect(cfg *Config) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	// Open a connection to the database
	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not open database connection: %w", err)
	}

	// Automatically migrate the schema, creating tables for your models
	if err := db.AutoMigrate(&model.Auth{}, &model.Secret{}); err != nil {
		return nil, fmt.Errorf("could not migrate database schema: %w", err)
	}

	log.Println("Database connection successful and schema migrated!")
	return db, nil
}
