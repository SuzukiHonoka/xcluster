package filter

import (
	"net/http"
	"xcluster/internal/api"
)

func ParseForm(w http.ResponseWriter, r *http.Request) bool {
	if err := r.ParseForm(); err != nil {
		err = api.Write(w, api.Response{
			Code:    http.StatusBadRequest,
			Message: err.Error(),
		})
		logger.LogError(err)
		return false
	}
	return true
}
