package update

import (
	"fmt"
	"net/http"
	"xcluster/internal/api"
	userApi "xcluster/internal/api/user"
	"xcluster/internal/user"
)

func update(w http.ResponseWriter, u *user.User, name, password, email string) bool {
	// check conflict
	if userApi.InfoConflict(w, name, email) {
		return false
	}
	// check if change password
	if password != "" {
		// remove all session related to user
		if !userApi.DeleteUserSessions(w, u.ID) {
			return false
		}
		cause := "password changed"
		msg := fmt.Sprintf("user: id=%d sessions deleted，cause=%s", u.ID, cause)
		logger.Log(msg)
	}
	if err := u.Update(name, password, email); err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "update user info failed", err)
		return false
	}
	return true
}
