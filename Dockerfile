FROM golang:1.7

RUN mkdir -p /go/src/github.com/eventials/golog

WORKDIR /go/src/github.com/eventials/golog

RUN go get github.com/getsentry/raven-go
