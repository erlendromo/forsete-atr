package httpsrouter

import (
	"testing"
	"time"

	"github.com/erlendromo/forsete-atr/src/api/router"
)

type serveHTTPSCase struct {
	router       router.Router
	expectedPass bool
}

var serveHTTPSCases []serveHTTPSCase = []serveHTTPSCase{
	{router: NewHTTPSRouter("8001", "aaa", "bbb"), expectedPass: false},
	{router: NewHTTPSRouter("8002", "ccc", "ddd"), expectedPass: false},
}

func TestServeTLS(t *testing.T) {
	t.Run("Serve HTTPS test", testServeTLS)
}

func testServeTLS(t *testing.T) {
	for _, httpsCase := range serveHTTPSCases {
		done := make(chan struct{})

		go func() {
			err := httpsCase.router.Serve()
			if (err == nil) != httpsCase.expectedPass {
				t.Errorf("router.Serve() failed: %v", err)
			}
			close(done)
		}()

		select {
		case <-time.After(1 * time.Second):
		case <-done:
		}
	}
}
