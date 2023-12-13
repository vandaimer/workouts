package config

import (
	"fmt"

	"github.com/spf13/viper"
)

var viperConfig = viper.New()

type Config struct {
	App AppConfig
	Db  DbConfig
}

type AppConfig struct {
	Port string
	Env  string
	Host string
}

type DbConfig struct {
	DB_DBName   string
	DB_Host     string
	DB_Port     string
	DB_Username string
	DB_Password string
}

func MustLoad(filenames ...string) *Config {
	viperConfig.AutomaticEnv()

	viperConfig.SetConfigFile(".env")
	err := viperConfig.ReadInConfig()
	if err != nil {
		fmt.Printf("fatal error config file: %w\n", err)
	}

	var config Config

	viperConfig.Unmarshal(&config.App)
	viperConfig.Unmarshal(&config.Db)

	return &config
}
