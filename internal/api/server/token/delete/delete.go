package delete

import (
	"fmt"
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
		err = api.Write(w, api.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		logger.LogIfError(err)
		return
	}
	if err = server.TokenID(tid).DeleteToken(); err != nil {
		err = fmt.Errorf("delete token failed, cause=%w", err)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusInternalServerError,
			Message: "delete token failed",
		})
		logger.LogIfError(err)
		return
	}
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "delete token success",
	})
	logger.LogIfError(err)
}
