package config

import (
	"log"
	"github.com/spf13/viper"
)

type Configuration struct {
	Server   ServerConfiguration
	MQServer MQConfiguration
}

func InitConfig() Configuration {
	viper.SetConfigName(".env")
	viper.AddConfigPath(".")

	var configuration Configuration

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("Error reading config file, %s", err)
	}
	err := viper.Unmarshal(&configuration)
	if err != nil {
		log.Fatalf("unable to decode into struct, %v", err)
	}

	return configuration
}
