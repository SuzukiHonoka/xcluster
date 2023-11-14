package logout

import (
	"fmt"
	"net/http"
	"xcluster/internal/api"
)

const logger = api.Logger("user/logout")

func Logout(w http.ResponseWriter, r *http.Request) {
	// parse session
	session, _ := api.GetSession(r)
	err := session.Delete()
	if err != nil {
		err = fmt.Errorf("delete session failed, cause=%w", err)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusInternalServerError,
			Message: "logout failed",
		})
		logger.LogIfError(err)
	}
	// set cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // invalidate now
	})
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "logout success",
	})
	logger.LogIfError(err)
	msg := fmt.Sprintf("session: %s (tenantID %d) deleted", session.ID, session.Lease.TenantID)
	logger.Log(msg)
}
