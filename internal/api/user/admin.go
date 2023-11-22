package user

import (
	"fmt"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/user"
)

func IsAdminFromSession(w http.ResponseWriter, r *http.Request) bool {
	u, ok := ParseUserFromSession(w, r)
	if !ok {
		return false
	}
	return IsAdmin(w, u)
}

func IsAdmin(w http.ResponseWriter, u *user.User) bool {
	if !u.Admin() {
		err := fmt.Errorf("user: %s (ID %d), trying to excced the user group permission", u.Name, u.ID)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusForbidden,
			Message: "admin group required",
		})
		logger.LogIfError(err)
		return false
	}
	return true
}
