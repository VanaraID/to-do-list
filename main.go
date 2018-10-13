package main

import (
	"github.com/VanaraID/to-do-list/response"
	"github.com/go-chi/chi"
)

// NetworthAPI nw api struct
type ServerVanara struct {
	router   *chi.Mux
	response *response.ResponseFormat
}

func main() {

	server := &ServerVanara{
		router:   chi.NewRouter(),
		response: response.NewResponse(),
	}

	server.Start("localhost:6969")
}
