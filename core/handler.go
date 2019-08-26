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
	Logger.Debug("RecoverPanic")
	if r := recover(); r != nil {
		Logger.Error("Server Panic:", r)
		InternalServerError(w)
	}
}

func LogRequest(req *Request, res *Response) {
	Logger.Debug("LogRequest")
	Logger.Info(
		req.Req.Proto,
		req.Req.Method,
		req.Req.URL.Path,
		// req.Req.ContentLength,
		len(res.Data),
		// err,
		req.Req.UserAgent())
}

func (h RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	req := Request{Req: r}

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
	LogRequest(&req, res)
}
