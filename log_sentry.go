package log

import (
	"fmt"

	"github.com/getsentry/raven-go"
)

// SentryLogger will ship all error, panic and fatal log to Sentry.
type SentryLogger struct {
}

func newSentryLogger(dsnURL string) *SentryLogger {
	raven.SetDSN(dsnURL)
	return &SentryLogger{}
}

// Panic will ship panic logs to Sentry.
func (logger *SentryLogger) Panic(format string, args ...interface{}) {
	logger.captureError(format, args...)
}

// Fatal will ship fatal logs to Sentry.
func (logger *SentryLogger) Fatal(format string, args ...interface{}) {
	logger.captureError(format, args...)
}

// Error will ship error logs to Sentry.
func (logger *SentryLogger) Error(format string, args ...interface{}) {
	logger.captureError(format, args...)
}

// Warning logs are ignored by this handler.
func (logger *SentryLogger) Warning(format string, args ...interface{}) {}

// Info logs are ignored by this handler.
func (logger *SentryLogger) Info(format string, args ...interface{}) {}

func (logger *SentryLogger) captureError(format string, args ...interface{}) {
	err := fmt.Errorf(format, args...)
	raven.CaptureErrorAndWait(err, nil)
}
