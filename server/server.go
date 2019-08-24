package server

import (
	"net/http"

	"github.com/ashon/gotest/config"
	"github.com/ashon/gotest/views"
)

func RunServer(cfg config.Config) {
	routes := map[string](func(http.ResponseWriter, *http.Request) error){
		"/hello": views.Hello,
		"/panic": views.RaisePanic,
		"/unexc": views.UnexpectedPanic,
	}

	for route, view := range routes {
		http.Handle(route, RequestHandler{view})
	}

	http.ListenAndServe(cfg.ListenAddress, nil)
}
