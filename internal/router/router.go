package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"xcluster/internal/api/handler"
	"xcluster/internal/api/middleware"
	"xcluster/internal/api/user/login"
	"xcluster/internal/api/user/logout"
	"xcluster/internal/api/user/signup"
)

var headers = map[string]string{
	"Content-Type": "application/json",
}

func NewRouter() *mux.Router {
	// api root
	api := mux.NewRouter().PathPrefix("/api").Subrouter()
	api.NotFoundHandler = http.HandlerFunc(handler.NotFound)
	api.Use(middleware.HeaderMiddleware(headers))
	api.Use(middleware.LogMiddleware)
	// api v1
	v1 := api.PathPrefix("/v1").Subrouter()
	// unauthenticated
	unauthenticated := v1.NewRoute().Subrouter()
	unauthenticatedUser := unauthenticated.PathPrefix("/user").Subrouter()
	unauthenticatedUser.HandleFunc("/signup", signup.Signup)
	unauthenticatedUser.HandleFunc("/login", login.Login)
	// authenticated
	authenticated := v1.NewRoute().Subrouter()
	authenticated.Use(middleware.SessionMiddleware)
	authenticatedUser := authenticated.PathPrefix("/user").Subrouter()
	authenticatedUser.HandleFunc("/logout", logout.Logout)
	authenticatedUser.HandleFunc("/info", nil)
	return api
}
