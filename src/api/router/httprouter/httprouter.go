package httprouter

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/erlendromo/forsete-atr/src/api/middleware"
	"github.com/erlendromo/forsete-atr/src/api/router"
	"github.com/erlendromo/forsete-atr/src/util"
)

type HTTPRouter struct {
	addr string
}

func NewHTTPRouter(addr string) *HTTPRouter {
	if intAddr, err := strconv.Atoi(addr); err != nil || addr == "" || intAddr < 1000 || intAddr > 9999 {
		addr = util.DEFAULT_API_PORT
	}

	return &HTTPRouter{
		addr: addr,
	}
}

func (r *HTTPRouter) Serve() error {
	mux := router.WithV2Endpoints(http.NewServeMux())
	log.Printf("Starting server on port %s...\n", r.addr)

	return http.ListenAndServe(
		fmt.Sprintf(":%s", r.addr),
		middleware.WithLogger(
			middleware.WithContexter(
				mux,
			),
		),
	)
}
