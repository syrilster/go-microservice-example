package internal

import (
	"fmt"
	"github.com/syrilster/go-microservice-example/internal/config"
	"github.com/syrilster/go-microservice-example/internal/middlewares"
	"net/http"
)

func StatusRoute() (route config.Route) {
	route = config.Route{
		Path:    "/status",
		Method:  http.MethodGet,
		Handler: middlewares.RuntimeHealthCheck(),
	}
	return route
}

type ServerConfig interface {
	Version() string
	BaseURL() string
}

func SetupServer(cfg ServerConfig) *config.Server {
	basePath := fmt.Sprintf("/%v", cfg.Version())
	server := config.NewServer().
		WithRoutes(
			"", StatusRoute(),
		).
		WithRoutes(
			basePath,
		)
	return server
}
