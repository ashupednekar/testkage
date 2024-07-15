package main

import "github.com/ashupednekar/testkage/cmd/server"

func main() {
	s := server.NewServer(":3000")
	s.Start()
}
