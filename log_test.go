package log

import (
	"testing"

	"github.com/eventials/golog"
	"github.com/eventials/golog/console"
)

func TestConsoleLog(t *testing.T) {
	logger := consolelog.NewConsoleLogger("")
	log.AddLogger(logger)
	log.Info("Console message")
	log.Info("Console message with parameters: %s %d", "test", 100)
}
