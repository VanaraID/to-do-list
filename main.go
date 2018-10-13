package main

import (
	"fmt"
	"net/http"
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
	message := "welcome to vanara"
	w.Write([]byte(message))
}

func handlerTodo(w http.ResponseWriter, r *http.Request) {
	message := "todos!"
	w.Write([]byte(message))
}
