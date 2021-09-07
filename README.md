# go-grpc-gateway-template

* [![Go Report Card](https://goreportcard.com/badge/github.com/aleksvdim/go-grpc-gateway-template?style=flat-square)](https://goreportcard.com/report/github.com/aleksvdim/go-grpc-gateway-template)

* [![Go Doc](https://img.shields.io/badge/godoc-reference-blue.svg?style=flat-square)](http://godoc.org/github.com/aleksvdim/go-grpc-gateway-template)

* [![Go Reference](https://pkg.go.dev/badge/github.com/aleksvdim/go-grpc-gateway-template.svg)](https://pkg.go.dev/github.com/aleksvdim/go-grpc-gateway-template)

* [![Release](https://img.shields.io/github/release/aleksvdim/go-grpc-gateway-template.svg?style=flat-square)](https://github.com/aleksvdim/go-grpc-gateway-template/releases/latest)

## Requirements

The project has been tested on Go 1.17. It uses some new features such as `go install`, `go:embed` (for third_party dependencies).
Other binary dependencies will be downloaded to the `bin` folder.

## First steps

### Binary dependencies

To download binary dependencies (e.g. buf, proto-gen-go, goose) run `make bin-deps`.

### Run tests

```bash
# run unit tests
make test

# run integration tests
# paste your connection string
export DSN="user=postgres password=postgres database=postgres sslmode=disable"

# run migrations
GOOSE_DRIVER=postgres GOOSE_DBSTRING=$DSN make migrate

# run integration tests
make test-integration
``` 

## Create migration

```bash
GOOSE_DRIVER=postgres GOOSE_DBSTRING=$DSN make migration
```
