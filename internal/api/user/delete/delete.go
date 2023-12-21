package delete

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	"xcluster/internal/api/user"
	"xcluster/internal/server"
)

func ServeUserDelete(w http.ResponseWriter, r *http.Request) {
	// post method only
	if !filter.MatchMethod(w, r, http.MethodDelete) {
		return
	}
	// admin only
	if !user.IsAdminFromSession(w, r) {
		return
	}
	// parse be deleted user
	if !filter.ParseForm(w, r) {
		return
	}
	vars := mux.Vars(r)
	sid := vars["id"]
	if filter.FieldsEmpty(w, sid) {
		return
	}
	// reuse error variable
	var err error
	// parse user
	u, ok := user.ParseUserFromID(w, sid)
	if !ok {
		return
	}
	// delete userID related
	var serverGroups server.Groups
	if serverGroups, err = server.GetGroups(u.ID); err != nil {
		msg := "get server group of user failed"
		api.WriteErrorLog(w, http.StatusInternalServerError, msg, err)
		return
	}
	// delete associated server
	if err = serverGroups.Delete(); err != nil {
		msg := "delete servers of group failed"
		api.WriteErrorLog(w, http.StatusInternalServerError, msg, err)
		return
	}
	// delete associated session
	if !user.DeleteUserSessions(w, u.ID) {
		return
	}
	// actual delete
	if err = u.Delete(); err != nil {
		err = fmt.Errorf("delete user failed, cause=%w", err)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusInternalServerError,
			Message: "delete user failed",
		})
		logger.LogIfError(err)
		return
	}
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "delete user success",
	})
	logger.LogIfError(err)
	msg := fmt.Sprintf("user: id=%s deleted", sid)
	logger.Log(msg)
}
