package delete

import (
	"github.com/gorilla/mux"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	"xcluster/internal/api/user"
	"xcluster/internal/server"
)

func ServeServerDelete(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodDelete) {
		return
	}
	fSid := mux.Vars(r)["sid"]
	if filter.FieldsEmpty(w, fSid) {
		return
	}
	s, err := server.ID(fSid).GetServer()
	if err != nil {
		msg := "server not found"
		api.WriteErrorLog(w, http.StatusInternalServerError, msg, err)
		return
	}
	// check if admin
	u, ok := user.ParseUserFromSession(w, r)
	if !ok {
		return
	}
	// check ownership if not an admin
	if !u.IsAdmin() {
		// get server group
		serverGroup, err := s.GroupID.GetGroup()
		if err != nil {
			msg := "server group not found"
			api.WriteErrorLog(w, http.StatusInternalServerError, msg, err)
			return
		}
		// if not match
		if !(serverGroup.User.ID == u.ID) {
			msg := "server group does not belongs to you"
			api.WriteErrorLog(w, http.StatusForbidden, msg, err)
			return
		}
	}
	err = s.Delete()
	if err != nil {
		msg := "delete server failed"
		api.WriteErrorLog(w, http.StatusInternalServerError, msg, err)
		return
	}
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "delete server success",
	})
	logger.LogIfError(err)
}
