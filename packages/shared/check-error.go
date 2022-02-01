package shared

//nolint
import (
	"os"

	"treuzedev/geheim/packages/logging"
)

func CheckError(e error, filePath *string) {
	if logger := logging.GetLogger(); e != nil {
		if filePath != nil {
			err := os.Remove(*filePath)
			if err != nil {
				logger.Log(logging.Error, logging.InfoLogLevel, e.Error())
			}
		}

		logger.Log(logging.Error, logging.InfoLogLevel, e.Error())
	}
}
