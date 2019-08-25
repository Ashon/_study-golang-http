package app

import "github.com/ashon/_study-golang-http/core"

var Routes = map[string](func(*core.Request) *core.Response){
	"/hello": Hello,
	"/panic": RaisePanic,
	"/unexc": UnexpectedPanic,
}
