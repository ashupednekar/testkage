package server

import (
	"fmt"
	"log"
	"net/http"

	"github.com/ashupednekar/testkage/internal/conf"
)

func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey")
}

func (s *Server) BuildRoutes() error {
	locations, err := conf.ReadConf("internal/conf/fixtures/sample.yaml")
	log.Println("Building routes...")
	if err != nil {
		return err
	}
	for _, loc := range locations {
		log.Printf("%v", loc)
		s.Router.HandleFunc(loc.Location, s.Handle)
	}
	return nil
}
