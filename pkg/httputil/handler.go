package httputil

type HttpHandler func(input interface{}) (output interface{}, err error)
