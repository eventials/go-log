package consolelog

import (
	"testing"
)

func TestConsoleLog(t *testing.T) {
	logger := NewConsoleLogger("")
	logger.Info("Console message")
	logger.Info("Console message with parameters: %s %d", "test", 100)
}
