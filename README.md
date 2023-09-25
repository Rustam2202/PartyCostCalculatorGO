
# Party Cost Calculator

Server application for manage event data and persons spent money. 

## API Specification
https://rustam2202.github.io/PartyCostCalculatorGO/

## Run Locally

Config parameters in ```cmd/config.yaml``` and start servers:

```bash
  make run
```

<!-- ## Up in docker

```bash
  make compose
``` -->

## Running Tests

To run tests and generate HTML coverage report 

```bash
  make test-cover-report
```


## Tech Stack

- [Gin-Gonic](https://github.com/gin-gonic/gin)
- [Viper](github.com/spf13/viper)
- [Zap Logger](https://github.com/uber-go/zap)
- PostgreSQL
- [PGX](github.com/jackc/pgx)
- [Mock Tests](github.com/pashagolub/pgxmock)
- [Kafka](https://github.com/segmentio/kafka-go)
<!-- - [Docker Compose](https://docs.docker.com/compose) -->
- REST-API
- [gRPC](https://github.com/grpc/grpc-go)
- [ProtoBuf](https://github.com/protocolbuffers/protobuf-go)
- [Swagger](https://github.com/swaggo/swag)
- Graceful Shutdown
- Panic Recovery
- [Prometheus](github.com/prometheus/client_golang/prometheus)


## Roadmap
- Fix docker-compose
- More tests 
