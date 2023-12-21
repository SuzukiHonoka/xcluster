package add

import (
	"fmt"
	"net"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	server2 "xcluster/internal/api/server"
	"xcluster/internal/server"
	"xcluster/pkg/random"
)

func ServeServerAdd(w http.ResponseWriter, r *http.Request) {
	if !filter.MatchMethod(w, r, http.MethodPost) {
		return
	}
	if !filter.ParseForm(w, r) {
		return
	}
	fToken := r.FormValue("token")
	fName := r.FormValue("name")
	fAddr := r.FormValue("addr")
	fPort := r.FormValue("port")
	//
	if filter.FieldsEmpty(w, fToken) {
		return
	}
	//
	// get token and verify
	token, err := server.TokenRaw(fToken).GetToken()
	if err != nil {
		msg := "token not exist"
		api.WriteErrorLog(w, http.StatusInternalServerError, msg, err)
		return
	}
	var ok bool
	if ok, err = token.Validate(); !ok {
		msg := "token invalid"
		api.WriteErrorLog(w, http.StatusBadRequest, msg, err)
		return
	}
	// use remote IP as host addr and combine the port if addr field empty
	if fAddr == "" {
		// port field cannot be empty
		if fPort == "" {
			// reject
			err = api.Write(w, api.Response{
				Code:    http.StatusBadRequest,
				Message: "missing server addr",
			})
			logger.LogIfError(err)
			return
		}
		host := api.RealIP(r)
		fAddr = net.JoinHostPort(host, fPort)
	}
	// if name field empty, use random name
	if fName == "" {
		suffix := random.String(4)
		fName = fmt.Sprintf("server_group-%d_%s", token.GroupID, suffix)
	}
	//
	if _, err = token.UseCapacity(1); err != nil {
		msg := "use token capacity failed"
		api.WriteErrorLog(w, http.StatusInternalServerError, msg, err)
		return
	}
	// actual add
	secretRaw, secret := server.NewRandomSecret()
	name := server.Name(fName)
	addr := server.Addr(fAddr)
	bound, err := server.NewServer(name, addr, secret, token.GroupID)
	if err != nil {
		msg := "add server failed"
		api.WriteErrorLog(w, http.StatusInternalServerError, msg, err)
		return
	}
	// return backend URL, read IP,Port from config
	boundData := server2.NewBound(bound.ServerID, secretRaw, "http://127.0.0.1:443/api/v1")
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "add server success",
		Data:    boundData,
	})
	logger.LogIfError(err)
}
