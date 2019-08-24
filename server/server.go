package server

import (
	"fmt"
	"net/http"

	"github.com/ashon/gotest/config"
	"github.com/ashon/gotest/logger"
)

// Build http server and run
func RunServer(cfg config.Config) {
	for route, view := range Routes {
		logger.Error(route, view)
		http.Handle(route, RequestHandler{view})
	}

	logger.Info(fmt.Sprintf("Server listening.. %s", cfg.ListenAddress))
	http.ListenAndServe(cfg.ListenAddress, nil)
}
