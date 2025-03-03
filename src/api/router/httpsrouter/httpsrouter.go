package httpsrouter

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/erlendromo/forsete-atr/src/api/middleware"
	"github.com/erlendromo/forsete-atr/src/api/router"
	"github.com/erlendromo/forsete-atr/src/util"
)

type HTTPSRouter struct {
	addr     string
	certFile string
	keyFile  string
}

func NewHTTPSRouter(addr, certFile, keyFile string) router.Router {
	if _, err := strconv.Atoi(addr); err != nil {
		addr = util.DEFAULT_API_PORT
	}

	return &HTTPSRouter{
		addr:     addr,
		certFile: certFile,
		keyFile:  keyFile,
	}
}

func (r *HTTPSRouter) Serve() error {
	mux := router.WithEndpoints(http.NewServeMux())
	log.Printf("Starting tls-server on port %s...\n", r.addr)

	return http.ListenAndServeTLS(
		fmt.Sprintf(":%s", r.addr),
		r.certFile,
		r.keyFile,
		middleware.WithLogger(
			mux,
		),
	)
}
