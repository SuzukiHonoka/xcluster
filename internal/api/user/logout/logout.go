package logout

import (
	"fmt"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
)

func ServeUserLogout(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodGet) {
		return
	}
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
	msg := fmt.Sprintln(session.ShortString(), "deleted")
	logger.Log(msg)
}
