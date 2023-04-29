package config

import (
	"fmt"

	"github.com/caarlos0/env"
)

type Config struct {
	HTTP_port int  `env:"HTTP_PORT" envDefault:"3033"`
	IsProd    bool `env:"IS_PROD" envDefault:"false"`
}

var config Config = Config{}

func GetConfig() (*Config, error) {
	if err := env.Parse(&config); err != nil {
		return nil, fmt.Errorf("read logger configuration failed: %w", err)
	}
	return &config, nil
}
