package core

type ResponseError interface {
	GetError() StatusError
}

type Response struct {
	Data string
	Err  StatusError
}

func (res Response) GetError() StatusError {
	return res.Err
}
