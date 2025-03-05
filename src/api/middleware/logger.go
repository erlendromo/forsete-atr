package middleware

import (
	"net/http"
	"time"

	"github.com/erlendromo/forsete-atr/src/util"
)

type Logger struct {
	next http.Handler
}

func (l *Logger) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	start := time.Now()
	zone, _ := start.Zone()

	srw := NewStatusResponseWriter(w)
	l.next.ServeHTTP(srw, r)
	srw.Done()

	log := util.NewLogData(
		start.Format(util.TIME_FORMAT),
		zone,
		r.URL.Path,
		r.Method,
		r.ContentLength > 0 && (r.Method == http.MethodPost || r.Method == http.MethodPut || r.Method == http.MethodPatch),
		srw.StatusCode,
		time.Since(start).Milliseconds(),
	)

	log.PrintLog()
}

func WithLogger(next http.Handler) http.Handler {
	return &Logger{
		next: next,
	}
}
