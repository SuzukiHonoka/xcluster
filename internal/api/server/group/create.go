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
	var err error
	// parse group name and user id
	name := r.FormValue("name")
	if name == "" {
		err = api.Write(w, api.Response{
			Code:    http.StatusBadRequest,
			Message: "server group name cannot empty",
		})
		logger.LogIfError(err)
		return
	}
	u, ok := user.ParseUserFromSession(w, r)
	if !ok {
		return
	}
	var g *server.Group
	if g, err = server.NewGroup(u.ID, name); err != nil {
		err = fmt.Errorf("create server group failed, cause=%w", err)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusInternalServerError,
			Message: "create server group failed",
		})
		logger.LogIfError(err)
		return
	}
	// return group info
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "create server group success",
		Data:    g,
	})
	logger.LogIfError(err)
	msg := fmt.Sprintf("group: id=%d, name=%s, userID=%d created", g.ID, g.Name, g.UserID)
	logger.Log(msg)
	return
}
