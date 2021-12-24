package config

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"treuzedev/geheim/shared"
)

var LOCAL_LOCATION string
var GLOBAL_LOCATION string

func Get() Config {
	cliFlags := parseCliFlags()
	config := readConfig()
	mergedConfig := mergeCliAndConfig(cliFlags, config)
	return checkConfigAndApplyDefaults(mergedConfig)
}

func init() {
	home, err := os.UserHomeDir()
	shared.CheckError(err)
	LOCAL_LOCATION = ".geheim/config.yaml"
	GLOBAL_LOCATION = fmt.Sprintf("%s/.geheim/config.yaml", home)
}

func parseCliFlags() CliFlags {
	secretKey := parseFlag("secretkey", "k", "", "A key to encrypt/decrypt files. If not specified, the program will try to get one from local/global config file.")
	flag.Parse()
	return CliFlags{SecretKey: secretKey}
}

func parseFlag(longFlag, shortFlag, defaultValue, usage string) string {
	var longFlagValue string
	flag.StringVar(&longFlagValue, longFlag, defaultValue, usage)
	var shortFlagValue string
	flag.StringVar(&shortFlagValue, shortFlag, defaultValue, fmt.Sprintf("See -%v", longFlag))
	if longFlagValue != defaultValue {
		return longFlagValue
	} else {
		return shortFlagValue
	}
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

func mergeCliAndConfig(cliFlags CliFlags, config Config) Config {
	if cliFlags.SecretKey != "" {
		config.SecretKey = cliFlags.SecretKey
	}
	return config
}

func checkConfigAndApplyDefaults(config Config) Config {
	if config.SecretKey == "" {
		shared.CheckError(fmt.Errorf("a secret key must be set"))
	}
	if shared.CompareStringSlices(config.Files, []string{}) {
		config.Files = []string{"secrets.geheim.yaml"}
	}
	return config
}
