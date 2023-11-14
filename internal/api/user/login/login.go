package login

import (
	"fmt"
	"net/http"
	"time"
	"xcluster/internal/api"
	options "xcluster/internal/api/user"
	"xcluster/internal/session"
	"xcluster/internal/user"
	"xcluster/internal/utils"
)

// Login authenticates the requested user, assign session if authentication success
func Login(w http.ResponseWriter, r *http.Request) {
	var err error
	if err = r.ParseForm(); err != nil {
		err = api.Write(w, api.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		logger.LogIfError(err)
		return
	}
	name := r.FormValue("name")
	password := r.FormValue("password")
	if utils.EmptyAny(name, password) {
		err = api.Write(w, api.Response{
			Code:    http.StatusBadRequest,
			Message: "fields can not be empty",
		})
		logger.LogIfError(err)
		return
	}
	var u user.User
	u, err = user.Name(name).GetUser()
	if err != nil {
		err = api.Write(w, api.Response{
			Code:    http.StatusBadRequest,
			Message: "fields can not be empty",
		})
		logger.LogIfError(err)
		return
	}
	var ok bool
	ok, err = u.Password.Compare(password)
	if err != nil {
		err = fmt.Errorf("compare user password failed, cause=%w", err)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusInternalServerError,
			Message: "password compare failed",
		})
		logger.LogIfError(err)
		return
	}
	if !ok {
		err = api.Write(w, api.Response{
			Code:    http.StatusForbidden,
			Message: "wrong password",
		})
		logger.LogIfError(err)
		return
	}
	// allocate session
	l := session.NewLease(u.ID, options.SessionDuration)
	var s *session.Session
	s, err = session.NewSession(l)
	if err != nil {
		err = fmt.Errorf("create user session failed, cause=%w", err)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusInternalServerError,
			Message: "create user session failed",
		})
		logger.LogIfError(err)
		return
	}
	// set cookie
	http.SetCookie(w, &http.Cookie{
		Name:  "session",
		Value: string(s.ID),
		Path:  "/",
		//Domain:     "",
		Expires: s.Lease.ExpireTime,
	})
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "login success",
	})
	logger.LogIfError(err)
	msg := fmt.Sprintf("session: %s (expiration %s) -> user: %s",
		s.ID, s.Lease.ExpireTime.Format(time.DateTime), u.Name)
	logger.Log(msg)
}
