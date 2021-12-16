package config

import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"treuzedev/geheim/shared"
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

func init() {
	home, err := os.UserHomeDir()
	shared.CheckError(err)
	LOCAL_LOCATION = ".geheim/config.yaml"
	GLOBAL_LOCATION = fmt.Sprintf("%s/.geheim/config.yaml", home)
}

func readConfig() Config {
	configLocation := getConfigLocation()
	return readYaml(configLocation)
}

func getConfigLocation() string {
	_, err := os.Stat(LOCAL_LOCATION)
	if !errors.Is(err, os.ErrNotExist) {
		return LOCAL_LOCATION
	}
	fmt.Println(err)
	_, err = os.Stat(GLOBAL_LOCATION)
	if !errors.Is(err, os.ErrNotExist) {
		return GLOBAL_LOCATION
	}
	shared.CheckError(err)
	return ""
}

func readYaml(configLocation string) Config {
	data, err := ioutil.ReadFile(configLocation)
	shared.CheckError(err)
	var config Config
	err = config.Parse(data)
	shared.CheckError(err)
	return config
}
