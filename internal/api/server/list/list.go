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
	var err error
	// get all groups from user
	var groups server.Groups
	if groups, err = server.GetGroups(u.ID); err != nil {
		msg := "get server groups from user failed"
		api.WriteErrorLog(w, http.StatusInternalServerError, msg, err)
		return
	}
	// get servers in each group
	var servers server.Servers
	if servers, err = groups.GetServers(); err != nil {
		msg := "get servers from server groups failed"
		api.WriteErrorLog(w, http.StatusInternalServerError, msg, err)
		return
	}
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "get servers success",
		Data:    servers,
	})
	logger.LogIfError(err)
}
