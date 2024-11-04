package database

import (
	"fmt"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"

	"github.com/mahinops/secretcli-web/model"
)

func Connect(cfg *Config) (*gorm.DB, error) {
	connStr := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		cfg.DBHost, cfg.DBPort, cfg.DBUser, cfg.DBPassword, cfg.DBName)

	db, err := gorm.Open(postgres.Open(connStr), &gorm.Config{})
	if err != nil {
		return nil, fmt.Errorf("could not open database connection: %w", err)
	}

	if err := db.AutoMigrate(&model.Auth{}, &model.Secret{}); err != nil {
		return nil, fmt.Errorf("could not migrate database schema: %w", err)
	}

	log.Println("Database connection successful and schema migrated!")
	return db, nil
}
