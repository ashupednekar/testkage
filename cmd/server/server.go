package server

import (
	"log"
	"net/http"

	"github.com/ashupednekar/testkage/internal/conf"
	"github.com/gorilla/mux"
)

type Server struct {
	Addr   string
	Router *mux.Router
	Conf   []conf.Location
}

func NewServer(addr string) *Server {
	return &Server{
		Addr:   addr,
		Router: mux.NewRouter(),
	}
}

func (s *Server) Start() error {
	log.Printf("Listening at %s", s.Addr)
	s.BuildRoutes()
	err := http.ListenAndServe(s.Addr, s.Router)
	if err != nil {
		log.Fatalf("could not start server")
	}
	log.Printf("Listening at :%s", s.Addr)
	return nil
}
