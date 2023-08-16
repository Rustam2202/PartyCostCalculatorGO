build:
	go build -o bin/party-calc ./cmd/main.go
run:
	go run ./cmd/main.go 
exe:
	./bin/party-calc
	
grpc-run:
	go run ./internal/server/grpc/server/server.go
grpcui-run:
	grpcui -plaintext localhost:50051
grpcui-kafka-run:
	grpcui -plaintext localhost:50052

docker-build:
	docker build --tag party-calc .
docker-run:
	docker run -p 8080:8080 party-calc
compose:
	docker-compose up

test:
	go test ./... -cover -coverprofile=coverage.out
test-cover-report:
	make test
	go tool cover -html=coverage.out

swag:
	swag init -g ./internal/server/http/server.go
lint:
	golangci-lint run

zookeeper-run:
	bin/windows/zookeeper-server-start.bat config/zookeeper.properties
kafka-run:
	make zookeeper-run
	bin/windows/kafka-server-start.bat config/server0.properties
	
proto:
	protoc --go_out=./internal/server/grpc/ --go-grpc_out=./internal/server/grpc/ ./protobuf/service.proto