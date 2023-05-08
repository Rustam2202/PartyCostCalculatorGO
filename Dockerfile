FROM golang:latest

WORKDIR ${workspaceFolder}

COPY go.mod ./

COPY *.go ./

RUN go build ./cmd
