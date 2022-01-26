package logging

import (
	"log"
	"os"
)

var (
	infoLogger  *log.Logger
	errorLogger *log.Logger
	logLevel    string
	lookup      bool
)

func init() {
	infoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
	setupLogLevel()
}

func setupLogLevel() {
	logLevel, lookup = os.LookupEnv(GEHEIM_LOG_LEVEL_ENV_VAR)
	if !lookup {
		logLevel = InfoLogLevel
	}
	if logLevel != NoLogsLevel && logLevel != InfoLogLevel && logLevel != DebugLogLevel {
		logLevel = InfoLogLevel
	}
}

func Log(logger, level, message string) {
	if level == logLevel || level < logLevel {
		switch logger {
		case Info:
			infoLogger.Println(message)
		case Error:
			errorLogger.Panicln(message)
		default:
			infoLogger.Println(message)
		}
	}
}
