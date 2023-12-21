package update

import (
	"fmt"
	"net/http"
	"strconv"
	"xcluster/internal/api"
	userApi "xcluster/internal/api/user"
)

func update(w http.ResponseWriter, r *http.Request, sid, name, password, email string) bool {
	var err error
	// grab requesting user
	u, ok := userApi.ParseUserFromSession(w, r)
	if !ok {
		return false
	}
	// check admin
	if !userApi.IsAdmin(w, u) {
		return false
	}
	// check if restful update
	if sid != "" {
		// check if update other
		if strconv.Itoa(int(u.ID)) != sid {
			u, ok = userApi.ParseUserFromID(w, sid)
			if !ok {
				return false
			}
		}
	}
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
	if userApi.InfoConflict(w, name, email) {
		return false
	}
	if err = u.Update(name, password, email); err != nil {
		err = fmt.Errorf("update user failed, cause=%w", err)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusInternalServerError,
			Message: "update user failed",
		})
		logger.LogIfError(err)
		return false
	}
	msg := fmt.Sprintf("user: id=%d updated", u.ID)
	logger.Log(msg)
	return true
}
