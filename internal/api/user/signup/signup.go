package signup

import (
	"fmt"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/user"
	"xcluster/internal/utils"
)

// Signup a user and store it to database
// only accept POST request
func Signup(w http.ResponseWriter, r *http.Request) {
	var err error
	// check request method
	if r.Method != http.MethodPost {
		err = api.Write(w, api.Response{
			Code:    http.StatusMethodNotAllowed,
			Message: "method not allowed",
		})
		logger.LogIfError(err)
		return
	}
	// parse form data
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
	email := r.FormValue("email")
	// check if any field empty
	if utils.EmptyAny(name, password, email) {
		err = api.Write(w, api.Response{
			Code:    http.StatusBadRequest,
			Message: "fields can not be empty",
		})
		logger.LogIfError(err)
		return
	}
	// add user
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
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "signup success",
	})
	logger.LogIfError(err)
	logger.Log(u.String())
}
