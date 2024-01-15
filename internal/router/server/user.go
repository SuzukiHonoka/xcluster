package server

import (
	"github.com/gorilla/mux"
	userDelete "xcluster/internal/api/user/delete"
	userInfo "xcluster/internal/api/user/info"
	userList "xcluster/internal/api/user/list"
	userLogin "xcluster/internal/api/user/login"
	userLogout "xcluster/internal/api/user/logout"
	userRegister "xcluster/internal/api/user/register"
	userUpdate "xcluster/internal/api/user/update"
)

func ApplyUserRoutes(authenticated, unauthenticated *mux.Router) {
	unauthenticatedUser := unauthenticated.PathPrefix("/user").Subrouter()
	unauthenticatedUser.HandleFunc("/register", userRegister.ServeUserRegister)
	unauthenticatedUser.HandleFunc("/login", userLogin.ServeUserLogin)
	//
	authenticatedUser := authenticated.PathPrefix("/user").Subrouter()
	authenticatedUser.HandleFunc("/logout", userLogout.ServeUserLogout)
	authenticatedUser.HandleFunc("/delete/{id:[0-9]+}", userDelete.ServeUserDelete)
	authenticatedUser.HandleFunc("/list", userList.ServeUserList)
	authenticatedUser.HandleFunc("/update/{id:[0-9]+}", userUpdate.ServeUserUpdateOther)
	authenticatedUser.HandleFunc("/update", userUpdate.ServeUserUpdateSelf)
	authenticatedUser.HandleFunc("/info", userInfo.ServeUserInfo)
}
