build:
	go build -o bin/party-calc ./cmd/main.go
	
run:
	go run ./cmd/main.go 

exe:
	./bin/party-calc

build-docker:
	docker build --tag party-calc .

run-docker:
	docker run -p 8080:8080 party-calc

compose:
	docker-compose up