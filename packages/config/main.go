package config

import (
	"errors"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"treuzedev/geheim/packages/logging"
	"treuzedev/geheim/packages/shared"
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
	logging.Log(logging.Info, fmt.Sprintf("Local config: %v", LOCAL_LOCATION))
	logging.Log(logging.Info, fmt.Sprintf("Global config: %v", GLOBAL_LOCATION))
}

func parseCliFlags() CliFlags {
	secretKey := parseStringFlag("secretkey", "k", "", "A key to encrypt/decrypt files. If not specified, the program will try to get one from local/global config file.")
	encrypt := parseBoolFlag("encrypt", "e", false, "Whether to encrypt the files defined in the config file. Defaults to 'false'. If both encrypt and decrypt flags are set to 'true', the encrypt flag takes precedence. If both the encrypt flag and the decrypt are set to 'false', the default behavior is to encrypt any unencrypted files, ie, the encrypt flag becomes 'true'.")
	decrypt := parseBoolFlag("decrypt", "d", false, "Whether to decrypt the files defined in the config file. Defaults to 'false'. If both encrypt and decrypt flags are set to 'true', the encrypt flag takes precedence. If both the encrypt flag and the decrypt are set to 'false', the default behavior is to encrypt any unencrypted files, ie, the encrypt flag becomes 'true'.")
	flag.Parse()
	cliFlags := CliFlags{SecretKey: *secretKey, Encrypt: *encrypt, Decrypt: *decrypt}
	logging.Log(logging.Info, fmt.Sprintf("CLI flags: --secretkey=*** --encrypt=%v --decrypt=%v", cliFlags.Encrypt, cliFlags.Decrypt))
	return cliFlags
}

func parseStringFlag(longFlag, shortFlag, defaultValue, usage string) *string {
	longFlagValue := flag.String(longFlag, defaultValue, usage)
	shortFlagValue := flag.String(shortFlag, defaultValue, fmt.Sprintf("See -%v", longFlag))
	if *longFlagValue != defaultValue {
		return longFlagValue
	} else {
		return shortFlagValue
	}
}

func parseBoolFlag(longFlag, shortFlag string, defaultValue bool, usage string) *bool {
	longFlagValue := flag.Bool(longFlag, defaultValue, usage)
	shortFlagValue := flag.Bool(shortFlag, defaultValue, fmt.Sprintf("See -%v", longFlag))
	if *longFlagValue != defaultValue {
		return longFlagValue
	} else {
		return shortFlagValue
	}
}

func readConfig() Config {
	configLocation := getConfigLocation()
	config := readYaml(configLocation)
	logging.Log(logging.Info, fmt.Sprintf("Config file from %v", configLocation))
	logging.Log(logging.Info, fmt.Sprintf("config.yaml: secretkey=***, files=%v", config.Files))
	return config
}

func getConfigLocation() string {
	_, err := os.Stat(LOCAL_LOCATION)
	if !errors.Is(err, os.ErrNotExist) {
		return LOCAL_LOCATION
	}
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
	config.Encrypt = cliFlags.Encrypt
	config.Decrypt = cliFlags.Decrypt
	logging.Log(logging.Info, fmt.Sprintf("Merged CLI and config.yaml: secretkey=***, encrypt=%v, decrypt=%v, files=%v", config.Encrypt, config.Decrypt, config.Files))
	return config
}

func checkConfigAndApplyDefaults(config Config) Config {
	if config.SecretKey == "" {
		shared.CheckError(fmt.Errorf("a secret key must be set"))
	}
	if shared.CompareStringSlices(config.Files, []string{}) {
		config.Files = []string{"secrets.geheim.yaml"}
	}
	if config.Encrypt && config.Decrypt {
		config.Decrypt = false
	} else if !config.Encrypt && !config.Decrypt {
		config.Encrypt = true
	}
	logging.Log(logging.Info, fmt.Sprintf("Final config with needed defaults: secretkey=***, encrypt=%v, decrypt=%v, files=%v", config.Encrypt, config.Decrypt, config.Files))
	return config
}
