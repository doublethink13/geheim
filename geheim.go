package main

import (
	"fmt"
	"treuzedev/geheim/packages/config"
	"treuzedev/geheim/packages/decrypt"
	"treuzedev/geheim/packages/encrypt"
	"treuzedev/geheim/packages/logging"
	"treuzedev/geheim/packages/shared"
)

const encrypted string = "encrypted"
const e string = "e"
const decrypted string = "decrypted"
const d string = "d"

func main() {
	config := config.Get()
	geheim(config)
}

func geheim(config config.Config) {
	if config.Check != "" {
		checkIfEncryptedOrDecrypted(config)
	} else {
		workOnFiles(config)
	}
}

func checkIfEncryptedOrDecrypted(config config.Config) {
	switch config.Check {
	case encrypted:
		checkState(e, config)
	case e:
		checkState(e, config)
	case decrypted:
		checkState(d, config)
	case d:
		checkState(d, config)
	default:
		shared.CheckError(fmt.Errorf("check flag needs to be one of '%v'/'%v' or '%v'/'%v'", encrypted, e, decrypted, d), nil)
	}
}

func checkState(state string, config config.Config) {
	for _, filePath := range config.Files {
		switch state {
		case e:
			logging.Log(logging.Info, logging.DebugLogLevel, fmt.Sprintf("Checking if file '%v' is encrypted", filePath))
			if !shared.CheckIfEncrypted(filePath) {
				shared.CheckError(fmt.Errorf("%v is not encrypted, and it should", filePath), nil)
			}
		case d:
			logging.Log(logging.Info, logging.DebugLogLevel, fmt.Sprintf("Checking if file '%v' is decrypted", filePath))
			if shared.CheckIfEncrypted(filePath) {
				shared.CheckError(fmt.Errorf("%v is not decrypted, and it should", filePath), nil)
			}
		default:
			shared.CheckError(fmt.Errorf("something is wrong with the binary - checkState"), nil)
		}
	}
}

func workOnFiles(config config.Config) {
	for _, filePath := range config.Files {
		if config.Encrypt {
			if !shared.CheckIfEncrypted(filePath) {
				encrypt.EncryptFile(filePath, config.SecretKey)
			}
		}
		if config.Decrypt {
			if shared.CheckIfEncrypted(filePath) {
				decrypt.Decrypt(filePath, config.SecretKey)
			}
		}
	}
}
