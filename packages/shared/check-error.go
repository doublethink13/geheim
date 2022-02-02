package shared

//nolint
import (
	"os"

	"treuzedev/geheim/packages/logging"
)

func CheckError(err error, filePath *string) {
	if logger := logging.GetLogger(); err != nil {
		if filePath != nil {
			err = os.Remove(*filePath)
			if err != nil {
				logger.Log(logging.Error, logging.InfoLogLevel, err.Error())
			}
		}

		logger.Log(logging.Error, logging.InfoLogLevel, err.Error())
	}
}
