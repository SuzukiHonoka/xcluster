package api

import (
	"net/http"
	"xcluster/internal/session"
)

func GetSession(r *http.Request) (*session.Session, error) {
	cookie, err := r.Cookie("session")
	if err == nil {
		err = cookie.Valid()
	}
	if err != nil {
		return nil, err
	}
	val := cookie.Value
	if val == "" {
		return nil, ErrSessionNotFound
	}
	id := session.ID(val)
	userSession, err := id.GetSession()
	if err != nil {
		return nil, err
	}
	// generally session will be deleted after the ttl in redis, we assume auto-delete works
	return userSession, nil
}
