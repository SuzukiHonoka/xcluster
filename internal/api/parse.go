package api

import (
	"encoding/json"
	"net/http"
	"strconv"
	"xcluster/pkg/utils"
)

func ParseInt(w http.ResponseWriter, val string, bitSize int) (int64, bool) {
	i, err := strconv.ParseInt(val, 10, bitSize)
	if err != nil {
		err = utils.ExtendError(ErrPayloadInvalid, err)
		WriteErrorAndLog(w, http.StatusBadRequest, "parse int error", err)
		return 0, false
	}
	return i, true
}

func ParseUint(w http.ResponseWriter, val string, bitSize int) (uint64, bool) {
	i, err := strconv.ParseUint(val, 10, bitSize)
	if err != nil {
		err = utils.ExtendError(ErrPayloadInvalid, err)
		WriteErrorAndLog(w, http.StatusBadRequest, "parse uint error", err)
		return 0, false
	}
	return i, true
}

func ParseJsonPayload(w http.ResponseWriter, r *http.Request) (map[string]interface{}, bool) {
	if r.Header.Get("Content-Type") != "application/json" {
		WriteError(w, http.StatusBadRequest, ErrPayloadNotFound)
		return nil, false
	}
	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		err = utils.ExtendError(ErrPayloadInvalid, err)
		WriteErrorAndLog(w, http.StatusBadRequest, "decode data failed", err)
		return nil, false
	}
	return payload, true
}
