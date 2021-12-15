package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
)

// in ./.geheim/config.yaml or, alternatively,  ~/.geheim/config.yaml
// read necessary configs for the binary to work
// remote repo url
// secret key

// encripts contents in a file named ./secrets.geheim.yaml by default
// other files can be specified
// keys are left as is
// values are encripted

var LOCAL_LOCATION string
var GLOBAL_LOCATION string

func Get() Config {
	return readConfig()
}

// TODO: error code
// TODO: log exit reason
func init() {
	home, err := os.UserHomeDir()
	if err != nil {
		os.Exit(2)
	}
	LOCAL_LOCATION = ".geheim/config.yaml"
	GLOBAL_LOCATION = fmt.Sprintf("%s/.geheim/config.yaml", home)
}

func readConfig() Config {
	configLocation := getConfigLocation()
	return readYaml(configLocation)
}

// TODO: error code for exiting
// TODO: how to exit gracefully?
// TODO: log exit reason
func getConfigLocation() string {
	_, err := os.Stat(LOCAL_LOCATION)
	if !errors.Is(err, os.ErrNotExist) {
		return LOCAL_LOCATION
	}
	_, err = os.Stat(GLOBAL_LOCATION)
	if !errors.Is(err, os.ErrNotExist) {
		return GLOBAL_LOCATION
	}
	os.Exit(1)
	return ""
}

// TODO: exit error
// TODO: log error
func readYaml(configLocation string) Config {
	data, err := ioutil.ReadFile(configLocation)
	if err != nil {
		os.Exit(3)
	}
	var config Config
	err = config.Parse(data)
	if err != nil {
		os.Exit(4)
	}
	return config
}
