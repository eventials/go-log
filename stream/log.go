package streamlog

import (
	"fmt"
	"io"

	golog "github.com/eventials/golog"
)

// StreamLogger will write all logs into a stream writer.
type StreamLogger struct {
	writer io.Writer
}

// NewStreamLogger return a StreamLogger.
func NewStreamLogger(writer io.Writer) *StreamLogger {
	return &StreamLogger{
		writer,
	}
}

// Name return the Logger name.
func (logger *StreamLogger) Name() string {
	return "stream"
}

// Panic write logs into a stream.
func (logger *StreamLogger) Panic(format string, args ...interface{}) {
	logger.all(golog.LPanic, format, args...)
}

// Fatal write logs into a stream.
func (logger *StreamLogger) Fatal(format string, args ...interface{}) {
	logger.all(golog.LFatal, format, args...)
}

// Error write logs into a stream.
func (logger *StreamLogger) Error(format string, args ...interface{}) {
	logger.all(golog.LError, format, args...)
}

// Warning write logs into a stream.
func (logger *StreamLogger) Warning(format string, args ...interface{}) {
	logger.all(golog.LWarning, format, args...)
}

// Info write logs into a stream.
func (logger *StreamLogger) Info(format string, args ...interface{}) {
	logger.all(golog.LInfo, format, args...)
}

func (logger *StreamLogger) all(logLevel golog.LogLevel, format string, args ...interface{}) {
	level := fmt.Sprintf("[%s] ", logLevel)
	format = level + format + "\n"

	msg := fmt.Sprintf(format, args...)
	logger.writer.Write([]byte(msg))
}
