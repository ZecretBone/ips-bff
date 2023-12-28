GO ?= go
GO_TAGS ?= go_json
GO_FILES = $(shell find . -name \*.go)
GO_BUILDFLAG ?= -v trimpath$(if $(GO_TAGS), -tags $(GO_TAGS),)
GO_TESTFLAGS ?= -tags "test$(if $)"
GOTESTFLAGS ?= -tags "test$(if $(GO_TAGS),$(comma)$(GO_TAGS),)" -short -race

all: generate build

generate: pregenerate
	@printf \\e[1m"Generate"\\e[0m\\n
	@cd proto && $(GO) generate

pregenerate:
	@printf \\e[1m"Install dependency"\\e[0m\\n
	@$(GO) install github.com/golang/mock/mockgen@v1.6.0

test:
	@printf \\e[1m"Run test"\\e[0m\\n
	@ENV=unittest $(GO) test $(GOTESTFLAGS) ./...

build: .bin/bff-api .bin/bff-grpc

go.sum: go.mod
	@printf \\e[1m"go mod tidy"\\e[0m\\n
	@$(GO) mod tidy

.bin/bff-api: go.mod go.sum $(GO_FILES)
	@printf \\e[1m"Build .bin/bff-api"\\e[0m\\n
	@cd cmd/bff-api && $(GO) build -o ../../.bin/bff-api .

.bin/bff-grpc: go.mod go.sum $(GO_FILES)
	@printf \\e[1m"Build .bin/bff-grpc"\\e[0m\\n
	@cd cmd/bff-grpc && $(GO) build -o ../../.bin/bff-grpc .