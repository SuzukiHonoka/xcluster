package filter

import (
	"net/http"
	"xcluster/internal/api"
	"xcluster/internal/utils"
)

func FieldsEmpty(w http.ResponseWriter, values ...string) bool {
	for _, value := range values {
		if utils.Empty(value) {
			err := api.Write(w, api.Response{
				Code:    http.StatusBadRequest,
				Message: "fields can not be empty",
			})
			logger.LogIfError(err)
			return true
		}
	}
	return false
}
