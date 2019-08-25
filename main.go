package main

import (
	"github.com/ashon/_study-golang-http/app"
	"github.com/ashon/_study-golang-http/core"
)

func main() {
	var cfg core.Config
	cfg.ListenAddress = ":8088"
	cfg.Routes = app.Routes

	core.RunServer(cfg)
}
