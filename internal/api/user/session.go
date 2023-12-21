package user

import (
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/user"
)

func DeleteUserSessionsFromSession(w http.ResponseWriter, r *http.Request) bool {
	u, ok := ParseUserFromSession(w, r)
	if !ok {
		return false
	}
	return DeleteUserSessions(w, u.ID)
}

func DeleteUserSessions(w http.ResponseWriter, uid user.ID) bool {
	var err error
	var userSessions user.Sessions
	if userSessions, err = uid.GetSessions(); err != nil {
		msg := "get user sessions failed"
		api.WriteErrorLog(w, http.StatusInternalServerError, msg, err)
		return false
	}
	if err = userSessions.Invalidate(); err != nil {
		msg := "invalidate user sessions failed"
		api.WriteErrorLog(w, http.StatusInternalServerError, msg, err)
		return false
	}
	if err = uid.DeleteSessions(); err != nil {
		msg := "delete user sessions failed"
		api.WriteErrorLog(w, http.StatusInternalServerError, msg, err)
		return false
	}
	return true
}
