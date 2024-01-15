package filter

import (
	"net/http"
	"xcluster/internal/api"
	"xcluster/pkg/utils"
)

func FieldsEmpty(w http.ResponseWriter, values ...string) bool {
	if utils.EmptyAny(values) {
		api.WriteError(w, http.StatusBadRequest, api.ErrPayloadNotFound)
		return true
	}
	return false
}
