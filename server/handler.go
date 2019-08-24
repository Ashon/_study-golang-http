package server

import (
	"fmt"
	"net/http"

	"github.com/ashon/gotest/exc"
	"github.com/ashon/gotest/logger"
	"github.com/ashon/gotest/request"
	"github.com/ashon/gotest/response"
)

type RequestHandler struct {
	HandleRequest func(*request.Request) *response.Response
}

func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	switch e := err.(type) {

	case exc.Error:
		logger.Error(fmt.Sprintf("HTTP %d - %s", e.Status(), e))
		http.Error(w, e.Error(), e.Status())

	default:
		InternalServerError(w)
	}
}

func InternalServerError(w http.ResponseWriter) {
	http.Error(w,
		http.StatusText(http.StatusInternalServerError),
		http.StatusInternalServerError)
}

func RecoverPanic(w http.ResponseWriter) {
	if r := recover(); r != nil {
		logger.Error("Server Panic:", r)
		InternalServerError(w)
	}
}

func (h RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := request.Request{Req: r}

	logger.Info(
		req.Req.Proto,
		req.Req.Method,
		req.Req.URL.Path,
		req.Req.ContentLength,
		req.Req.UserAgent())

	// Request Middleware

	// Handle Request
	defer RecoverPanic(w)
	res := h.HandleRequest(&req)

	// Handler Error
	responseError := res.GetError()
	if responseError.Err != nil {
		HandleError(w, r, responseError)
	}

	// response
	w.Write([]byte(res.Data))
}
