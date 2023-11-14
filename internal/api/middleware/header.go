package middleware

import "net/http"

func HeaderMiddleware(headers map[string]string) func(next http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			for header, value := range headers {
				w.Header().Set(header, value)
				next.ServeHTTP(w, r)
			}
		})
	}
}
