package core

import (
	"fmt"
	"net/http"
)

// Build http server and run
func RunServer(cfg Config) {
	for route, view := range cfg.Routes {
		Logger.Error(route, view)
		http.Handle(route, RequestHandler{view})
	}

	Logger.Info(fmt.Sprintf("Server listening.. %s", cfg.ListenAddress))
	http.ListenAndServe(cfg.ListenAddress, nil)
}
