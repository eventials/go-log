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
	Name() string
}

var (
	handlers = make(map[string]Logger)
)

// AddLogger adds a Logger handlers.
func AddLogger(logger Logger) error {
	if logger == nil {
		return fmt.Errorf("logger can't be nil.")
	}

	handlers[logger.Name()] = logger

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
