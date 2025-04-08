package middleware

import "net/http"

type StatusResponseWriter struct {
	written        bool
	status         int
	responseWriter http.ResponseWriter
}

func NewStatusResponseWriter(w http.ResponseWriter) *StatusResponseWriter {
	return &StatusResponseWriter{
		written:        false,
		status:         0,
		responseWriter: w,
	}
}

func (w *StatusResponseWriter) Write(b []byte) (int, error) {
	if !w.written {
		w.WriteHeader(http.StatusOK)
	}
	return w.responseWriter.Write(b)
}

func (w *StatusResponseWriter) Header() http.Header {
	return w.responseWriter.Header()
}

func (w *StatusResponseWriter) WriteHeader(statusCode int) {
	if !w.written {
		w.written = true
		w.responseWriter.WriteHeader(statusCode)
	}
	w.status = statusCode
}

func (w *StatusResponseWriter) Status() int {
	return w.status
}

func (w *StatusResponseWriter) Written() bool {
	return w.written
}
