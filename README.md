# go-log [![Build Status](https://travis-ci.org/eventials/golog.svg?branch=master)](https://travis-ci.org/eventials/golog) [![Go Report Card](https://goreportcard.com/badge/github.com/eventials/golog)](https://goreportcard.com/report/github.com/eventials/golog) [![GoDoc](https://godoc.org/github.com/eventials/golog?status.svg)](http://godoc.org/github.com/eventials/golog)

Logging library for golang projects.

## Example

```go
logger := consolelog.NewConsoleLogger("")

log.AddLogger(logger)

log.Info("info message")
log.Info("info message with parameters: %d", 100)
log.Error("error: %s", "an error")
```

## Loggers

| Logger | Description | Dependency |
|:------:|:-----------:|:----------:|
| Console       | Logs everything to Stderr                | None   |
| Sentry        | Logs errors, panics and fatals to Sentry | [raven-go](https://github.com/getsentry/raven-go) |
| Stream        | Logs everything to io.Stream             | None   |
| TimedRotating | Logs everything to rolling files         | [lumberjack](https://github.com/natefinch/lumberjack) |

## License

MIT
