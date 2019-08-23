package server

import (
	"log"
	"net/http"

	"github.com/ashon/gotest/exc"
)

type RequestHandler struct {
	HandleRequest func(w http.ResponseWriter, r *http.Request) error
}

func (h RequestHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	error := h.HandleRequest(w, r)

	if error != nil {
		switch e := error.(type) {

		case exc.Error:
			log.Printf("HTTP %d - %s", e.Status(), e)
			http.Error(w, e.Error(), e.Status())

		default:
			http.Error(w,
				http.StatusText(http.StatusInternalServerError),
				http.StatusInternalServerError)
		}
	}
}
