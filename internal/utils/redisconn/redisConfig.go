package redisconn

import (
	"log"

	"github.com/spf13/viper"
)

type RedisConfig struct {
	RedisHost string
	RedisPort int
}

func LoadRedisConfig(configFile string) (*RedisConfig, error) {
	viper.SetConfigName(configFile) // name of the config file (without extension)
	viper.SetConfigType("env")      // type of the configuration file
	viper.AddConfigPath(".")        // look for the file in the working directory

	// Read the config file if it exists
	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config file found, proceeding with environment variables only.")
	}

	viper.AutomaticEnv()

	redisConfig := &RedisConfig{
		RedisHost: viper.GetString("REDIS_HOST"),
		RedisPort: viper.GetInt("REDIS_PORT"),
	}
	return redisConfig, nil
}
