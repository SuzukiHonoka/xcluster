package api

import (
	"fmt"
	"net/http"
	"strings"
)

func ServeNotAllowed(w http.ResponseWriter, r *http.Request) {
	if strings.HasPrefix(r.URL.Path, "/api") {
		w.Header().Set("Content-Type", "application/json")
		Write(w, Response{
			Code:    http.StatusMethodNotAllowed,
			Message: "method not allowed",
		})
		return
	}
	w.WriteHeader(http.StatusMethodNotAllowed)
	_, err := fmt.Fprintln(w, "method not allowed")
	logger.LogIfError(err)
}
