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

// LoadConfig reads configuration from the specified file or environment variables
func LoadConfig(configFile string) (*Config, error) {
	// Set the config file name based on the flag passed
	viper.SetConfigName(configFile) // name of the config file (without extension)
	viper.SetConfigType("env")      // type of the configuration file
	viper.AddConfigPath(".")        // look for the file in the working directory

	// Read the config file if it exists
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config file found, proceeding with environment variables only.")
	}

	// Automatically read environment variables
	viper.AutomaticEnv()

	// Load the configuration values
	config := &Config{
		DBHost:        viper.GetString("DB_HOST"),
		DBPort:        viper.GetInt("DB_PORT"),
		DBUser:        viper.GetString("DB_USER"),
		DBPassword:    viper.GetString("DB_PASSWORD"),
		DBName:        viper.GetString("DB_NAME"),
		EncryptionKey: viper.GetString("ENCRYPTION_KEY"), // Encryption key
	}

	return config, nil
}
