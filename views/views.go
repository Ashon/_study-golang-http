package views

import (
	"errors"
	"net/http"

	"github.com/ashon/gotest/exc"
)

// Returns simple greetings message
func Hello(w http.ResponseWriter, req *http.Request) error {
	w.Write([]byte("hello world"))

	return nil
}

// Raises Panic
func RaisePanic(w http.ResponseWriter, r *http.Request) error {
	return exc.StatusError{Code: 500, Err: errors.New("exceptions")}
}

// Unexpected
func UnexpectedPanic(w http.ResponseWriter, r *http.Request) error {
	panic("nope")
}
