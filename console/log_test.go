package consolelog

import (
	"testing"

	"github.com/eventials/golog/console"
)

func TestConsoleLog(t *testing.T) {
	logger := consolelog.NewConsoleLogger("")
	logger.Info("Console message")
	logger.Info("Console message with parameters: %s %d", "test", 100)
}
