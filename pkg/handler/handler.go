package handler

type GenericHandler interface {
	HandlerFunc(input interface{}) (output interface{}, err error)
	ObjectAddress() interface{}
}
