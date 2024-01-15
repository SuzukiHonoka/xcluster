package delete

import (
	"github.com/gorilla/mux"
	"net/http"
	"strconv"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	"xcluster/internal/server"
)

func ServeServerTokenDelete(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodDelete) {
		return
	}
	fTokenID := mux.Vars(r)["tid"]
	if filter.FieldsEmpty(w, fTokenID) {
		return
	}
	tid, err := strconv.ParseUint(fTokenID, 10, 32)
	if err != nil {
		api.WriteError(w, http.StatusBadRequest, err)
		return
	}
	if err = server.TokenID(tid).DeleteToken(); err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "delete token failed", err)
		return
	}
	api.Write(w, api.NewResponse(http.StatusOK, "delete token success", nil))
}
