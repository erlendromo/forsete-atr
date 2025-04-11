package middleware

import (
	"context"
	"errors"
	"net/http"

	"github.com/erlendromo/forsete-atr/src/config"
	"github.com/erlendromo/forsete-atr/src/util"
)

type Contexter struct {
	next http.Handler
}

func (c *Contexter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Set timeout
	timeoutCtx, cancelTimeout := context.WithTimeout(r.Context(), config.GetConfig().TIMEOUT)
	defer cancelTimeout()

	// Reassign request and responsewriter
	r = r.WithContext(timeoutCtx)
	srw := NewStatusResponseWriter(w)

	// Serve request
	done := make(chan struct{})
	go func() {
		defer close(done)
		c.next.ServeHTTP(srw, r)
	}()

	// Defer timeout, cancel or success
	select {
	case <-timeoutCtx.Done():
		if !srw.Written() {
			util.ERROR(srw, http.StatusRequestTimeout, errors.New("timeout reached, try again"))
		}
	case <-r.Context().Done():

		// THIS IS ONLY WRITING THE HEADER, NOT THE BODY FOR SOME REASON

		if !srw.Written() {
			util.ERROR(srw, 499, errors.New("request cancelled by client"))
		}
	case <-done:
		// The request completed successfully
	}
}

func WithContexter(next http.Handler) http.Handler {
	return &Contexter{
		next: next,
	}
}
