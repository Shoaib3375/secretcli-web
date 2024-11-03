package redisconn

import (
	"log"

	"github.com/spf13/viper"
)

// Config holds the application configuration
type RedisConfig struct {
	RedisHost string
	RedisPort int
}

// LoadConfig reads configuration from the specified file or environment variables
func LoadRedisConfig(configFile string) (*RedisConfig, error) {
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
	redisConfig := &RedisConfig{
		RedisHost: viper.GetString("REDIS_HOST"),
		RedisPort: viper.GetInt("REDIS_PORT"),
	}

	return redisConfig, nil
}
