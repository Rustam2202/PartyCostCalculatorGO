FROM golang:latest

WORKDIR /app

COPY . .

RUN go build -o /bin/party-calc ./cmd/main.go

ENTRYPOINT [ "/bin/party-calc" ]
