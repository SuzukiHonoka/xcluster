package update

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	userApi "xcluster/internal/api/user"
)

func ServeUserUpdateOther(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodPut) {
		return
	}
	u, ok := userApi.ParseUserFromSession(w, r)
	if !ok {
		return
	}
	// block request if not an admin
	if !userApi.IsAdmin(w, u) {
		return
	}
	vars := mux.Vars(r)
	sid, ok := vars["id"]
	if !ok {
		api.WriteError(w, http.StatusBadRequest, api.ErrUserNotFound)
		return
	}
	data, ok := parseData(w, r)
	if !ok {
		return
	}
	u, ok = userApi.ParseUserFromID(w, sid)
	if !ok {
		return
	}
	if !update(w, u, data.Name, data.Password, data.Email) {
		return
	}
	api.Write(w, api.NewResponse(http.StatusOK, "update user success", u))
	msg := fmt.Sprintf("user: id=%s updated", sid)
	logger.Log(msg)
}
