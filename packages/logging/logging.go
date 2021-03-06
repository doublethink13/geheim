package logging

import (
	"log"
	"os"
)

//nolint
var logger *GeheimLogger

const (
	Info  = "info"
	Error = "error"
)

const (
	NoLogsLevel    = "0"
	InfoLogLevel   = "1"
	DebugLogLevel  = "2"
	geheimLogLevel = "GEHEIM_LOG_LEVEL"
)

type GeheimLogger struct {
	logLevel    string
	infoLogger  *log.Logger
	errorLogger *log.Logger
}

type LoggerProvider func() *GeheimLogger

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

func newGeheimLogger() *GeheimLogger {
	logLevel, lookup := os.LookupEnv(geheimLogLevel)

	if !lookup {
		logLevel = InfoLogLevel
	}

	if logLevel != NoLogsLevel && logLevel != InfoLogLevel && logLevel != DebugLogLevel {
		logLevel = InfoLogLevel
	}

	logger := GeheimLogger{
		logLevel:    logLevel,
		infoLogger:  log.New(os.Stderr, "INFO: ", log.Ldate|log.Ltime),
		errorLogger: log.New(os.Stderr, "ERROR: ", log.Ldate|log.Ltime),
	}

	return &logger
}

func GetLogger() *GeheimLogger {
	if logger == nil {
		logger = newGeheimLogger()
	}

	return logger
}
