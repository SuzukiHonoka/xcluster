package delete

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	"xcluster/internal/api/user"
	"xcluster/internal/server"
	"xcluster/pkg/utils"
)

func ServeUserDelete(w http.ResponseWriter, r *http.Request) {
	// post method only
	if !filter.MatchMethod(w, r, http.MethodDelete) {
		return
	}
	// admin only
	u, ok := user.ParseUserFromSession(w, r)
	if !ok {
		return
	}
	if !u.IsAdmin() {
		w.WriteHeader(http.StatusForbidden)
		return
	}
	// parse be deleted user
	vars := mux.Vars(r)
	sid := vars["id"]
	if filter.FieldsEmpty(w, sid) {
		return
	}
	// parse user
	target, ok := user.ParseUserFromID(w, sid)
	if !ok {
		return
	}
	// check if current user match the target user
	if target.ID == u.ID {
		err := utils.AppendErrorInfo(api.ErrPayloadInvalid, "can't delete current user")
		api.WriteError(w, http.StatusBadRequest, err)
		return
	}
	// delete userID related
	serverGroups, err := server.GetGroups(target.ID)
	if err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "get server group of user failed", err)
		return
	}
	// delete associated server
	if err = serverGroups.Delete(); err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "delete servers of group failed", err)
		return
	}
	// delete associated session
	if !user.DeleteUserSessions(w, target.ID) {
		return
	}
	// actual delete
	if err = target.Delete(); err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "delete user failed", err)
		return
	}
	api.Write(w, api.NewResponse(http.StatusOK, "delete user success", target))
	msg := fmt.Sprintf("user: id=%s deleted", sid)
	logger.Log(msg)
}
