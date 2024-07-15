package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	addr   string
	Router *mux.Router
}

func NewServer(addr string) *Server {
	return &Server{
		addr:   addr,
		Router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	err := http.ListenAndServe(s.addr, s.Router)
	if err != nil {
		log.Fatalf("could not start server")
	}
	fmt.Printf("Listening at :%s", s.addr)
	return nil
}
