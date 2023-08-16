FROM golang:latest

WORKDIR /app

COPY . .

COPY go.mod go.sum ./
RUN go mod download

RUN go build -o /bin/party-calc ./cmd/main.go

EXPOSE 8080

ENTRYPOINT [ "/bin/party-calc" ]
