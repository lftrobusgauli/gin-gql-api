package config

import (
	"github.com/caarlos0/env/v6"
)

type DBConfig struct {
	Username string `env:"DB_USERNAME" envDefault:"pg"`
	Password string `env:"DB_PASSWORD" envDefault:"password"`
	Database string `env:"DB_NAME" envDefault:"thegame"`
	DbPort   string `env:"DB_PORT" envDefault:"5432"`
}

type AppConfig struct {
	// Configuration for database storage
	DBConfig DBConfig
}

// New returns instance of AppConfig that exposes the method to load the configuation from environment.
func New() *AppConfig {
	return &AppConfig{}
}

// LoadFromEnv loads the configuration into `AppConfig` from environment.
func (config *AppConfig) LoadFromEnv() error {
	// Load the configuration from envinment variables and inject those values
	// within the `AppConfig` struct.
	return env.Parse(config)
}
