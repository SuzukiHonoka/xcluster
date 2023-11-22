package user

import (
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	"xcluster/internal/user"
)

func InfoConflict(w http.ResponseWriter, name, email string) bool {
	var err error
	if name == "" && email == "" {
		return false
	}
	if name != "" {
		if filter.FieldsEmpty(w, name) {
			return true
		}
		if _, err = user.Name(name).GetUser(); err == nil {
			err = api.Write(w, api.Response{
				Code:    http.StatusBadRequest,
				Message: "user name not available",
			})
			logger.LogIfError(err)
			return true
		}
	}
	if email != "" {
		if filter.FieldsEmpty(w, email) {
			return true
		}
		if _, err = user.Email(email).GetUser(); err == nil {
			err = api.Write(w, api.Response{
				Code:    http.StatusBadRequest,
				Message: "user email not available",
			})
			logger.LogIfError(err)
			return true
		}
	}
	return false
}
