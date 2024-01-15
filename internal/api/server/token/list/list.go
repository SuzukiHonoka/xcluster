package list

import (
	"github.com/gorilla/mux"
	"net/http"
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
	id, ok := api.ParseUint(w, fGid, 32)
	if !ok {
		return
	}
	gid := server.GroupID(id)
	tokens, err := gid.GetTokens()
	if err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "get tokens failed", err)
		return
	}
	api.Write(w, api.NewResponse(http.StatusOK, "get tokens success", tokens))
}
