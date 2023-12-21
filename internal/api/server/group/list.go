package group

import (
	"fmt"
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
		err = fmt.Errorf("get server groups failed, cause=%w", err)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusInternalServerError,
			Message: "get server groups failed",
		})
		logger.LogIfError(err)
		return
	}
	// return groups info
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "get server groups success",
		Data:    gs,
	})
	logger.LogIfError(err)
	return
}
