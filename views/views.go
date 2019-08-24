package views

import (
	"errors"

	"github.com/ashon/gotest/exc"
	"github.com/ashon/gotest/request"
	"github.com/ashon/gotest/response"
)

// Returns simple greetings message
func Hello(*request.Request) *response.Response {
	res := &response.Response{Data: "hello world"}
	return res
}

// Raises Panic
func RaisePanic(*request.Request) *response.Response {
	return &response.Response{
		Err: exc.StatusError{
			Code: 500, Err: errors.New("server error yeah~~")}}
}

// Unexpected
func UnexpectedPanic(*request.Request) *response.Response {
	panic("force raise panic")
}
