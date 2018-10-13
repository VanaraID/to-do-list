package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/middleware"
)

func (s *ServerVanara) initRoutes() {
	s.router.Get("/", s.handlerIndex())
	s.router.Get("/todo", s.handlerTodo())
}

func (s *ServerVanara) initMiddleware() {
	s.router.Use(middleware.RequestID)
	s.router.Use(middleware.RealIP)
	s.router.Use(middleware.Logger)
	s.router.Use(middleware.Recoverer)
	s.router.Use(middleware.Timeout(60 * time.Second))
}

// Start start api service
func (s *ServerVanara) Start(host string) {
	s.initMiddleware()
	s.initRoutes()

	fmt.Printf("server started at %s\n", host)
	err := http.ListenAndServe(host, s.router)
	if err != nil {
		fmt.Println(err.Error())
	}
}
