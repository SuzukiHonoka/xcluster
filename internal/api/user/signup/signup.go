package signup

import (
	"fmt"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	userApi "xcluster/internal/api/user"
	"xcluster/internal/server"
	"xcluster/internal/user"
)

// ServeUserSignup signup a user and store it to database
// only accept POST request
func ServeUserSignup(w http.ResponseWriter, r *http.Request) {
	var err error
	// check request method
	if !filter.MatchMethod(w, r, http.MethodPost) {
		return
	}
	// parse form data
	if !filter.ParseForm(w, r) {
		return
	}
	name := r.FormValue("name")
	password := r.FormValue("password")
	email := r.FormValue("email")
	// check if any field empty
	if filter.FieldsEmpty(w, name, password, email) {
		return
	}
	// check if name or email in use
	if userApi.InfoConflict(w, name, email) {
		return
	}
	// actually add user
	var u *user.User
	if u, err = user.NewUser(name, password, email); err != nil {
		// hide the actual error since it might contain sensitive data
		err = fmt.Errorf("create user failed, cause=%w", err)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusInternalServerError,
			Message: "create user failed",
		})
		logger.LogIfError(err)
		return
	}
	// add default group
	if _, err = server.NewGroup(u.ID, "default"); err != nil {
		err = fmt.Errorf("create default group for user failed, cause=%w", err)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusInternalServerError,
			Message: "create default group for user failed",
		})
		return
	}
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "signup success",
	})
	logger.LogIfError(err)
	logger.Log(u.String())
}
