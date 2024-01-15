package middleware

import (
	"net/http"
	"xcluster/internal/api"
	"xcluster/pkg/utils"
)

func SessionMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// parse session
		session, err := api.GetSession(r)
		if err != nil {
			w.WriteHeader(http.StatusUnauthorized)
			err = utils.ExtendError(api.ErrSessionNotFound, err)
			api.WriteError(w, http.StatusUnauthorized, err)
			return
		}
		if session.Expired() {
			w.WriteHeader(http.StatusUnauthorized)
			api.WriteError(w, http.StatusUnauthorized, api.ErrSessionExpired)
			return
		}
		//msg := fmt.Sprintln(session.ShortString(), "being used")
		//logger.Log(msg)
		next.ServeHTTP(w, r)
	})
}
