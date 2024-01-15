package filter

import (
	"net/http"
	"xcluster/internal/api"
)

func MatchMethod(w http.ResponseWriter, r *http.Request, method string) bool {
	if r.Method != method {
		api.Write(w, api.Response{
			Code:    http.StatusMethodNotAllowed,
			Message: "method not allowed",
		})
		return false
	}
	return true
}
