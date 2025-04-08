package httprouter

import (
	"testing"
	"time"

	"github.com/erlendromo/forsete-atr/src/api/router"
)

type serveHTTPCase struct {
	router       router.Router
	expectedPass bool
}

var serveHTTPCases []serveHTTPCase = []serveHTTPCase{
	{router: NewHTTPRouter("test"), expectedPass: true},
	{router: NewHTTPRouter("9090"), expectedPass: true},
}

func TestServe(t *testing.T) {
	t.Run("Serve HTTP test", testServe)
}

func testServe(t *testing.T) {
	for _, httpCase := range serveHTTPCases {
		done := make(chan struct{})

		go func() {
			err := httpCase.router.Serve()
			if (err == nil) != httpCase.expectedPass {
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
