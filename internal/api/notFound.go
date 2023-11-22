package api

import (
	"net/http"
	"strings"
)

func NotFound(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/api") {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := Write(w, Response{
		Code:    http.StatusNotFound,
		Message: "api not found",
	})
	logger.LogIfError(err)
}
