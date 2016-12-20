package streamlog

import (
	"log"
	"testing"
)

type StreamMock struct {
}

func (s *StreamMock) Write(p []byte) (n int, err error) {
	log.Printf(string(p[:]))
	return 0, nil
}

func TestConsoleLog(t *testing.T) {
	logger := NewStreamLogger(&StreamMock{})
	logger.Info("Stream message")
	logger.Info("Stream message with parameters: %s %d", "test", 100)
}
