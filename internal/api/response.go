package api

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type Response struct {
	Code    int         `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"` // note: use array as always
}

func NewResponse(code int, message string, data interface{}) Response {
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
		msg := fmt.Sprintf("response marshal failed, err=%s", err)
		message := Response{
			Code:    http.StatusInternalServerError,
			Message: msg,
		}
		//log.Println("api:", msg)
		return message.Byte()
	}
	return b
}
