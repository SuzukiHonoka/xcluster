package filter

import (
	"net/http"
	"xcluster/internal/api"
)

func ParseForm(w http.ResponseWriter, r *http.Request) bool {
	if err := r.ParseForm(); err != nil {
		api.WriteError(w, http.StatusBadRequest, err)
		return false
	}
	return true
}
