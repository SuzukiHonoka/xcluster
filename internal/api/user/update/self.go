package update

import (
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	userApi "xcluster/internal/api/user"
)

func ServeUserUpdateSelf(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodPost) {
		return
	}
	u, ok := userApi.ParseUserFromSession(w, r)
	if !ok {
		return
	}
	data, ok := parseData(w, r)
	if !ok {
		return
	}
	if !update(w, u, data.Name, data.Password, data.Email) {
		return
	}
	api.Write(w, api.NewResponse(http.StatusOK, "update current user info success", nil))
}
