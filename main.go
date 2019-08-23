package main

import (
	"github.com/ashon/gotest/config"
	"github.com/ashon/gotest/server"
)

func main() {
	var cfg config.Config
	cfg.ListenAddress = ":8088"

	server.RunServer(cfg)
}
