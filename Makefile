LOCAL_BIN := $(CURDIR)/bin
THIRD_PARTY_FOLDER := $(CURDIR)/third_party
SWAGGER_URL := /swagger.json

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

.PHONY: bin-deps
bin-deps: .install-protoc-deps

.PHONY: buf-build
buf-build:
	"$(LOCAL_BIN)/buf" build

.PHONY: download-swagger
download-swagger:
	$(info Downloading swagger-ui)
	tmp=$$(mktemp -d) && \
		git clone --depth=1 https://github.com/swagger-api/swagger-ui.git $$tmp && \
	 	sed -i '' "s|https://petstore.swagger.io/v2/swagger.json|${SWAGGER_URL}|g" $$tmp/dist/index.html && \
		mkdir -p $(THIRD_PARTY_FOLDER)/swagger-ui && \
	 	mv $$tmp/dist/* $(THIRD_PARTY_FOLDER)/swagger-ui && \
	 	rm -rf $$tmp

.PHONY: generate
generate: buf-build download-swagger
	"$(LOCAL_BIN)/buf" generate

.PHONY: build
build: generate
	go build -o $$GOBIN/sinkapi
