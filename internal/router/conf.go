package router

import (
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v2"
)

type Location struct {
	Location string `yaml:location`
	Port1    int    `yaml:port1`
	Port2    int    `yaml:port2`
}

type Conf struct {
	Conf []Location
}

func ReadConf(path string) []Location {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		log.Fatalf("error reading conf file: %v", err)
	}
	var config Conf
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		log.Fatalf("error unmarhalling conf: %v", err)
	}
	return config.Conf
}
