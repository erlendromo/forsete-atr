package httpsrouter

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	appcontext "github.com/erlendromo/forsete-atr/src/api/app_context"
	"github.com/erlendromo/forsete-atr/src/api/middleware"
	"github.com/erlendromo/forsete-atr/src/api/router"
	"github.com/erlendromo/forsete-atr/src/util"
)

type HTTPSRouter struct {
	addr     string
	certFile string
	keyFile  string
}

func NewHTTPSRouter(addr, certFile, keyFile string) *HTTPSRouter {
	if intAddr, err := strconv.Atoi(addr); err != nil || addr == "" || intAddr < 1000 || intAddr > 9999 {
		addr = util.DEFAULT_API_PORT
	}

	return &HTTPSRouter{
		addr:     addr,
		certFile: certFile,
		keyFile:  keyFile,
	}
}

func (r *HTTPSRouter) Serve() error {
	mux := router.WithV2Endpoints(http.NewServeMux(), appcontext.GetAppContext())

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
