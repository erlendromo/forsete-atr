package router

import (
	"encoding/json"
	"net/http"
)

type Router interface {
	Serve() error
}

// TODO Add endpoints
func WithEndpoints(mux *http.ServeMux) *http.ServeMux {
	mux.HandleFunc("GET /forsete-atr/v1/test", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode(map[string]string{
			"Hello": "World!",
		})
	})
	mux.HandleFunc("POST /forsete-atr/v1/test", func(w http.ResponseWriter, r *http.Request) {
		var body string
		json.NewDecoder(r.Body).Decode(&body)
		json.NewEncoder(w).Encode(map[string]string{
			"Body": body,
		})
	})

	return mux
}
