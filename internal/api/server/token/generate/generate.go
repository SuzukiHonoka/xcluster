package generate

import (
	"net/http"
	"time"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	"xcluster/internal/api/user"
	"xcluster/internal/server"
)

func ServeServerTokenGenerate(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodPost) {
		return
	}
	if !filter.ParseForm(w, r) {
		return
	}
	var err error
	fGroupID := r.FormValue("groupID")
	fCapacity := r.FormValue("capacity")
	fDuration := r.FormValue("duration") // in minutes
	if filter.FieldsEmpty(w, fGroupID, fCapacity, fDuration) {
		return
	}
	pGroupID, ok := api.ParseUint(w, fGroupID, 32)
	if !ok {
		return
	}
	pDuration, ok := api.ParseUint(w, fDuration, 32)
	if !ok {
		return
	}
	pCapacity, ok := api.ParseUint(w, fCapacity, 16)
	if !ok {
		return
	}
	// check group existence
	groupID := server.GroupID(pGroupID)
	// cross-group check, bypass admin
	u, ok := user.ParseUserFromSession(w, r)
	if !ok {
		return
	}
	if u.IsAdmin() {
		if _, err = groupID.GetGroup(); err != nil {
			err = api.Write(w, api.Response{
				Code:    http.StatusBadRequest,
				Message: "group id not exist",
			})
			logger.LogIfError(err)
			return
		}
	} else {
		ok, err = server.HasGroupID(u.ID, groupID)
		if err != nil || !ok {
			err = api.Write(w, api.Response{
				Code:    http.StatusBadRequest,
				Message: "group id not allowed or invalid",
			})
			logger.LogIfError(err)
			return
		}
	}
	tokenCapacity := server.TokenCapacity(pCapacity)
	duration := time.Duration(pDuration) * time.Minute
	var token *server.Token
	if token, err = server.NewToken(groupID, tokenCapacity, duration); err != nil {
		msg := "generate token failed"
		api.WriteErrorLog(w, http.StatusInternalServerError, msg, err)
		return
	}
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "generate token success",
		Data:    token,
	})
	logger.LogIfError(err)
}
