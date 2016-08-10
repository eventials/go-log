package log

import (
	"fmt"
	"os"
)

// LogLevel is a string type for log level.
type LogLevel string

// Log levels
const (
	LPanic   LogLevel = "PANIC"
	LFatal   LogLevel = "FATAL"
	LError   LogLevel = "ERROR"
	LWarning LogLevel = "WARNING"
	LInfo    LogLevel = "INFO"
)

// Logger defines the interface to logger methods.
type Logger interface {
	Panic(format string, args ...interface{})
	Fatal(format string, args ...interface{})
	Error(format string, args ...interface{})
	Warning(format string, args ...interface{})
	Info(format string, args ...interface{})
}

var (
	handlers = make(map[string]Logger)
)

// AddHandler adds a handler to the given loggerName.
func AddHandler(loggerName string, args ...interface{}) error {
	var logger Logger

	switch loggerName {
	case "console":
		logger = newConsoleLogger()
	case "sentry":
		if len(args) == 0 {
			return fmt.Errorf("Missing paramenter: DSN")
		}

		dsnURL, ok := args[0].(string)

		if !ok {
			return fmt.Errorf("DSN paramenter must be string")
		}

		logger = newSentryLogger(dsnURL)
	}

	handlers[loggerName] = logger
	return nil
}

// Panic log.
func Panic(format string, args ...interface{}) {
	for _, logger := range handlers {
		logger.Panic(format, args...)
	}

	msg := fmt.Sprintf(format, args...)
	panic(msg)
}

// Fatal log.
func Fatal(format string, args ...interface{}) {
	for _, logger := range handlers {
		logger.Fatal(format, args...)
	}

	os.Exit(1)
}

// Error log.
func Error(format string, args ...interface{}) {
	for _, logger := range handlers {
		logger.Error(format, args...)
	}
}

// Warning log.
func Warning(format string, args ...interface{}) {
	for _, logger := range handlers {
		logger.Warning(format, args...)
	}
}

// Info log.
func Info(format string, args ...interface{}) {
	for _, logger := range handlers {
		logger.Info(format, args...)
	}
}
