package list

import (
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	user2 "xcluster/internal/api/user"
	"xcluster/internal/user"
)

func ServeUserList(w http.ResponseWriter, r *http.Request) {
	// get method only
	if !filter.MatchMethod(w, r, http.MethodGet) {
		return
	}
	// admin only
	if !user2.IsAdminFromSession(w, r) {
		return
	}
	users, err := user.All()
	if err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "get users failed", err)
		return
	}
	api.Write(w, api.NewResponse(http.StatusOK, "get users success", users))
}
