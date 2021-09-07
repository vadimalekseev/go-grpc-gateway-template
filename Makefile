LOCAL_BIN := $(CURDIR)/bin
SWAGGER_FOLDER := $(CURDIR)/swagger/swagger-ui
SWAGGER_URL := /swagger.json

GOLANGCI_LINT_VER=v1.42.1

export GOBIN=$(LOCAL_BIN)

.PHONY: .install-protoc-deps
.install-protoc-deps:
	$(info Downloading protoc plugins)
	go mod tidy
	go install \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-grpc-gateway \
    github.com/grpc-ecosystem/grpc-gateway/v2/protoc-gen-openapiv2 \
    google.golang.org/protobuf/cmd/protoc-gen-go \
	google.golang.org/grpc/cmd/protoc-gen-go-grpc \
	github.com/bufbuild/buf/cmd/buf

.PHONY: .install-golangci-lint
.install-golangci-lint:
	$(info Downloading golangci-lint)
	go install github.com/golangci/golangci-lint/cmd/golangci-lint@$(GOLANGCI_LINT_VER)

.PHONY: .goose
.goose:
	$(info Downloading goose)
	go mod tidy
	go install github.com/pressly/goose/v3/cmd/goose

.PHONY: bin-deps
bin-deps: .install-protoc-deps .goose .install-golangci-lint

.PHONY: buf-build
buf-build:
	"$(LOCAL_BIN)/buf" build

.PHONY: download-swagger
download-swagger:
	$(info Downloading swagger-ui)
	tmp=$$(mktemp -d) && \
	git clone --depth=1 https://github.com/swagger-api/swagger-ui.git $$tmp && \
	sed -i -e "s|https://petstore.swagger.io/v2/swagger.json|${SWAGGER_URL}|g" $$tmp/dist/index.html && \
	mkdir -p $(SWAGGER_FOLDER)/swagger-ui && \
	mv $$tmp/dist/* $(SWAGGER_FOLDER)/swagger-ui && \
	rm -rf $$tmp

.PHONY: generate
generate: buf-build download-swagger
	PATH=$(LOCAL_BIN):$$PATH $(LOCAL_BIN)/buf generate

.PHONY: build
build: download-swagger
	$(info Building app)
	go build -o $$GOBIN/echoapi cmd/echoapi/main.go

.PHONY: migration
migration:
	$(LOCAL_BIN)/goose -dir migrations create rename_me sql

.PHONY: migrate
migrate:
	$(LOCAL_BIN)/goose -dir migrations up

.PHONY: test
test:
	go test ./...

.PHONY: test-integration
test-integration:
	go test ./... -tags integration

.PHONY: lint
lint:
	$(LOCAL_BIN)/golangci-lint run --config .golangci.yaml
	$(LOCAL_BIN)/buf lint
