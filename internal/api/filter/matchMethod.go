package filter

import (
	"net/http"
	"xcluster/internal/api"
)

func MatchMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		err := api.Write(w, api.Response{
			Code:    http.StatusMethodNotAllowed,
			Message: "method not allowed",
		})
		logger.LogIfError(err)
		return false
	}
	return true
}
