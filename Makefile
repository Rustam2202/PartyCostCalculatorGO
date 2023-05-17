build:
	go build -o bin/party-calc ./cmd/main.go
	
run:
	go run ./cmd/main.go

exe:
	./bin/party-calc -config=.

build-dockerfile:
	docker build --tag party-calc .

run-dockerfile:
	docker run -p 8080:8080 party-calc
