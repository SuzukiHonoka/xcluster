package middleware

import (
	"fmt"
	"net/http"
	"xcluster/internal/api"
)

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// parse session
		s, err := api.GetSession(r)
		if err != nil {
			err = api.Write(w, api.Response{
				Code:    http.StatusForbidden,
				Message: "session not found",
			})
			logger.LogIfError(err)
			return
		}
		msg := fmt.Sprintln(s.ShortString(), "being used")
		logger.Log(msg)
		next.ServeHTTP(w, r)
	})
}
