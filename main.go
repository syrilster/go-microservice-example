package main

import (
	"github.com/syrilster/go-microservice-example/internal"
	"github.com/syrilster/go-microservice-example/internal/config"
)

func main() {
	cfg := config.NewApplicationConfig()
	server := internal.SetupServer(cfg)
	server.Start("", cfg.ServerPort())
}
