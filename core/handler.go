package core

import (
	"fmt"
	"net/http"
)

type RequestHandler struct {
	HandleRequest func(*Request) *Response
}

func HandleError(w http.ResponseWriter, r *http.Request, err error) {
	switch e := err.(type) {

	case Error:
		Logger.Error(fmt.Sprintf("HTTP %d - %s", e.Status(), e))
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
		Logger.Error("Server Panic:", r)
		InternalServerError(w)
	}
}

func (h RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := Request{Req: r}

	Logger.Info(
		req.Req.Proto,
		req.Req.Method,
		req.Req.URL.Path,
		req.Req.ContentLength,
		req.Req.UserAgent())

	// TODO: Request Middleware
	// Pre-process for request

	// Handle Request
	defer RecoverPanic(w)
	res := h.HandleRequest(&req)

	// Handler Error
	responseError := res.GetError()
	if responseError.Err != nil {
		HandleError(w, r, responseError)
	}

	// Response
	w.Write([]byte(res.Data))
}
