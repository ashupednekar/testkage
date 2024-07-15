package server

import (
	"fmt"
	"net/http"
)

func (s *Server) Hey(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "hey")
}
