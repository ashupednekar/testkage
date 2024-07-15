package router

import (
	"github.com/ashupednekar/testkage/cmd/server"
)

func BuildRouter(s *server.Server) error {
	conf, err := ReadConf("fixtures/sample.yaml")
	if err != nil {
		return err
	}
	for _, location := range conf {
	}
	return nil
}
