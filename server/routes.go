package server

import (
	"github.com/ashon/gotest/request"
	"github.com/ashon/gotest/response"
	"github.com/ashon/gotest/views"
)

var Routes = map[string](func(*request.Request) *response.Response){
	"/hello": views.Hello,
	"/panic": views.RaisePanic,
	"/unexc": views.UnexpectedPanic,
}
