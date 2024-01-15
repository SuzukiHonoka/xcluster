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
	if ok = IsAdmin(w, u); !ok {
		w.WriteHeader(http.StatusForbidden)
	}
	return ok
}

func IsAdmin(w http.ResponseWriter, u *user.User) bool {
	if !u.IsAdmin() {
		err := fmt.Errorf("%w: user=%s (ID %d)", api.ErrUserExceedGroupPermission, u.Name, u.ID)
		api.WriteErrorAndLog(w, http.StatusForbidden, "admin group required", err)
		return false
	}
	return true
}
