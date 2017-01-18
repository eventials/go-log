package filelog

import (
	"fmt"
	"time"

	golog "github.com/eventials/golog"
	"gopkg.in/natefinch/lumberjack.v2"
)

// TimedRotatingLogger will write all logs into a stream writer.
type TimedRotatingLogger struct {
	logger *lumberjack.Logger
}

// NewTimedRotatingLogger return a TimedRotatingl.
func NewTimedRotatingLogger(filename string, maxsize int, maxage int, maxbackups int) *TimedRotatingLogger {
	return &TimedRotatingLogger{
		&lumberjack.Logger{
			Filename:   filename,
			MaxSize:    maxsize, // megabytes
			MaxBackups: maxbackups,
			MaxAge:     maxage, //days
		},
	}
}

// Name return the Logger name.
func (l *TimedRotatingLogger) Name() string {
	return "stream"
}

// Panic write logs into a stream.
func (l *TimedRotatingLogger) Panic(format string, args ...interface{}) {
	l.all(golog.LPanic, format, args...)
}

// Fatal write logs into a stream.
func (l *TimedRotatingLogger) Fatal(format string, args ...interface{}) {
	l.all(golog.LFatal, format, args...)
}

// Error write logs into a stream.
func (l *TimedRotatingLogger) Error(format string, args ...interface{}) {
	l.all(golog.LError, format, args...)
}

// Warning write logs into a stream.
func (l *TimedRotatingLogger) Warning(format string, args ...interface{}) {
	l.all(golog.LWarning, format, args...)
}

// Info write logs into a stream.
func (l *TimedRotatingLogger) Info(format string, args ...interface{}) {
	l.all(golog.LInfo, format, args...)
}

func (l *TimedRotatingLogger) all(logLevel golog.LogLevel, format string, args ...interface{}) {
	header := fmt.Sprintf("%s [%s] ", time.Now().Format("2006/01/02 15:04:05"), logLevel)
	format = header + format + "\n"

	msg := fmt.Sprintf(format, args...)
	l.logger.Write([]byte(msg))
}
