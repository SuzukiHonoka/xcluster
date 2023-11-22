package user

import (
	"fmt"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/user"
)

func DeleteUserSessionsFromSession(w http.ResponseWriter, r *http.Request) bool {
	u, ok := ParseUserFromSession(w, r)
	if !ok {
		return false
	}
	var err error
	var userSessions user.Sessions
	if userSessions, err = u.ID.GetSessions(); err != nil {
		err = fmt.Errorf("get user sessions failed, cause=%w", err)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusInternalServerError,
			Message: "get user sessions failed",
		})
		logger.LogIfError(err)
		return false
	}
	if err = userSessions.Invalidate(); err != nil {
		err = fmt.Errorf("invalidate user sessions failed, cause=%w", err)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusInternalServerError,
			Message: "invalidate user sessions failed",
		})
		logger.LogIfError(err)
		return false
	}
	if err = u.ID.DeleteSessions(); err != nil {
		err = fmt.Errorf("delete user sessions failed, cause=%w", err)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusInternalServerError,
			Message: "delete user sessions failed",
		})
		logger.LogIfError(err)
		return false
	}
	return true
}
