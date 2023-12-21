package api

import (
	"net/http"
	"strconv"
)

func ParseInt(w http.ResponseWriter, val string, bitSize int) (int64, bool) {
	i, err := strconv.ParseInt(val, 10, bitSize)
	if err != nil {
		err = Write(w, Response{
			Code:    http.StatusBadRequest,
			Message: "parse int error: " + err.Error(),
		})
		logger.LogIfError(err)
		return 0, false
	}
	return i, true
}

func ParseUint(w http.ResponseWriter, val string, bitSize int) (uint64, bool) {
	i, err := strconv.ParseUint(val, 10, bitSize)
	if err != nil {
		err = Write(w, Response{
			Code:    http.StatusBadRequest,
			Message: "parse uint error: " + err.Error(),
		})
		logger.LogIfError(err)
		return 0, false
	}
	return i, true
}
