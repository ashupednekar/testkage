package router

import (
	"fmt"
	"testing"

	"github.com/ashupednekar/testkage/cmd/server"
	"github.com/stretchr/testify/assert"
)

func TestReadConf(t *testing.T) {
	_, err := ReadConf("fixtures/sample.yaml")
	if err != nil {
		t.Fatalf("error reading file")
	}
}

func TestReadConfLength(t *testing.T) {
	conf, _ := ReadConf("fixtures/sample.yaml")
	fmt.Println(conf)
	assert.Equal(t, len(conf), 3)
}

func TestReadConfContent(t *testing.T) {
	conf, _ := ReadConf("fixtures/sample.yaml")
	loc := conf[0]
	assert.Equal(t, loc.Location, "/one")
	assert.Equal(t, loc.Port1, 3001)
	assert.Equal(t, loc.Port2, 3002)
}

func TestBuildRouter(t *testing.T) {
	s := server.NewServer(":3000") // we won't be starting server here, so port doesn't matter
	BuildRouter(s)
}
