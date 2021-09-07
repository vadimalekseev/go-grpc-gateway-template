# go-grpc-gateway-template

# Contribute

## Bin deps

To download binary dependencies (e.g. proto-gen-go, goose) run make bin-deps.

To run tests set up your database and run:

```bash
export TEST_DSN=<your_connection_string> # example: "user=postgres password=postgres database=go-echo sslmode=disable"
./bin/goose -dir migrations postgres $TEST_DSN up
go test ./...
``` 
