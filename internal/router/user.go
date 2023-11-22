package router

import (
	"github.com/gorilla/mux"
	"xcluster/internal/api/user/delete"
	"xcluster/internal/api/user/list"
	"xcluster/internal/api/user/login"
	"xcluster/internal/api/user/logout"
	"xcluster/internal/api/user/signup"
	"xcluster/internal/api/user/update"
)

func applyUserRoutes(authenticated, unauthenticated *mux.Router) {
	unauthenticatedUser := unauthenticated.PathPrefix("/user").Subrouter()
	unauthenticatedUser.HandleFunc("/signup", signup.Signup)
	unauthenticatedUser.HandleFunc("/login", login.Login)
	//
	authenticatedUser := authenticated.PathPrefix("/user").Subrouter()
	authenticatedUser.HandleFunc("/logout", logout.Logout)
	authenticatedUser.HandleFunc("/delete/{id:[0-9]+}", delete.Delete)
	authenticatedUser.HandleFunc("/list", list.List)
	authenticatedUser.HandleFunc("/update/{id:[0-9]+}", update.Other)
	authenticatedUser.HandleFunc("/update", update.Self)
	authenticatedUser.HandleFunc("/info", nil)
}
