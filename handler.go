package main

import (
	"net/http"
	"strconv"

	"github.com/VanaraID/to-do-list/model"
)

func (s *ServerVanara) handlerIndex() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		s.response.Payload(w, http.StatusOK, model.GetTodos())
	}

}

func (s *ServerVanara) handlerTodo() http.HandlerFunc {

	return func(w http.ResponseWriter, r *http.Request) {
		if id := r.URL.Query().Get("id"); id != "" {

			idInt, err := strconv.Atoi(id)

			if err != nil {
				s.response.Payload(w, http.StatusInternalServerError, err.Error())
			}

			s.response.Payload(w, http.StatusOK, model.SelectTodo(idInt))
			return
		}

		s.response.Payload(w, http.StatusInternalServerError, "parameter id is a must.")

	}

}
