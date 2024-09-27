# REP Test API

This project implements a simple API for a test using Go, Gin, Docker and redis. The API provides endpoints for packs management, including. It also includes end-to-end tests and mocks for database

## Getting Started

### Prerequisites

- Go 1.21.5 or later
- Docker
- Docker Compose

### Installation

1. Clone the repository
2. Install dependencies: go mod tidy

### Running the Application

1. Re-Generate Swagger documentation (only if you change the api): 
```
swag init -g cmd/main.go
```
2. Start Docker services: 
```
docker-compose up -d
```
3. Run the application: 
```
go run ./cmd/main.go
```
 3.1. You can also run the application with vsCode (a config file is provided in the root of the project)
    

### Running Tests

1. Run unit tests: 
```
go test ./...
```

### Generating Mocks

Generate mocks for database and service interfaces:

```
mockgen -source=internal/core/port/pack.go -destination=internal/core/port/mock_pack_port.go -package=port
mockgen -source=internal/core/port/config.go -destination=internal/core/port/mock_config_port.go -package=port
mockgen -source=internal/core/port/cache.go -destination=internal/core/port/mock_cache_port.go -package=port

```

## API Documentation

The API documentation is generated using Swagger. After running the application, you can access the Swagger UI at:

http://localhost:8080/swagger/index.html

## Project Structure

- `cmd/main.go`: Entry point of the application.
- `internal/adapter/handler/http`: Contains the HTTP handlers for the API endpoints.
- `internal/core/service`: Contains the business logic and service layer.
- `internal/core/port`: Contains the interfaces for the API.
- `internal/core/domain`: Contains the domain objects for the API.

