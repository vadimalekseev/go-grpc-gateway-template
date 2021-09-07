# go-grpc-gateway-template

## Requirements

The project has been tested on Go 1.17. It uses some new features such as `go install`, `go:embed` (for third_party dependencies).
Other binary dependencies will be downloaded to the `bin` folder.

## First steps

### Binary dependencies

To download binary dependencies (e.g. buf, proto-gen-go, goose) run `make bin-deps`.

### Run tests

```bash
# paste your connection string
export DSN="user=postgres password=postgres database=postgres sslmode=disable"

# run migrations
GOOSE_DRIVER=postgres GOOSE_DBSTRING=$DSN make migrate

# run tests
make test

# run integration tests
make test-integration
``` 

## Create migration

```bash
GOOSE_DRIVER=postgres GOOSE_DBSTRING="user=postgres password=postgres database=go-sink sslmode=disable" make migration
```
