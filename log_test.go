package log

import (
	"testing"
)

func TestConsoleLog(t *testing.T) {
	Info("Console message")
	Info("Console message with parameters: %s %d", "test", 100)
}
