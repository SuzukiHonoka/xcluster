package group

import (
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	"xcluster/internal/api/user"
	"xcluster/internal/server"
)

func ServeList(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodGet) {
		return
	}
	// list groups own by userid
	u, ok := user.ParseUserFromSession(w, r)
	if !ok {
		return
	}
	gs, err := server.GetGroups(u.ID)
	if err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "get server groups from user failed", err)
		return
	}
	// return groups info
	api.Write(w, api.NewResponse(http.StatusOK, "get server groups from user success", gs))
}
