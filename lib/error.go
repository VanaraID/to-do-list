package lib

import (
	"encoding/json"
	"net/http"
)

type ErrorResponse struct {
	Code    int    `json:"code"`
	Message string `json:"message"`
}

func ResponseError(w http.ResponseWriter, e ErrorResponse) {
	data := struct {
		Err ErrorResponse `json:"error"`
	}{e}
	res, err := json.Marshal(data)
	if err != nil {
		http.Error(w, "Internal Server Error", 500)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(e.Code)
	w.Write(res)
}

func HTTPError(w http.ResponseWriter, code int, message string) {

	e := ErrorResponse{
		Code:    code,
		Message: message,
	}

	ResponseError(w, e)
}
