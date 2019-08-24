package response

import (
	"github.com/ashon/gotest/exc"
)

type ResponseError interface {
	GetError() exc.StatusError
}

type Response struct {
	Data string
	Err  exc.StatusError
}

func (res Response) GetError() exc.StatusError {
	return res.Err
}
