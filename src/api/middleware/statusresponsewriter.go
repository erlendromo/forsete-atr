package middleware

import "net/http"

type StatusResponseWriter struct {
	StatusCode     int
	responseWriter http.ResponseWriter
}

func NewStatusResponseWriter(w http.ResponseWriter) *StatusResponseWriter {
	return &StatusResponseWriter{
		StatusCode:     0,
		responseWriter: w,
	}
}

func (w *StatusResponseWriter) Write(b []byte) (int, error) {
	return w.responseWriter.Write(b)
}

func (w *StatusResponseWriter) Header() http.Header {
	return w.responseWriter.Header()
}

func (w *StatusResponseWriter) WriteHeader(statusCode int) {
	w.StatusCode = statusCode
	w.responseWriter.WriteHeader(statusCode)
}

func (w *StatusResponseWriter) Done() {
	if w.StatusCode == 0 {
		w.StatusCode = http.StatusOK
	}
}
