package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/middleware"
)

var headers = map[string]string{
	"Content-Type": "application/json",
}

func NewRouter() *mux.Router {
	// api root
	root := mux.NewRouter().PathPrefix("/api").Subrouter()
	root.NotFoundHandler = http.HandlerFunc(api.NotFound)
	root.Use(middleware.HeaderMiddleware(headers))
	root.Use(middleware.LogMiddleware)
	// api v1
	v1 := root.PathPrefix("/v1").Subrouter()
	// unauthenticated
	unauthenticated := v1.NewRoute().Subrouter()
	// authenticated
	authenticated := v1.NewRoute().Subrouter()
	authenticated.Use(middleware.SessionMiddleware)
	// apply
	applyUserRoutes(authenticated, unauthenticated)
	return root
}
