package database

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type Config struct {
	DBHost     string
	DBPort     int
	DBUser     string
	DBPassword string
	DBName     string
}

func LoadConfig(configFile string) (*Config, error) {
	viper.SetConfigName(configFile) // name of the config file (without extension)
	viper.SetConfigType("env")      // type of the configuration file
	viper.AddConfigPath(".")        // look for the file in the working directory

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config file found, proceeding with environment variables only.")
	}

	viper.AutomaticEnv()

	config := &Config{
		DBHost:     viper.GetString("DB_HOST"),
		DBPort:     viper.GetInt("DB_PORT"),
		DBUser:     viper.GetString("DB_USER"),
		DBPassword: viper.GetString("DB_PASSWORD"),
		DBName:     viper.GetString("DB_NAME"),
	}
	return config, nil
}
