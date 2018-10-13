package response

import (
	"encoding/json"
	"net/http"
)

type ResponseFormat struct {
	Code         int         `json:"code"`
	ResponseJSON interface{} `json:"response"`
}

func (r *ResponseFormat) Send(w http.ResponseWriter) {
	res, err := json.Marshal(r)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(r.Code)
	w.Write(res)
}

func (r *ResponseFormat) Payload(w http.ResponseWriter, code int, response interface{}) {

	e := ResponseFormat{
		Code:         code,
		ResponseJSON: response,
	}

	e.Send(w)
}

func NewResponse() *ResponseFormat {
	return &ResponseFormat{}
}
