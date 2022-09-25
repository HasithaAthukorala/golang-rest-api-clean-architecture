package server

import (
	"github.com/gorilla/mux"
	"net/http"
)

// TODO: add authentication handler and location based handler middlewares here

func ResponseHeadersMiddleware(headers map[string]string) mux.MiddlewareFunc {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
			for k, v := range headers {
				w.Header().Set(k, v)
			}
			next.ServeHTTP(w, req)
		})
	}
}
