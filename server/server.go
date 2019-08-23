package server

import (
	"net/http"

	"github.com/ashon/gotest/config"
	"github.com/ashon/gotest/views"
)

func RunServer(cfg config.Config) {
	http.Handle("/hello", RequestHandler{views.Hello})
	http.Handle("/panic", RequestHandler{views.RaisePanic})
	http.Handle("/unexc", RequestHandler{views.UnexpectedPanic})

	http.ListenAndServe(cfg.ListenAddress, nil)
}
