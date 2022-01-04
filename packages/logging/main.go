package logging

import (
	"log"
	"os"
)

var infoLogger *log.Logger
var errorLogger *log.Logger
var logLevel string
var lookup bool

func init() {
	infoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
	logLevel, lookup = os.LookupEnv("GEHEIM_LOG_LEVEL")
	if !lookup {
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
