GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
API_PROTO_FILES=$(shell find api -name *.proto)
BUF_INSTALLED := $(shell command -v buf 2> /dev/null)

ifeq ($(GOHOSTOS), windows)
	#the `find.exe` is different from `find` in bash/shell.
	#to see https://docs.microsoft.com/en-us/windows-server/administration/windows-commands/find.
	#changed to use git-bash.exe to run find cli or other cli friendly, caused of every developer has a Git.
	Git_Bash= $(subst cmd\,bin\bash.exe,$(dir $(shell where git)))
	INTERNAL_PROTO_FILES=$(shell $(Git_Bash) -c "find internal -name *.proto")
	API_PROTO_FILES=$(shell $(Git_Bash) -c "find api -name *.proto")
else
	INTERNAL_PROTO_FILES=$(shell find internal -name *.proto)
	API_PROTO_FILES=$(shell find api -name *.proto)
endif

.PHONY: errors
# generate errors code
errors:
	protoc --proto_path=. \
		   --proto_path=./third_party \
		   --go_out=paths=source_relative:. \
		   --go-errors_out=paths=source_relative:. \
		   $(API_PROTO_FILES)

.PHONY: validate
# generate validate code
validate:
	protoc --proto_path=. \
		   --proto_path=./third_party \
		   --go_out=paths=source_relative:. \
		   --validate_out=paths=source_relative,lang=go:. \
		   $(API_PROTO_FILES)

.PHONY: config
# generate internal proto
config:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
	       $(INTERNAL_PROTO_FILES)

.PHONY: api
# generate api proto
api:
	protoc --proto_path=. \
		   --proto_path=./api \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-http_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && GOPROXY="https://goproxy.cn,direct" go build -ldflags '-w -s -extldflags "-static" -X main.Version=$(VERSION)' -tags musl -o ./bin/ ./...

.PHONY: generate
# generate
generate:
	go generate ./...

.PHONY: lint
# golangci-lint
lint:
	go mod tidy
	GOGC=20 golangci-lint run ./...

.PHONY: swagger
# proto -> swagger.json
swagger:
	protoc --proto_path=. \
        --proto_path=./third_party \
        --openapiv2_out . \
        --openapiv2_opt logtostderr=true \
        --openapiv2_opt json_names_for_fields=false \
        $(API_PROTO_FILES)

.PHONY: wire
# wire
wire:
	wire ./...

.PHONY: all
# generate all
all:
	make api;
	make errors;
	make validate;
	make config;
	go mod tidy;
	make swagger;
	make wire;
	make build;
