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

test:
	go test ./... -cover -coverprofile=coverage.out
test-cover-report:
	make test
	go tool cover -html=coverage.out

swag:
	swag init -g ./internal/server/server.go
lint:
	golangci-lint run
	
kafka-producer-run:
	go run ./internal/grpc/producer .
proto:
	protoc --go_out=./internal/server/grpc/ --go-grpc_out=./internal/server/grpc/ ./protobuf/service.proto