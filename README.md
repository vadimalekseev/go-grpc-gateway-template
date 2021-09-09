# go-grpc-gateway-template

* [![Go Report Card](https://goreportcard.com/badge/github.com/aleksvdim/go-grpc-gateway-template?style=flat-square)](https://goreportcard.com/report/github.com/aleksvdim/go-grpc-gateway-template)

* [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/aleksvdim/go-grpc-gateway-template)

* [![Go Reference](https://pkg.go.dev/badge/github.com/aleksvdim/go-grpc-gateway-template.svg)](https://pkg.go.dev/github.com/aleksvdim/go-grpc-gateway-template)

* [![Release](https://img.shields.io/github/release/aleksvdim/go-grpc-gateway-template.svg?style=flat-square)](https://github.com/aleksvdim/go-grpc-gateway-template/releases/latest)

## TODO

* health check;
* metrics

## Goal

Bringing practices together when building yet another web service.

## Example usage

Project contains echo server with PostgreSQL integration. It supports gRPC and HTTP transports 
using [gRPC to JSON proxy generator](https://github.com/grpc-ecosystem/grpc-gateway). 
Proto files generated with [buf](https://buf.build/). 

## Requirements

The project has been tested on Go 1.17. It uses some new features such as `go install`, `go:embed` (for third party dependencies like Swagger UI).
Other binary dependencies will be downloaded to the `bin` folder.

## First steps

### Third party dependencies

To download third party dependencies (e.g. buf, proto-gen-go, goose, Swagger UI) run `make deps`.

### Build project

Run `make build`. It will download Swagger UI if it does not exist and build `echoapi/main.go` file.

### Lint before commit 

Run `make lint`. It will check *.proto files with [buf](https://buf.build/) and *.go files with golangci-lint

### Set up dependent services

Run `docker compose up` to run PostgreSQL.

### Run tests

```bash
# Run unit tests
make test

# Run integration tests
# Paste your connection string. 
# If you ran docker compose up, you can copy-paste the code below.
export DSN="user=postgres password=postgres database=postgres sslmode=disable" 

# run migrations
GOOSE_DRIVER=postgres GOOSE_DBSTRING=$DSN make migrate

# run integration tests
make test-integration
``` 

### Create migration

```bash
GOOSE_DRIVER=postgres GOOSE_DBSTRING=$DSN make migration
```

The new migration will be added to the migration folder. Do not forget to rename it.

### Apply migrations

```bash
GOOSE_DRIVER=postgres GOOSE_DBSTRING=$DSN make migrate
```

### Create new service

* Add a new service to `proto/echo/v1/echo.proto` file or add a new proto file to the proto folder;

* Then run `make` - it will generate new *.pb.go files in the pkg folder.;
See [generate code with buf](https://docs.buf.build/tour/generate-code) for more information.

## After fork

1. Rename the name of the module in the `go.mod` file;
2. Replace the default `go_package_prefix` with the new module name in buf.gen.yaml file;
3. Fix badges references in README.md file;
4. Change author in LICENSE file;
