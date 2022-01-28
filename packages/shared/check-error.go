package shared

import (
	"os"
	"treuzedev/geheim/packages/logging"
)

func CheckError(e error, filePath *string) {
	if e != nil {
		if filePath != nil {
			err := os.Remove(*filePath)
			if err != nil {
				logging.logger(logging.Error, logging.InfoLogLevel, e.Error())
			}
		}
		logging.logger(logging.Error, logging.InfoLogLevel, e.Error())
	}
}
