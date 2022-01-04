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

func Get() (config Config) {
	cliFlags := parseCliFlags()
	tmpConfig := readConfig()
	mergedConfig := mergeCliAndConfig(cliFlags, tmpConfig)
	return checkConfigAndApplyDefaults(mergedConfig)
}

func init() {
	home, err := os.UserHomeDir()
	shared.CheckError(err, nil)
	LOCAL_LOCATION = ".geheim/config.yaml"
	GLOBAL_LOCATION = fmt.Sprintf("%s/.geheim/config.yaml", home)
	logging.Log(logging.Info, fmt.Sprintf("Local config: %v", LOCAL_LOCATION))
	logging.Log(logging.Info, fmt.Sprintf("Global config: %v", GLOBAL_LOCATION))
}

func parseCliFlags() (cliFlags CliFlags) {
	secretKey := parseStringFlag("secretkey", "k", "", "A key to encrypt/decrypt files. If not specified, the program will try to get one from local/global config file.")
	encrypt := parseBoolFlag("encrypt", "e", false, "Whether to encrypt the files defined in the config file. Defaults to 'false'. If both encrypt and decrypt flags are set to 'true', the encrypt flag takes precedence. If both the encrypt flag and the decrypt are set to 'false', the default behavior is to encrypt any unencrypted files, ie, the encrypt flag becomes 'true'.")
	decrypt := parseBoolFlag("decrypt", "d", false, "Whether to decrypt the files defined in the config file. Defaults to 'false'. If both encrypt and decrypt flags are set to 'true', the encrypt flag takes precedence. If both the encrypt flag and the decrypt are set to 'false', the default behavior is to encrypt any unencrypted files, ie, the encrypt flag becomes 'true'.")
	flag.Parse()
	cliFlags = CliFlags{SecretKey: *secretKey, Encrypt: *encrypt, Decrypt: *decrypt}
	logging.Log(logging.Info, fmt.Sprintf("CLI flags: --secretkey=*** --encrypt=%v --decrypt=%v", cliFlags.Encrypt, cliFlags.Decrypt))
	return cliFlags
}

func parseStringFlag(longFlag, shortFlag, defaultValue, usage string) (flagValue *string) {
	longFlagValue := flag.String(longFlag, defaultValue, usage)
	shortFlagValue := flag.String(shortFlag, defaultValue, fmt.Sprintf("See -%v", longFlag))
	if *longFlagValue != defaultValue {
		return longFlagValue
	} else {
		return shortFlagValue
	}
}

func parseBoolFlag(longFlag, shortFlag string, defaultValue bool, usage string) (flagValue *bool) {
	longFlagValue := flag.Bool(longFlag, defaultValue, usage)
	shortFlagValue := flag.Bool(shortFlag, defaultValue, fmt.Sprintf("See -%v", longFlag))
	if *longFlagValue != defaultValue {
		return longFlagValue
	} else {
		return shortFlagValue
	}
}

func readConfig() (config Config) {
	configLocation := getConfigLocation()
	config = readYaml(configLocation)
	logging.Log(logging.Info, fmt.Sprintf("Config file from %v", configLocation))
	logging.Log(logging.Info, fmt.Sprintf("config.yaml: secretkey=***, files=%v", config.Files))
	return config
}

func getConfigLocation() (location string) {
	_, err := os.Stat(LOCAL_LOCATION)
	if !errors.Is(err, os.ErrNotExist) {
		return LOCAL_LOCATION
	}
	_, err = os.Stat(GLOBAL_LOCATION)
	if !errors.Is(err, os.ErrNotExist) {
		return GLOBAL_LOCATION
	}
	shared.CheckError(err, nil)
	return ""
}

func readYaml(configLocation string) (config Config) {
	data, err := ioutil.ReadFile(configLocation)
	shared.CheckError(err, nil)
	err = config.Parse(data)
	shared.CheckError(err, nil)
	return config
}

func mergeCliAndConfig(cliFlags CliFlags, config Config) (newConfig Config) {
	if cliFlags.SecretKey != "" {
		newConfig.SecretKey = cliFlags.SecretKey
	}
	newConfig.Encrypt = cliFlags.Encrypt
	newConfig.Decrypt = cliFlags.Decrypt
	newConfig.Files = config.Files
	logging.Log(logging.Info, fmt.Sprintf("Merged CLI and config.yaml: secretkey=***, encrypt=%v, decrypt=%v, files=%v", newConfig.Encrypt, newConfig.Decrypt, newConfig.Files))
	return newConfig
}

func checkConfigAndApplyDefaults(config Config) (newConfig Config) {
	newConfig = config
	if config.SecretKey == "" {
		shared.CheckError(fmt.Errorf("a secret key must be set"), nil)
	}
	if shared.CompareStringSlices(config.Files, []string{}) {
		newConfig.Files = []string{"secrets.geheim.yaml"}
	}
	if config.Encrypt && config.Decrypt {
		newConfig.Decrypt = false
	} else if !config.Encrypt && !config.Decrypt {
		newConfig.Encrypt = true
	}
	logging.Log(logging.Info, fmt.Sprintf("Final config with needed defaults: secretkey=***, encrypt=%v, decrypt=%v, files=%v", newConfig.Encrypt, newConfig.Decrypt, newConfig.Files))
	return newConfig
}
