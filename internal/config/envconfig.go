package config

import (
	"github.com/caarlos0/env"
	log "github.com/sirupsen/logrus"
)

type envConfig struct {
	LogLevel string `env:"LOG_LEVEL"`

	ServerPort int    `env:"SERVER_PORT" envDefault:"8080"`
	Version    string `env:"VERSION" envDefault:"v1"`
	BaseUrl    string `env:"BASE_URL"`
}

func newEnvironmentConfig() *envConfig {
	cfg := &envConfig{}
	if err := env.Parse(cfg); err != nil {
		log.Fatal("cannot find configs for server: %v \n", err)
	}
	return cfg
}
