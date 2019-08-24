package main

import (
	"github.com/ashon/gotest/app"
	"github.com/ashon/gotest/core"
)

func main() {
	var cfg core.Config
	cfg.ListenAddress = ":8088"
	cfg.Routes = app.Routes

	core.RunServer(cfg)
}
