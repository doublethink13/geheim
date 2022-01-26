package config

//nolint
import (
	"flag"
	"fmt"
	"os"

	"treuzedev/geheim/packages/logging"
	"treuzedev/geheim/packages/shared"
)

func Get(getLocation GetFirstExistingLocation, reader Reader) (config Config) {
	localLocation, globalLocation := setup()

	cliFlags := parseCliFlags()
	tmpConfig := readConfig(getLocation, localLocation, globalLocation, reader)
	mergedConfig := mergeCliAndConfig(cliFlags, tmpConfig)
	finalConfig := checkConfigAndApplyDefaults(mergedConfig)

	return finalConfig
}

func setup() (localLocation, globalLocation string) {
	home, err := os.UserHomeDir()
	shared.CheckError(err, nil)

	localLocation = ".geheim/config.yaml"
	globalLocation = fmt.Sprintf("%s/.geheim/config.yaml", home)

	logging.Log(
		logging.Info,
		logging.DebugLogLevel,
		fmt.Sprintf("Local config: %v", localLocation),
	)
	logging.Log(
		logging.Info,
		logging.DebugLogLevel,
		fmt.Sprintf("Global config: %v", globalLocation),
	)

	return localLocation, globalLocation
}

func parseCliFlags() (cliFlags CliFlags) {
	check := parseCheckFlag()
	secretkey := parseSecretKeyFlag()
	encrypt := parseEncryptFlag()
	decrypt := parseDecryptFlag()

	flag.Parse()

	cliFlags = CliFlags{
		Check:     check,
		SecretKey: secretkey,
		Encrypt:   encrypt,
		Decrypt:   decrypt,
	}

	logging.Log(
		logging.Info,
		logging.DebugLogLevel,
		fmt.Sprintf(
			"CLI flags: --check=%s --secretkey=*** --encrypt=%v --decrypt=%v",
			cliFlags.Check,
			cliFlags.Encrypt,
			cliFlags.Decrypt,
		),
	)

	return cliFlags
}

func parseCheckFlag() (checkFlag string) {
	checkUsage := "Whether to check if files are encrypted or decrypted. Defaults to an empty string, '', ie, its not active by default. If set to 'encrypted'/'e' or 'decrypted'/'d', checks if all files are in the specified state, and throws an error otherwise. When set, no encryption/decryption occurs" //nolint
	parseStringFlag(&checkFlag, "check", "c", "", checkUsage)

	return
}

func parseSecretKeyFlag() (secretkeyFlag string) {
	secretKeyUsage := "A key to encrypt/decrypt files. If not specified, the program will try to get one from local/global config file" //nolint
	parseStringFlag(&secretkeyFlag, "secretkey", "k", "", secretKeyUsage)

	return
}

func parseEncryptFlag() (encryptFlag bool) {
	encryptUsage := "Whether to encrypt the files defined in the config file. Defaults to 'false'. If both encrypt and decrypt flags are set to 'true', the encrypt flag takes precedence. If both the encrypt flag and the decrypt are set to 'false', the default behavior is to encrypt any unencrypted files, ie, the encrypt flag becomes 'true'" //nolint
	parseBoolFlag(&encryptFlag, "encrypt", "e", false, encryptUsage)

	return
}

func parseDecryptFlag() (decryptFlag bool) {
	decryptUsage := "Whether to decrypt the files defined in the config file. Defaults to 'false'. If both encrypt and decrypt flags are set to 'true', the encrypt flag takes precedence. If both the encrypt flag and the decrypt are set to 'false', the default behavior is to encrypt any unencrypted files, ie, the encrypt flag becomes 'true'" //nolint
	parseBoolFlag(&decryptFlag, "decrypt", "d", false, decryptUsage)

	return
}

func parseStringFlag(flagValue *string, longFlag, shortFlag, defaultValue, usage string) {
	flag.StringVar(
		flagValue,
		longFlag,
		defaultValue,
		usage,
	)
	flag.StringVar(
		flagValue,
		shortFlag,
		defaultValue,
		fmt.Sprintf("See -%v", longFlag),
	)
}

func parseBoolFlag(flagValue *bool, longFlag, shortFlag string, defaultValue bool, usage string) {
	flag.BoolVar(
		flagValue,
		longFlag,
		defaultValue,
		usage,
	)
	flag.BoolVar(
		flagValue,
		shortFlag,
		defaultValue,
		fmt.Sprintf("See -%v", longFlag),
	)
}

func readConfig(getLocation GetFirstExistingLocation, localLocation, globalLocation string, reader Reader) (config Config) {
	configLocation := getLocation([]string{localLocation, globalLocation})
	if configLocation != "" {
		config = readYaml(configLocation, reader)
	}

	logging.Log(
		logging.Info,
		logging.DebugLogLevel,
		fmt.Sprintf("Config file from %v", configLocation),
	)
	logging.Log(
		logging.Info,
		logging.DebugLogLevel,
		fmt.Sprintf("config.yaml: secretkey=***, files=%v", config.Files),
	)

	return config
}

// var getConfigLocation = func(localLocation, globalLocation string) (finalLocation string) {
// 	file, err := os.Stat(localLocation)
// 	if !errors.Is(err, os.ErrNotExist) {
// 		shared.CheckError(err, nil)
// 	}
// 	if file != nil {
// 		return localLocation
// 	}

// 	file, err = os.Stat(globalLocation)
// 	if !errors.Is(err, os.ErrNotExist) {
// 		shared.CheckError(err, nil)
// 	}
// 	if file != nil {
// 		return globalLocation
// 	}

// 	return ""
// }

func readYaml(configLocation string, reader Reader) (config Config) {
	data, err := reader(configLocation)
	shared.CheckError(err, nil)

	err = config.Parse(data)
	shared.CheckError(err, nil)

	return config
}

func mergeCliAndConfig(cliFlags CliFlags, config Config) (newConfig Config) {
	newConfig.Check = cliFlags.Check

	if cliFlags.SecretKey != "" {
		newConfig.SecretKey = cliFlags.SecretKey
	} else {
		newConfig.SecretKey = config.SecretKey
	}

	newConfig.Encrypt = cliFlags.Encrypt
	newConfig.Decrypt = cliFlags.Decrypt
	newConfig.Files = config.Files

	logging.Log(
		logging.Info,
		logging.DebugLogLevel,
		fmt.Sprintf(
			"Merged CLI and config.yaml: --check=%v secretkey=***, encrypt=%v, decrypt=%v, files=%v",
			newConfig.Check,
			newConfig.Encrypt,
			newConfig.Decrypt,
			newConfig.Files,
		),
	)

	return newConfig
}

func checkConfigAndApplyDefaults(config Config) (newConfig Config) {
	newConfig = config

	if config.SecretKey == "" && config.Check == "" {
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

	logging.Log(
		logging.Info,
		logging.DebugLogLevel,
		fmt.Sprintf(
			"Final config with needed defaults: check=%v, secretkey=***, encrypt=%v, decrypt=%v, files=%v",
			newConfig.Check,
			newConfig.Encrypt,
			newConfig.Decrypt,
			newConfig.Files,
		),
	)

	return newConfig
}
