package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"

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
	http.Error(w, "parameter id is a must", http.StatusInternalServerError)
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
