package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

	"github.com/VanaraID/to-do-list/model"
)

type Error struct {
	HTTPCode int    `json:"-"`
	Code     int    `json:"code,omitempty"`
	Massage  string `json:"massage"`
}

func JSONError(e Error) []byte {
	data := struct {
		Err Error `json:"error"`
	}{e}
	b, err := json.Marshal(data)
	if err != nil {
		return []byte(err.Error())
	}

	return b
}

func DisplayError(w http.ResponseWriter) {
	e := Error{
		HTTPCode: http.StatusForbidden,
		Code:     403,
		Massage:  "An Error Occured",
	}
	w.Header().Set("Content-Type", "application/json")
	w.Write(JSONError(e))

}

func main() {
	http.HandleFunc("/", handlerIndex)
	http.HandleFunc("/todo", handlerTodo)

	var address = "localhost:6969"
	fmt.Printf("server started at %s\n", address)
	err := http.ListenAndServe(address, nil)
	if err != nil {
		fmt.Println(err.Error())
	}
}

func handlerIndex(w http.ResponseWriter, r *http.Request) {
	OutputJSON(w, model.GetTodos())
}

func handlerTodo(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {

		idInt, err := strconv.Atoi(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		OutputJSON(w, model.SelectTodo(idInt))
		return
	}
	DisplayError(w)
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
}
