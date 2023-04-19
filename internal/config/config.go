package config

import (
	"github.com/kelseyhightower/envconfig"
)

type Config struct {
	DisableHealthLog      bool  `envconfig:"DISABLE_HEALTH_LOG" default:"false"`
	ShutdownDelaySeconds  int64 `envconfig:"SHUTDOWN_DELAY_SECONDS" default:"0"`
	ReadinessDelaySeconds int64 `envconfig:"READINESS_DELAY_SECONDS" default:"0"`
	LivenessDelaySeconds  int64 `envconfig:"LIVENESS_DELAY_SECONDS" default:"0"`
}

func GetConfig() (*Config, error) {
	var c Config
	if err := envconfig.Process("", &c); err != nil {
		return nil, err
	}

	return &c, nil
}
