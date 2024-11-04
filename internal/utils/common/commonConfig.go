package common

import (
	"log"

	"github.com/spf13/viper"
)

type CommonConfig struct {
	SecretEncKey string
	JWTSecretKey string
	JWTExpiry    string
}

func LoadCommonConfig(configFile string) (*CommonConfig, error) {
	viper.SetConfigName(configFile)
	viper.SetConfigType("env")
	viper.AddConfigPath(".")

	if err := viper.ReadInConfig(); err != nil {
		log.Println("No config file found, proceeding with environment variables only.")
	}

	viper.AutomaticEnv()

	commonConfig := &CommonConfig{
		SecretEncKey: viper.GetString("SECRET_ENC_KEY"),
		JWTSecretKey: viper.GetString("JWT_SECRET_KEY"),
		JWTExpiry:    viper.GetString("JWT_EXPIRY"),
	}
	return commonConfig, nil
}
