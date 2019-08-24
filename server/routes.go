package server

import (
	"net/http"

	"github.com/ashon/gotest/views"
)

var Routes = map[string](func(http.ResponseWriter, *http.Request) error){
	"/hello": views.Hello,
	"/panic": views.RaisePanic,
	"/unexc": views.UnexpectedPanic,
}
