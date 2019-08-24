package app

import "github.com/ashon/gotest/core"

var Routes = map[string](func(*core.Request) *core.Response){
	"/hello": Hello,
	"/panic": RaisePanic,
	"/unexc": UnexpectedPanic,
}
