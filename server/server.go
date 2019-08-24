package server

import (
	"net/http"

	"github.com/ashon/gotest/config"
	"github.com/ashon/gotest/logger"
)

// Build http server and run
func RunServer(cfg config.Config) {
	for route, view := range Routes {
		logger.Info(route, view)
		http.Handle(route, RequestHandler{view})
	}

	http.ListenAndServe(cfg.ListenAddress, nil)
}
