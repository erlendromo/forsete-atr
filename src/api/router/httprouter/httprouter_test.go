package httprouter

import (
	"testing"
	"time"

	appcontext "github.com/erlendromo/forsete-atr/src/api/app_context"
	"github.com/erlendromo/forsete-atr/src/api/router"
	"github.com/erlendromo/forsete-atr/src/database/mock"
)

type serveHTTPCase struct {
	router       router.Router
	expectedPass bool
}

var serveHTTPCases []serveHTTPCase = []serveHTTPCase{
	{router: NewHTTPRouter("test"), expectedPass: true},
	{router: NewHTTPRouter("9090"), expectedPass: true},
}

func setup() {
	appcontext.InitAppContext(
		mock.NewMockDatabase(),
	)
}

func TestServe(t *testing.T) {
	setup()
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
