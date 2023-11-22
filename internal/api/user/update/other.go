package update

import (
	"fmt"
	"github.com/gorilla/mux"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
)

func Other(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodPut) {
		return
	}
	if !filter.ParseForm(w, r) {
		return
	}
	var err error
	vars := mux.Vars(r)
	sid, ok := vars["id"]
	if !ok {
		err = api.Write(w, api.Response{
			Code:    http.StatusBadRequest,
			Message: "id not found",
		})
		logger.LogIfError(err)
		return
	}
	name := r.FormValue("name")
	password := r.FormValue("password")
	email := r.FormValue("email")
	if !update(w, r, sid, name, password, email) {
		return
	}
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "update user success",
	})
	logger.LogIfError(err)
	msg := fmt.Sprintf("user: id=%s updated", sid)
	logger.Log(msg)
}
