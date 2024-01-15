package group

import (
	"fmt"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	"xcluster/internal/api/user"
	"xcluster/internal/server"
)

func ServeServerCreate(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodPost) {
		return
	}
	if !filter.ParseForm(w, r) {
		return
	}
	// parse group name and user id
	name := r.FormValue("name")
	if filter.FieldsEmpty(w, name) {
		return
	}
	u, ok := user.ParseUserFromSession(w, r)
	if !ok {
		return
	}
	g, err := server.NewGroup(u.ID, name)
	if err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "create server group failed", err)
		return
	}
	// return group info
	api.Write(w, api.NewResponse(http.StatusOK, "create server group success", g))
	msg := fmt.Sprintf("group: id=%d, name=%s, userID=%d created", g.ID, g.Name, g.UserID)
	logger.Log(msg)
}
