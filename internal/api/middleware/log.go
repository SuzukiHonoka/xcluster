package middleware

import (
	"log"
	"net/http"
	"xcluster/internal/utils"
)

func LogMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		addr := utils.RealIP(r)
		method := r.Method
		uri := r.RequestURI
		log.Printf("%s -> %s %s", addr, method, uri)
		next.ServeHTTP(w, r)
	})
}
