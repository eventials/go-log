package consolelog

import (
	"fmt"
	"log"
	"os"

	golog "github.com/eventials/golog"
)

// ConsoleLogger will print all logs into the console.
type ConsoleLogger struct {
	std *log.Logger
}

// NewConsoleLogger return a ConsoleLogger.
func NewConsoleLogger(prefix string) *ConsoleLogger {
	return &ConsoleLogger{
		log.New(os.Stderr, prefix, log.LstdFlags),
	}
}

// Name return the Logger name.
func (logger *ConsoleLogger) Name() string {
	return "console"
}

// Panic print logs into the console.
func (logger *ConsoleLogger) Panic(format string, args ...interface{}) {
	logger.all(golog.LPanic, format, args...)
}

// Fatal print logs into the console.
func (logger *ConsoleLogger) Fatal(format string, args ...interface{}) {
	logger.all(golog.LFatal, format, args...)
}

// Error print logs into the console.
func (logger *ConsoleLogger) Error(format string, args ...interface{}) {
	logger.all(golog.LError, format, args...)
}

// Warning print logs into the console.
func (logger *ConsoleLogger) Warning(format string, args ...interface{}) {
	logger.all(golog.LWarning, format, args...)
}

// Info print logs into the console.
func (logger *ConsoleLogger) Info(format string, args ...interface{}) {
	logger.all(golog.LInfo, format, args...)
}

func (logger *ConsoleLogger) all(logLevel golog.LogLevel, format string, args ...interface{}) {
	level := fmt.Sprintf("[%s] ", logLevel)
	format = level + format + "\n"

	msg := fmt.Sprintf(format, args...)
	logger.std.Output(2, msg)
}
