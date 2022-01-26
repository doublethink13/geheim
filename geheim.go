package geheim

//nolint
import (
	"errors"
	"fmt"
	"io/ioutil"
	"os"

	"treuzedev/geheim/packages/config"
	"treuzedev/geheim/packages/decrypt"
	"treuzedev/geheim/packages/encrypt"
	"treuzedev/geheim/packages/logging"
	"treuzedev/geheim/packages/shared"
)

const (
	encrypted string = "encrypted"
	e         string = "e"
	decrypted string = "decrypted"
	d         string = "d"
)

func main() {
	config := config.Get(getConfigFileLocation, ioutil.ReadFile)
	Geheim(config)
}

func Geheim(config config.Config) {
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
				errorMessage := fmt.Errorf("%v is not encrypted, and it should", filePath)
				shared.CheckError(errorMessage, nil)
			}
		case d:
			logging.Log(logging.Info, logging.DebugLogLevel, fmt.Sprintf("Checking if file '%v' is decrypted", filePath))

			if shared.CheckIfEncrypted(filePath) {
				errorMessage := fmt.Errorf("%v is not decrypted, and it should", filePath)
				shared.CheckError(errorMessage, nil)
			}
		default:
			errorMessage := fmt.Errorf("something is wrong with the binary - checkState")
			shared.CheckError(errorMessage, nil)
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

func getConfigFileLocation(locations []string) (firstExistingLocation string) {
	for _, location := range locations {
		file, err := os.Stat(location)

		if !errors.Is(err, os.ErrNotExist) {
			shared.CheckError(err, nil)
		}

		if file != nil {
			return location
		}
	}

	return ""
}
