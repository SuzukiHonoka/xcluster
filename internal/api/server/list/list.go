package list

import (
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	"xcluster/internal/api/user"
	"xcluster/internal/server"
)

func ServeServerList(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodGet) {
		return
	}
	// get user
	u, ok := user.ParseUserFromSession(w, r)
	if !ok {
		return
	}
	// get all groups from user
	groups, err := server.GetGroups(u.ID)
	if err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "get server groups from user failed", err)
		return
	}
	// get servers in each group
	servers, err := groups.GetServers()
	if err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "get servers from server groups failed", err)
		return
	}
	api.Write(w, api.NewResponse(http.StatusOK, "get servers success", servers))
}
