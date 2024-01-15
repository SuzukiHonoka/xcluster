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
	if err := session.Delete(); err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "logout failed", err)
		return
	}
	// set cookie
	http.SetCookie(w, &http.Cookie{
		Name:   "session",
		Value:  "",
		Path:   "/",
		MaxAge: -1, // invalidate now
	})
	api.Write(w, api.NewResponse(http.StatusOK, "logout success", nil))
	msg := fmt.Sprintln(session.ShortString(), "deleted")
	logger.Log(msg)
}
