package list

import (
	"fmt"
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/api/filter"
	user2 "xcluster/internal/api/user"
	"xcluster/internal/user"
)

func List(w http.ResponseWriter, r *http.Request) {
	// get method only
	if !filter.MatchMethod(w, r, http.MethodGet) {
		return
	}
	// admin only
	if !user2.IsAdminFromSession(w, r) {
		return
	}
	all, err := user.All()
	if err != nil {
		err = fmt.Errorf("get users failed, cause=%w", err)
		logger.LogError(err)
		err = api.Write(w, api.Response{
			Code:    http.StatusInternalServerError,
			Message: "get users failed",
		})
		logger.LogIfError(err)
		return
	}
	api.NewResponse(http.StatusOK, "get users success", all)
	err = api.Write(w, api.Response{
		Code:    http.StatusOK,
		Message: "get users success",
		Data:    all,
	})
	logger.LogIfError(err)
}
