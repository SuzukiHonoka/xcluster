package middleware

import (
	"log"
	"net/http"
	"xcluster/internal/api"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addr := api.RealIP(r)
		method := r.Method
		uri := r.RequestURI
		log.Printf("%s -> %s %s", addr, method, uri)
		next.ServeHTTP(w, r)
	})
}
