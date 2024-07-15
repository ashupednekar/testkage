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
	Conf   map[string]conf.Location
}

func NewServer(addr string) (*Server, error) {
	locations, err := conf.ReadConf("internal/conf/fixtures/sample.yaml")
	log.Println("Building routes...")
	if err != nil {
		return &Server{}, err
	}
	return &Server{
		Addr:   addr,
		Router: mux.NewRouter(),
		Conf:   locations,
	}, nil
}

func (s *Server) Start() error {
	s.BuildRoutes()
	log.Printf("Listening at %s", s.Addr)
	err := http.ListenAndServe(s.Addr, s.Router)
	if err != nil {
		log.Fatalf("could not start server")
	}
	log.Printf("Listening at :%s", s.Addr)
	return nil
}
