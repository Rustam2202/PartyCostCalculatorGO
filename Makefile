build:
	go build -o bin/party-calc ./cmd/main.go
run:
	go run ./cmd/main.go -confpath=./cmd
exe:
	./bin/party-calc
	
docker-build:
	docker build --tag party-calc .
docker-run:
	docker run -p 8080:8080 -e DB_HOST=127.0.0.1 -e DB_PORT=5432 -e DB_USER="postgres" -e DB_PASSWORD="password" -e DB_NAME="partycalc"  party-calc
compose:
	docker-compose up -d --build

test:
	go test ./... -cover -coverprofile=coverage.out
test-cover-report:
	make test
	go tool cover -html=coverage.out

grpcui-run:
	grpcui -plaintext localhost:50051
swag:
	swag fmt
	swag init -g ./internal/server/http/server.go
	npx @redocly/cli build-docs ./docs/swagger.json -o ./docs/index.html
lint:
	golangci-lint run
proto:
	protoc --go_out=./internal/server/grpc/ --go-grpc_out=./internal/server/grpc/ ./protobuf/service.proto
	
# zookeeper-run:
# 	bin/zookeeper-server-start.sh config/zookeeper.properties
# 	bin/windows/zookeeper-server-start.bat config/zookeeper.properties
# kafka-run:
# 	bin/kafka-server-start.sh config/server.properties
# 	bin/windows/kafka-server-start.bat config/server.properties
