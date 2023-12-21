package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// Write writes response to the http.ResponseWriter
// note that it will write the status code by use Response.Code since it uses standard http code for now
func Write(w http.ResponseWriter, r Response) error {
	// log any writes
	//log.Println(r)
	w.WriteHeader(r.Code)
	return json.NewEncoder(w).Encode(r)
}

// WriteErrorLog writes the shallow-error and log it
func WriteErrorLog(w http.ResponseWriter, code int, msg string, err error) {
	err = fmt.Errorf("%s, cause=%w", msg, err)
	logger.LogError(err)
	err = Write(w, Response{
		Code:    code,
		Message: msg,
	})
	logger.LogIfError(err)
}

// WriteError writes the error
func WriteError(w http.ResponseWriter, code int, err error) {
	err = Write(w, Response{
		Code:    code,
		Message: err.Error(),
	})
	logger.LogIfError(err)
}
