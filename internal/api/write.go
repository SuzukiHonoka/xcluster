package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Write writes response to the http.ResponseWriter, log write error if any
func Write(w http.ResponseWriter, r Response) {
	// log any writes
	//log.Println(r)
	//w.WriteHeader(r.Code)
	err := json.NewEncoder(w).Encode(r)
	logger.LogIfError(err)
}

// WriteErrorAndLog writes the shallow-error and log it
func WriteErrorAndLog(w http.ResponseWriter, code int, msg string, err error) {
	err = fmt.Errorf("%s, cause=%w", msg, err)
	// hide the actual error since it might contain sensitive data
	logger.LogError(err)
	Write(w, Response{
		Code:    code,
		Message: msg,
	})
}

// WriteError writes the error
func WriteError(w http.ResponseWriter, code int, err error) {
	Write(w, Response{
		Code:    code,
		Message: err.Error(),
	})
}
