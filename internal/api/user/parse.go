package user

import (
	"net/http"
	"strconv"
	"xcluster/internal/api"
	"xcluster/internal/user"
	"xcluster/pkg/utils"
)

func ParseUserFromID(w http.ResponseWriter, sid string) (*user.User, bool) {
	id, err := strconv.Atoi(sid)
	if err != nil {
		err = utils.AppendErrorInfo(api.ErrUserInfoConflict, "parse user id failed")
		api.WriteError(w, http.StatusBadRequest, err)
		return nil, false
	}
	// check if id exist
	u, err := user.ID(id).GetUser()
	if err != nil {
		api.WriteError(w, http.StatusBadRequest, api.ErrUserNotFound)
		return nil, false
	}
	return u, true
}

func ParseUserFromSession(w http.ResponseWriter, r *http.Request) (*user.User, bool) {
	// parse session
	session, _ := api.GetSession(r)
	// get logged-in user
	u, err := user.ID(session.Lease.TenantID).GetUser()
	if err != nil {
		api.WriteErrorAndLog(w, http.StatusBadRequest, api.ErrSessionInvalid.Error(), err)
		return nil, false
	}
	return u, true
}
