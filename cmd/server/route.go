package server

import (
	"fmt"
	"net/http"

	"github.com/ashupednekar/testkage/internal/conf"
)

func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey")
}

func BuildRouter(s *Server) error {
	locations, err := conf.ReadConf("fixtures/sample.yaml")
	if err != nil {
		return err
	}
	for _, loc := range locations {
		fmt.Printf("%v", loc)
		s.Router.HandleFunc(loc.Location, s.Handle)
	}
	return nil
}
