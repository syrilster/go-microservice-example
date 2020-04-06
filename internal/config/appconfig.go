package config

import (
	"github.com/syrilster/go-microservice-example/internal/currencyexchange"
	"github.com/syrilster/go-microservice-example/internal/customhttp"
	"net/http"
	"time"
)

type ApplicationConfig struct {
	envValues              *envConfig
	currencyExchangeClient currencyexchange.ClientInterface
}

func (cfg *ApplicationConfig) Version() string {
	return cfg.envValues.Version
}

func (cfg *ApplicationConfig) ServerPort() int {
	return cfg.envValues.ServerPort
}

func (cfg *ApplicationConfig) BaseURL() string {
	return cfg.envValues.BaseUrl
}

func (cfg *ApplicationConfig) CurrencyExchangeClient() currencyexchange.ClientInterface {
	return cfg.currencyExchangeClient
}

func NewApplicationConfig() *ApplicationConfig {
	envValues := newEnvironmentConfig()
	httpCommand := NewHttpCommand()
	ceClient := currencyexchange.NewClient(envValues.CurrencyExchangeEndpoint, httpCommand)

	return &ApplicationConfig{
		envValues:              envValues,
		currencyExchangeClient: ceClient,
	}
}

func NewHttpCommand() customhttp.HTTPCommand {
	httpCommand := customhttp.New(
		customhttp.WithHTTPClient(&http.Client{Timeout: 5 * time.Second}),
	).Build()

	return httpCommand
}
