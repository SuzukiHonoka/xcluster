package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

type Response struct {
	Code    int           `json:"code"`
	Message string        `json:"message"`
	Data    []interface{} `json:"data,omitempty"`
}

func NewResponse(code int, message string, data []interface{}) Response {
	return Response{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func (r Response) String() string {
	return string(r.Byte())
}

func (r Response) Byte() []byte {
	b, err := json.Marshal(r)
	if err != nil {
		detail := fmt.Sprintf("response marshal failed, err=%s", err)
		message := Response{
			Code:    http.StatusInternalServerError,
			Message: detail,
		}
		log.Println("api:", detail)
		return message.Byte()
	}
	return b
}

// Write writes response to the http.ResponseWriter
// note that it will write the status code by use Response.Code since it uses standard http code for now
func Write(w http.ResponseWriter, r Response) error {
	// log ant writes
	log.Println(r)
	w.WriteHeader(r.Code)
	return json.NewEncoder(w).Encode(r)
}
