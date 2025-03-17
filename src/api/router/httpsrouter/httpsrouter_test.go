package httpsrouter

import (
	"testing"
	"time"

	"github.com/erlendromo/forsete-atr/src/api/router"
)

type serveHTTPSCase struct {
	router   router.Router
	expected bool
}

var serveHTTPSCases []serveHTTPSCase = []serveHTTPSCase{
	{
		router:   NewHTTPSRouter("8001", "aaa", "bbb"),
		expected: false,
	},
	{
		router:   NewHTTPSRouter("8002", "ccc", "ddd"),
		expected: false,
	},
}

func TestServeTLS(t *testing.T) {
	t.Run("Serve HTTPS test", testServeTLS)
}

func testServeTLS(t *testing.T) {
	for _, httpsCase := range serveHTTPSCases {
		done := make(chan struct{})

		go func() {
			if err := httpsCase.router.Serve(); err != nil && httpsCase.expected == true {
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
