package main

import (
	"log"

	"github.com/ashupednekar/testkage/cmd/server"
)

func main() {
	s, err := server.NewServer(":3000")
	if err != nil {
		log.Fatal(err)
	}
	s.Start()
}
