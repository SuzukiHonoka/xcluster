package api

import (
	"net/http"
	"strings"
)

func ServeNotFound(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api") {
		w.Header().Set("Content-Type", "application/json")
		err := Write(w, Response{
			Code:    http.StatusNotFound,
			Message: "api not found",
		})
		logger.LogIfError(err)
		return
	}
	http.NotFound(w, r)
}
