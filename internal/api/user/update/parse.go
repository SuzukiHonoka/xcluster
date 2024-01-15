package update

import (
	"net/http"
	"xcluster/internal/api"
)

type Data struct {
	Name     string
	Email    string
	Password string
}

func parseData(w http.ResponseWriter, r *http.Request) (Data, bool) {
	var data Data
	payload, ok := api.ParseJsonPayload(w, r)
	if !ok {
		return data, false
	}
	var ok2, ok3 bool
	data.Name, ok = payload["name"].(string)
	data.Email, ok2 = payload["email"].(string)
	data.Password, ok3 = payload["password"].(string)
	ok = ok || ok2 || ok3
	if !ok {
		api.WriteError(w, http.StatusBadRequest, api.ErrPayloadInvalid)
	}
	return data, ok
}
