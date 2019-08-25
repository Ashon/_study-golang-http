package app

import (
	"errors"

	"github.com/ashon/_study-golang-http/core"
)

// Returns simple greetings message
func Hello(*core.Request) *core.Response {
	res := &core.Response{Data: "hello world"}
	return res
}

// Raises 500
func RaisePanic(*core.Request) *core.Response {
	return &core.Response{
		Err: core.StatusError{
			Code: 500,
			Err:  errors.New("server error yeah~~")}}
}

// Unexpected
func UnexpectedPanic(*core.Request) *core.Response {
	panic("force raise panic")
}
