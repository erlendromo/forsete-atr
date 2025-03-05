package router

import (
	"fmt"
	"net/http"

	"github.com/erlendromo/forsete-atr/src/api/handler"
	"github.com/erlendromo/forsete-atr/src/util"
)

type Router interface {
	Serve() error
}

// TODO Add endpoints
func WithEndpoints(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc(fmt.Sprintf("GET %s", util.TIPNOTE_ENDPOINT), handler.GetTipnote)
	mux.HandleFunc(fmt.Sprintf("GET %s", util.BASIC_ENDPOINT), handler.GetBasic)

	return mux
}
