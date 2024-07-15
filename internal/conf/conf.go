package conf

import (
	"fmt"
	"io/ioutil"

	"gopkg.in/yaml.v2"
)

type Location struct {
	Location string `yaml:location`
	Port1    int    `yaml:port1`
	Port2    int    `yaml:port2`
}

type Conf map[string]Location

func ReadConf(path string) (map[string]Location, error) {
	data, err := ioutil.ReadFile(path)
	if err != nil {
		fmt.Printf("error reading conf file: %v", err)
		return nil, err
	}
	var config Conf
	err = yaml.Unmarshal(data, &config)
	if err != nil {
		fmt.Printf("error unmarhalling conf: %v", err)
		return nil, err
	}
	return config, nil
}
