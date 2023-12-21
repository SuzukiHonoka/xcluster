package update

import (
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
)

func ServeUserUpdateSelf(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodPost) {
		return
	}
	if !filter.ParseForm(w, r) {
		return
	}
	name := r.FormValue("name")
	password := r.FormValue("password")
	email := r.FormValue("email")
	if !update(w, r, "", name, password, email) {
		return
	}
	err := api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "update user success",
	})
	logger.LogIfError(err)
}
