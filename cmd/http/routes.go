package main

import (
	handlerPkg "github.com/jekiapp/hi-mod-arch/pkg/handler"
	"net/http"
)

type Handler struct {
	CheckoutPageHandler handlerPkg.GenericHandler
}

func (h Handler) routes(mux *http.ServeMux) {
	mux.HandleFunc("/checkout", handlerPkg.HttpGenericHandler(h.CheckoutPageHandler))
}
