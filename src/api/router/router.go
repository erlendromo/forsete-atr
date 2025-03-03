package router

import (
	"net/http"

	"github.com/erlendromo/forsete-atr/src/api/handler"
)

type Router interface {
	Serve() error
}

// TODO Add endpoints
func WithEndpoints(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("GET /forsete-atr/v1/yaml/", handler.GetYaml)
	mux.HandleFunc("GET /forsete-atr/v1/basic/", handler.GetBasic)

	return mux
}
