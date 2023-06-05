build:
	go build -o bin/party-calc ./cmd/main.go
	
run:
	go run ./cmd/main.go -srvcfg=./internal/server/config/ -dbcfg=./internal/database/config

exe:
	./bin/party-calc -srvcfg=./internal/server/config/ -dbcfg=./internal/database/config

build-docker:
	docker build --tag party-calc .

run-docker:
	docker run -p 8080:8080 party-calc

compose:
	docker-compose up