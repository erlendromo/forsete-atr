package httpsrouter

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/erlendromo/forsete-atr/src/api/middleware"
	"github.com/erlendromo/forsete-atr/src/api/router"
)

type HTTPSRouter struct {
	addr     string
	certFile string
	keyFile  string
}

func NewHTTPSRouter(addr, certFile, keyFile string) *HTTPSRouter {
	if intAddr, err := strconv.Atoi(addr); err != nil || addr == "" || intAddr < 1000 || intAddr > 9999 {
		addr = "8080"
	}

	return &HTTPSRouter{
		addr:     addr,
		certFile: certFile,
		keyFile:  keyFile,
	}
}

func (r *HTTPSRouter) Serve() error {
	mux := router.WithV2Endpoints(http.NewServeMux())

	log.Printf("Starting tls-server on port %s...\n", r.addr)

	return http.ListenAndServeTLS(
		fmt.Sprintf(":%s", r.addr),
		r.certFile,
		r.keyFile,
		middleware.WithLogger(
			middleware.WithContexter(
				mux,
			),
		),
	)
}
