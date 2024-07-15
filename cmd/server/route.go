package server

import (
	"fmt"
	"log"
	"net/http"
)

func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
	path := r.URL.Path
	fmt.Printf("%s", path)
	ForwardRequests(r, s.Conf[path].Port1)
	ForwardRequests(r, s.Conf[path].Port2)
}

func (s *Server) BuildRoutes() error {
	for path, loc := range s.Conf {
		log.Printf("%v", loc)
		s.Router.HandleFunc(path, s.Handle)
	}
	return nil
}
