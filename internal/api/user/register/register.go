package register

import (
	"fmt"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	userApi "xcluster/internal/api/user"
	"xcluster/internal/server"
	"xcluster/internal/user"
)

// ServeUserRegister register a user and store it to database
// only accept POST request
func ServeUserRegister(w http.ResponseWriter, r *http.Request) {
	// check request method
	if !filter.MatchMethod(w, r, http.MethodPost) {
		return
	}
	if !userApi.AllowRegister {
		api.Write(w, api.NewResponse(http.StatusPreconditionFailed, "register is not allowed", nil))
		return
	}
	payload, ok := api.ParseJsonPayload(w, r)
	if !ok {
		return
	}
	name, ok := payload["name"].(string)
	password, ok2 := payload["password"].(string)
	email, ok3 := payload["email"].(string)
	if !ok || !ok2 || !ok3 {
		api.WriteError(w, http.StatusBadRequest, api.ErrPayloadInvalid)
		return
	}
	// check if any field empty
	if filter.FieldsEmpty(w, name, password, email) {
		return
	}
	// check if name or email already in use
	if userApi.InfoConflict(w, name, email) {
		return
	}
	// actually add the user
	u, err := user.NewUser(name, password, email)
	if err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "create user failed", err)
		return
	}
	// add default group
	if _, err = server.NewGroup(u.ID, "default"); err != nil {
		api.WriteErrorAndLog(w, http.StatusInternalServerError, "create default group failed", err)
		return
	}
	api.Write(w, api.NewResponse(http.StatusOK, "register success", nil))
	msg := fmt.Sprintf("user registred: %s", u)
	logger.Log(msg)
}
