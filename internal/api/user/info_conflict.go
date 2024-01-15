package user

import (
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/user"
	"xcluster/pkg/utils"
)

func InfoConflict(w http.ResponseWriter, name, email string) bool {
	if name != "" {
		if _, err := user.Name(name).GetUser(); err == nil {
			err = utils.AppendErrorInfo(api.ErrUserInfoConflict, "user name not available")
			api.WriteError(w, http.StatusBadRequest, err)
			return true
		}
	}
	if email != "" {
		if _, err := user.Email(email).GetUser(); err == nil {
			err = utils.AppendErrorInfo(api.ErrUserInfoConflict, "user email not available")
			api.WriteError(w, http.StatusBadRequest, err)
			return true
		}
	}
	return false
}
