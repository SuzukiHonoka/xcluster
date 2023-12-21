package filter

import (
	"net/http"
	"xcluster/internal/api"
	"xcluster/pkg/utils"
)

func FieldsEmpty(w http.ResponseWriter, values ...string) bool {
	if utils.EmptyAny(values) {
		msg := "fields can not be empty"
		api.WriteErrorLog(w, http.StatusBadRequest, msg, nil)
		return true
	}
	return false
}
