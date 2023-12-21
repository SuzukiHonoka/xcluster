package info

import (
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	"xcluster/internal/api/user"
)

func ServeUserInfo(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodGet) {
		return
	}
	u, ok := user.ParseUserFromSession(w, r)
	if !ok {
		return
	}
	err := api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "get user info success",
		Data:    u,
	})
	logger.LogIfError(err)
}
