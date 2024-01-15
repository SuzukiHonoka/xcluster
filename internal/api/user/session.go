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
	userSessions, err := uid.GetSessions()
	if err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "get user sessions failed", err)
		return false
	}
	if err = userSessions.Invalidate(); err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "invalidate user sessions failed", err)
		return false
	}
	if err = uid.DeleteSessions(); err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "delete user sessions failed", err)
		return false
	}
	return true
}
