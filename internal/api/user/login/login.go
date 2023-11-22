package login

import (
	"fmt"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	options "xcluster/internal/api/user"
	"xcluster/internal/session"
	"xcluster/internal/user"
)

// Login authenticates the requested user, assign session if authentication success
func Login(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodPost) {
		return
	}
	if !filter.ParseForm(w, r) {
		return
	}
	var err error
	name := r.FormValue("name")
	password := r.FormValue("password")
	if filter.FieldsEmpty(w, name, password) {
		return
	}
	var u *user.User
	u, err = user.Name(name).GetUser()
	if err != nil {
		err = api.Write(w, api.Response{
			Code:    http.StatusBadRequest,
			Message: "user not found",
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
	l := session.NewLease(uint(u.ID), options.SessionDuration)
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
	if _, err = user.SaveSession(s); err != nil {
		err = fmt.Errorf("save user session failed, cause=%w", err)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusInternalServerError,
			Message: "save user session failed",
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
		Expires: s.Lease.ExpirationTime,
	})
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "login success",
	})
	logger.LogIfError(err)
	msg := fmt.Sprintf("%s -> user: %s", s.ShortString(), u.Name)
	logger.Log(msg)
}
