package core

type Config struct {
	ListenAddress string
	Routes        map[string](func(*Request) *Response)
}
