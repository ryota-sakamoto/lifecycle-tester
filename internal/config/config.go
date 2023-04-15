package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	EnableHealthLog bool `envconfig:"ENABLE_HEALTH_LOG" default:"true"`
}

func GetConfig() (*Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
