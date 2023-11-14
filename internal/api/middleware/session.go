package middleware

import (
	"net/http"
	"xcluster/internal/api"
)

const loggerSession = api.Logger("middleware-session")

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// parse session
		if _, err := api.GetSession(r); err != nil {
			loggerSession.LogError(err)
			err = api.Write(w, api.Response{
				Code:    http.StatusForbidden,
				Message: "session not found",
			})
			loggerSession.LogIfError(err)
			return
		}
		next.ServeHTTP(w, r)
	})
}
