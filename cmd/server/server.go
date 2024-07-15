package server

import (
	"net/http"

	"github.com/gorilla/mux"
)

type Server struct {
	addr   string
	router *mux.Router
}

func NewServer(addr string) Server {
	return &Server{
		addr: addr,
    router: mux.NewRouter()
	}
}

func (s *Server) Start() error {
	return http.ListenAndServe(s.addr, s.router)
}
