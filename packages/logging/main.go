package logging

import (
	"log"
	"os"
)

var infoLogger *log.Logger
var errorLogger *log.Logger

func init() {
	infoLogger = log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime)
	errorLogger = log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime)
}

func Log(level, message string) {
	switch level {
	case Info:
		infoLogger.Println(message)
	case Error:
		errorLogger.Panicln(message)
	default:
		infoLogger.Println(message)
	}
}
