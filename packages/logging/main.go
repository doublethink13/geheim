package logging

import (
	"log"
	"os"
)

const (
	Info  = "info"
	Error = "error"
)

const (
	NoLogsLevel   = "0"
	InfoLogLevel  = "1"
	DebugLogLevel = "2"
)

const GEHEIM_LOG_LEVEL_ENV_VAR = "GEHEIM_LOG_LEVEL"

type GeheimLogger struct {
	logLevel    string
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

func (l *GeheimLogger) Log(logger, level, message string) {
	if level == l.logLevel || level < l.logLevel {
		switch logger {
		case Info:
			l.infoLogger.Println(message)
		case Error:
			l.errorLogger.Panicln(message)
		default:
			l.infoLogger.Println(message)
		}
	}
}

func NewGeheimLogger() GeheimLogger {
	logLevel, lookup := os.LookupEnv(GEHEIM_LOG_LEVEL_ENV_VAR)

	if !lookup {
		logLevel = InfoLogLevel
	}

	if logLevel != NoLogsLevel && logLevel != InfoLogLevel && logLevel != DebugLogLevel {
		logLevel = InfoLogLevel
	}

	return GeheimLogger{
		logLevel:    logLevel,
		infoLogger:  log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime),
	}
}
