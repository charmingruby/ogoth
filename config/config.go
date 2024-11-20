package config

import (
	env "github.com/caarlos0/env/v6"
)

type environment struct {
	Environment          string `env:"ENVIRONMENT" envDefault:"dev"`
	ServerHost           string `env:"SERVER_HOST" envDefault:"http://localhost"`
	ServerPort           string `env:"SERVER_PORT" envDefault:"3000"`
	GoogleClientID       string `env:"GOOGLE_CLIENT_ID,required"`
	GoogleClientSecretID string `env:"GOOGLE_CLIENT_SECRET_ID,required"`
}

func New() (Config, error) {
	environment := environment{}
	if err := env.Parse(&environment); err != nil {
		return Config{}, err
	}

	cfg := Config{
		Environment: environment.Environment,
		ServerConfig: serverConfig{
			Host: environment.ServerHost,
			Port: environment.ServerPort,
		},
		GoogleConfig: googleConfig{
			ClientID:       environment.GoogleClientID,
			ClientSecretID: environment.GoogleClientSecretID,
		},
	}

	return cfg, nil
}

type Config struct {
	Environment  string
	ServerConfig serverConfig
	GoogleConfig googleConfig
}

type serverConfig struct {
	Host string
	Port string
}

type googleConfig struct {
	ClientID       string
	ClientSecretID string
}
