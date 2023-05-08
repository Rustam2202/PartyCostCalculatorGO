FROM golang:latest

WORKDIR /app

COPY go.mod ./

COPY *.go ./

RUN go build ./cmd