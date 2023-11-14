package handler

import (
	"net/http"
	"strings"
	"xcluster/internal/api"
)

const PartNotFound = api.Logger("handler/notFound")

func NotFound(w http.ResponseWriter, r *http.Request) {
	if !strings.HasPrefix(r.URL.Path, "/api") {
		http.NotFound(w, r)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	err := api.Write(w, api.Response{
		Code:    http.StatusNotFound,
		Message: "api not found",
	})
	PartNotFound.LogIfError(err)
}
