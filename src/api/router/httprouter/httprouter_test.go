package httprouter

import (
	"testing"
	"time"

	"github.com/erlendromo/forsete-atr/src/api/router"
)

type serveHTTPCase struct {
	router   router.Router
	expected bool
}

var serveHTTPCases []serveHTTPCase = []serveHTTPCase{
	{
		router:   NewHTTPRouter("test"),
		expected: true,
	},
	{
		router:   NewHTTPRouter("9090"),
		expected: true,
	},
}

func TestServe(t *testing.T) {
	for _, httpCase := range serveHTTPCases {
		done := make(chan struct{})

		go func() {
			if err := httpCase.router.Serve(); err != nil && httpCase.expected == true {
				t.Errorf("router.Serve() failed: %v", err)
			}
			close(done)
		}()

		select {
		case <-time.After(1 * time.Second):
			t.Log("Timeout reached, stopping server")
		case <-done:
			t.Log("Server stopped gracefully")
		}
	}
}
