package config

import "github.com/caarlos0/env/v9"

type Config struct {
	LogLevel string `env:"LOG_LEVEL" envDefault:"debug"`

	CurrentEnvironment string `env:"CURRENT_ENVIRONMENT" envDefault:"TEST"`
	AppPort            int    `env:"APP_PORT" envDefault:"8080"`
}

func NewConfig() (*Config, error) {
	cfg := &Config{}
	if err := env.Parse(cfg); err != nil {
		return nil, err
	}

	return cfg, nil
}
