package user

import (
	"net/http"
	"strconv"
	"xcluster/internal/api"
	"xcluster/internal/user"
)

func ParseUserFromID(w http.ResponseWriter, sid string) (*user.User, bool) {
	var err error
	var id int
	if id, err = strconv.Atoi(sid); err != nil {
		err = api.Write(w, api.Response{
			Code:    http.StatusBadRequest,
			Message: "parse user id failed",
		})
		logger.LogIfError(err)
		return nil, false
	}
	// check if id exist
	var u *user.User
	if u, err = user.ID(id).GetUser(); err != nil {
		err = api.Write(w, api.Response{
			Code:    http.StatusBadRequest,
			Message: "user not found",
		})
		logger.LogIfError(err)
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
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusBadRequest,
			Message: "user not found from session",
		})
		logger.LogIfError(err)
		return nil, false
	}
	return u, true
}
