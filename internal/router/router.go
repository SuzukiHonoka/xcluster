package router

import (
	"github.com/gorilla/mux"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/middleware"
	"xcluster/internal/router/server"
)

var headers = map[string]string{
	"Content-Type": "application/json",
}

func NewServerRouter() *mux.Router {
	// api root
	root := mux.NewRouter().PathPrefix("/api").Subrouter()
	root.NotFoundHandler = http.HandlerFunc(api.ServeNotFound)
	root.MethodNotAllowedHandler = http.HandlerFunc(api.ServeNotAllowed)
	root.Use(middleware.CorsMiddleware("http://localhost:5173"))
	root.Use(middleware.HeaderMiddleware(headers))
	root.Use(middleware.LogMiddleware)
	// only allow https due to security reasons
	root.StrictSlash(true)
	// api v1
	v1 := root.PathPrefix("/v1").Subrouter()
	// unauthenticated
	unauthenticated := v1.NewRoute().Subrouter()
	// authenticated
	authenticated := v1.NewRoute().Subrouter()
	authenticated.Use(middleware.SessionMiddleware)
	// apply routers
	server.ApplyUserRoutes(authenticated, unauthenticated)
	server.ApplyServerRoutes(authenticated, unauthenticated)
	return root
}
