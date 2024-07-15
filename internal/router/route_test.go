package router

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestReadConf(t *testing.T) {
	conf := ReadConf("fixtures/sample.yaml")
	fmt.Println(conf)
	assert.Equal(t, len(conf), 3)
	loc := conf[0]
	assert.Equal(t, loc.Location, "/one")
	assert.Equal(t, loc.Port1, 3001)
	assert.Equal(t, loc.Port2, 3002)
}
