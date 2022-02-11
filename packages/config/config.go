package config

//nolint
import (
	"flag"
	"fmt"

	"treuzedev/geheim/packages/logging"
	"treuzedev/geheim/packages/shared"
)

func Get(configLocation string, reader Reader) (config Config) {
	cliFlags := parseCliFlags()
	tmpConfig := readConfig(configLocation, reader)
	mergedConfig := mergeCliAndConfig(cliFlags, tmpConfig)
	finalConfig := checkConfigAndApplyDefaults(mergedConfig)

	return finalConfig
}

func parseCliFlags() (cliFlags CliFlags) {
	logger := logging.GetLogger()

	setFlagUsage()

	check := parseCheckFlag()
	secretkey := parseSecretKeyFlag()
	encrypt := parseEncryptFlag()
	decrypt := parseDecryptFlag()

	flag.Parse()

	cliFlags = CliFlags{
		Check:     *check,
		SecretKey: *secretkey,
		Encrypt:   *encrypt,
		Decrypt:   *decrypt,
	}

	logger.Log(
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

func setFlagUsage() {
	usage := `geheim usage:

	--check, -c
		Whether to check if files are encrypted or decrypted
		Defaults to an empty string, '', ie, its not active by default
		If set to 'encrypted'/'e' or 'decrypted'/'d', checks if all files are in the specified state
		When set, no encryption/decryption occurs
	
	--secretkey, -k 
		A key to encrypt/decrypt files
		If not specified, the program will try to get one from local/global config file
	
	--encrypt, -e
		Whether to encrypt the files defined in the config file
		Defaults to 'false'
		If both encrypt and decrypt flags are set to 'true', the encrypt flag takes precedence
		If both encrypt and decrypt flags are set to 'false', the encrypt flag becomes 'true'
	
	--decrypt, -d
		Whether to decrypt the files defined in the config file
		Defaults to 'false'
		If both encrypt and decrypt flags are set to 'true', the encrypt flag takes precedence
		If both encrypt and decrypt flags are set to 'false', the encrypt flag becomes 'true'
	
`
	//nolint:forbidigo
	flag.Usage = func() {
		fmt.Print(usage)
	}
}

func parseCheckFlag() (flag *string) {
	var checkFlag string

	parseStringFlag(&checkFlag, "check", "c", "")

	return &checkFlag
}

func parseSecretKeyFlag() (flag *string) {
	var secretkeyFlag string

	parseStringFlag(&secretkeyFlag, "secretkey", "k", "")

	return &secretkeyFlag
}

func parseEncryptFlag() (flag *bool) {
	var encryptFlag bool

	parseBoolFlag(&encryptFlag, "encrypt", "e", false)

	return &encryptFlag
}

func parseDecryptFlag() (flag *bool) {
	var decryptFlag bool

	parseBoolFlag(&decryptFlag, "decrypt", "d", false)

	return &decryptFlag
}

func parseStringFlag(flagValue *string, longFlag, shortFlag, defaultValue string) {
	flag.StringVar(
		flagValue,
		longFlag,
		defaultValue,
		"",
	)
	flag.StringVar(
		flagValue,
		shortFlag,
		defaultValue,
		"",
	)
}

func parseBoolFlag(flagValue *bool, longFlag, shortFlag string, defaultValue bool) {
	flag.BoolVar(
		flagValue,
		longFlag,
		defaultValue,
		"",
	)
	flag.BoolVar(
		flagValue,
		shortFlag,
		defaultValue,
		"",
	)
}

func readConfig(configLocation string, reader Reader) (config Config) {
	logger := logging.GetLogger()

	if configLocation != "" {
		config = readYaml(configLocation, reader)
	}

	logger.Log(
		logging.Info,
		logging.DebugLogLevel,
		fmt.Sprintf("Config file from %v", configLocation),
	)
	logger.Log(
		logging.Info,
		logging.DebugLogLevel,
		fmt.Sprintf("config.yaml: secretkey=***, files=%v", config.Files),
	)

	return config
}

func readYaml(configLocation string, reader Reader) (config Config) {
	data, err := reader(configLocation)
	shared.CheckError(err, nil)

	err = config.Parse(data)
	shared.CheckError(err, nil)

	return config
}

func mergeCliAndConfig(cliFlags CliFlags, config Config) (newConfig Config) {
	logger := logging.GetLogger()

	newConfig.Check = cliFlags.Check

	if cliFlags.SecretKey != "" {
		newConfig.SecretKey = cliFlags.SecretKey
	} else {
		newConfig.SecretKey = config.SecretKey
	}

	newConfig.Encrypt = cliFlags.Encrypt
	newConfig.Decrypt = cliFlags.Decrypt
	newConfig.Files = config.Files

	logger.Log(
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
	logger := logging.GetLogger()

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

	logger.Log(
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
