package api

import (
	"errors"
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
		return nil, errors.New("session field empty")
	}
	id := session.ID(val)
	var ss *session.Session
	if ss, err = id.GetSession(); err != nil {
		return nil, err
	}
	// generally session will be deleted after the ttl in redis, we assume auto-delete works
	return ss, nil
}
