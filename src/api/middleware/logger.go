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

	recieveLog := util.NewRecieveLog(
		start.Format(util.TIME_FORMAT),
		zone,
		r.RemoteAddr,
		r.URL.Path,
		r.Method,
		r.Form,
	)

	recieveLog.PrintLog("INFO")

	srw := NewStatusResponseWriter(w)
	l.next.ServeHTTP(srw, r)
	srw.Done()

	took := time.Since(start).Milliseconds()
	unit := "ms"

	if took >= 1000 {
		took = int64(time.Since(start).Seconds())
		unit = "s"
	}

	responseLog := util.NewResponseLog(
		srw.StatusCode,
		took,
		unit,
	)

	var logType string
	if srw.StatusCode < 200 {
		logType = "MISC"
	} else if srw.StatusCode < 300 {
		logType = "SUCCESS"
	} else if srw.StatusCode < 400 {
		logType = "MISC"
	} else if srw.StatusCode < 500 {
		logType = "CLIENT ERROR"
	} else {
		logType = "SERVER ERROR"
	}

	responseLog.PrintLog(logType)
}

func WithLogger(next http.Handler) http.Handler {
	return &Logger{
		next: next,
	}
}
