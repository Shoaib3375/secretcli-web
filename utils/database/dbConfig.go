package database

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	DBHost        string
	DBPort        int
	DBUser        string
	DBPassword    string
	DBName        string
	EncryptionKey string
}

// LoadConfig reads configuration from .env file or environment variables
func LoadConfig() (*Config, error) {
	viper.SetConfigName(".env") // name of the .env file (without extension)
	viper.SetConfigType("env")  // type of the configuration file
	viper.AddConfigPath(".")    // optionally look for the .env file in the working directory

	// Read the .env file if it exists
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No .env file found, proceeding with environment variables only.")
	}

	// Automatically read environment variables
	viper.AutomaticEnv()

	config := &Config{
		DBHost:        viper.GetString("DB_HOST"),
		DBPort:        viper.GetInt("DB_PORT"),
		DBUser:        viper.GetString("DB_USER"),
		DBPassword:    viper.GetString("DB_PASSWORD"),
		DBName:        viper.GetString("DB_NAME"),
		EncryptionKey: viper.GetString("ENCRYPTION_KEY"), // Add this line
	}

	return config, nil
}
