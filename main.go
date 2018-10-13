package main

import (
	"encoding/json"
	"fmt"
	"net/http"

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
	message := "todos!"
	w.Write([]byte(message))
}

func OutputJSON(w http.ResponseWriter, o interface{}) {
	res, err := json.Marshal(o)
	if err != nil {
		w.Write([]byte(err.Error()))
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.Write(res)
	w.Write([]byte("\n"))
}
