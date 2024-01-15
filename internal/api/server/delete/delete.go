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
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "server not found", err)
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
			api.WriteErrorAndLog(w, http.StatusInternalServerError, "server group not found", err)
			return
		}
		// if not match
		if serverGroup.User.ID != u.ID {
			api.WriteErrorAndLog(w, http.StatusForbidden, "server group user mismatch", err)
			return
		}
	}
	err = s.Delete()
	if err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "delete server failed", err)
		return
	}
	api.Write(w, api.NewResponse(http.StatusOK, "delete server success", nil))
}
