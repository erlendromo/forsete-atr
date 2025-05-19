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
	// Start timer
	start := time.Now()
	zone, _ := start.Zone()

	// Print receive-log
	receiveLog := util.NewReceiveLog(
		start.Format(util.TIME_FORMAT),
		zone,
		r.RemoteAddr,
		r.URL.Path,
		r.Method,
	)
	receiveLog.PrintLog(util.INFO)

	// Do request
	srw := NewStatusResponseWriter(w)
	l.next.ServeHTTP(srw, r)

	// Calculate time elapsed
	elapsed := time.Since(start)
	var took int64
	var unit string
	switch {
	case elapsed >= time.Minute:
		took = int64(elapsed.Minutes())
		unit = "m"
	case elapsed >= time.Second:
		took = int64(elapsed.Seconds())
		unit = "s"
	case elapsed >= time.Millisecond:
		took = elapsed.Milliseconds()
		unit = "ms"
	default:
		took = elapsed.Microseconds()
		unit = "μs"
	}

	// Assign log-type
	var logType string
	if srw.Status() < 200 {
		logType = util.MISC
	} else if srw.Status() < 300 {
		logType = util.SUCCESS
	} else if srw.Status() < 400 {
		logType = util.MISC
	} else if srw.Status() < 500 {
		logType = util.CLIENT_ERROR
	} else {
		logType = util.SERVER_ERROR
	}

	// Print response-log
	responseLog := util.NewResponseLog(
		srw.Status(),
		took,
		unit,
	)
	responseLog.PrintLog(logType)
}

func WithLogger(next http.Handler) http.Handler {
	return &Logger{
		next: next,
	}
}
