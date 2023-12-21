package server

import (
	"github.com/gorilla/mux"
	"net/http"
	serverAdd "xcluster/internal/api/server/add"
	serverControl "xcluster/internal/api/server/control"
	serverDelete "xcluster/internal/api/server/delete"
	serverGroup "xcluster/internal/api/server/group"
	serverList "xcluster/internal/api/server/list"
	serverTokenDelete "xcluster/internal/api/server/token/delete"
	serverTokenGenerate "xcluster/internal/api/server/token/generate"
	serverTokenList "xcluster/internal/api/server/token/list"
)

func ApplyServerRoutes(authenticated, unauthenticated *mux.Router) {
	// authenticated server
	authenticatedServer := authenticated.PathPrefix("/server").Subrouter()
	authenticatedServer.HandleFunc("/add", serverAdd.ServeServerAdd).Methods(http.MethodPost)
	authenticatedServer.HandleFunc("/list", serverList.ServeServerList)
	authenticatedServer.HandleFunc("/delete/{sid}", serverDelete.ServeServerDelete)
	// group
	groupRouter := authenticatedServer.PathPrefix("/group").Subrouter()
	groupRouter.HandleFunc("/create", serverGroup.ServeServerCreate)
	groupRouter.HandleFunc("/list", serverGroup.ServeList)
	// todo: [admin] get specified user groups
	groupRouter.HandleFunc("/list/{id:[0-9]+}", nil)
	// token
	tokenRouter := authenticatedServer.PathPrefix("/token").Subrouter()
	tokenRouter.HandleFunc("/generate", serverTokenGenerate.ServeServerTokenGenerate)
	tokenRouter.HandleFunc("/list/{gid:[0-9]+}", serverTokenList.ServeServerGroupTokenList)
	tokenRouter.HandleFunc("/delete/{tid:[0-9]+}", serverTokenDelete.ServeServerTokenDelete)
	// server websocket
	unauthenticatedServer := unauthenticated.PathPrefix("/server").Subrouter()
	unauthenticatedServer.HandleFunc("/control", serverControl.ServeServerControl)
}
