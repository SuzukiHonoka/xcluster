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
			api.WriteErrorAndLog(w, http.StatusBadRequest, "group id not exist", err)
			return
		}
	} else {
		ok, err = server.HasGroupID(u.ID, groupID)
		if err != nil || !ok {
			//"group id not allowed or invalid"
			api.WriteErrorAndLog(w, http.StatusBadRequest, "group verification failed", err)
			return
		}
	}
	tokenCapacity := server.TokenCapacity(pCapacity)
	duration := time.Duration(pDuration) * time.Minute
	var token *server.Token
	if token, err = server.NewToken(groupID, tokenCapacity, duration); err != nil {
		msg := "generate token failed"
		api.WriteErrorAndLog(w, http.StatusInternalServerError, msg, err)
		return
	}
	api.Write(w, api.NewResponse(http.StatusOK, "generate token success", token))
}
