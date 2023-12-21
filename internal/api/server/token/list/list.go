package list

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	"xcluster/internal/server"
)

func ServeServerGroupTokenList(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodGet) {
		return
	}
	fGid := mux.Vars(r)["gid"]
	if filter.FieldsEmpty(w, fGid) {
		return
	}
	var err error
	//
	var id uint64
	if id, err = strconv.ParseUint(fGid, 10, 32); err != nil {
		api.WriteError(w, http.StatusBadRequest, err)
		return
	}
	gid := server.GroupID(id)
	var tokens server.Tokens
	if tokens, err = gid.GetTokens(); err != nil {
		msg := "get tokens failed"
		api.WriteErrorLog(w, http.StatusInternalServerError, msg, err)
		return
	}
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "get tokens success",
		Data:    tokens,
	})
	logger.LogIfError(err)
}
