package main

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/VanaraID/to-do-list/lib"
	"github.com/VanaraID/to-do-list/model"
)

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
	lib.HTTPJSON(w, model.GetTodos())
}

func handlerTodo(w http.ResponseWriter, r *http.Request) {
	if id := r.URL.Query().Get("id"); id != "" {

		idInt, err := strconv.Atoi(id)

		if err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}

		lib.HTTPJSON(w, model.SelectTodo(idInt))
		return
	}

	lib.HTTPError(w, http.StatusInternalServerError, "parameter id is a must.")
}
