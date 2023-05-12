hello:
	echo "Hello"
build:
	go build -o bin/party-calc ./cmd/main.go
run:
	go run ./cmd/main.go
compile:
	GOOS=windows GOARCH=386 go build -o bin/main-windows-386 ./cmd/main.go