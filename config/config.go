package config

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	ServiceName   string `mapstructure:"SERVICE_NAME"`
	ServerAddress string `mapstructure:"SERVER_ADDRESS"`
	ServerPort    string `mapstructure:"SERVER_PORT"`
	MongoHostname string `mapstructure:"MONGO_HOSTNAME"`
	MongoPort     string `mapstructure:"MONGO_PORT"`
	DatabaseName  string `mapstructure:"DATABASE_NAME"`
}

func LoadConfig() (config Config, err error) {
	viper.AutomaticEnv()
	if err = viper.Unmarshal(&config); err != nil {
		return Config{}, err
	}
	return
}

func LocalLoadConfig() (config Config, err error) {
	viper.SetConfigType("env")
	viper.SetConfigName(".env")
	viper.AddConfigPath("./config")

	if err = viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading .env file: %s", err)
	}
	if err = viper.Unmarshal(&config); err != nil {
		log.Fatalf("Error unmarshalling .env file: %s", err)
	}
	return
}
