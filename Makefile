GOPATH:=$(shell go env GOPATH)
VERSION=$(shell git describe --tags --always)
API_PROTO_FILES=$(shell find api -name *.proto)

.PHONY: init
# init env
init:
	go mod tidy

.PHONY: api
api:
	protoc --proto_path=. \
	       --proto_path=./third_party \
 	       --go_out=paths=source_relative:. \
 	       --go-grpc_out=paths=source_relative:. \
	       $(API_PROTO_FILES)

.PHONY: build
# build
build:
	mkdir -p bin/ && go build -ldflags "-X main.Version=$(VERSION)" -o ./bin/ ./...

# .PHONY: wire
# # generate wire
# wire:
# 	cd cmd/server && wire

.PHONY: all
# generate all
all:
	make init;
	make api;
# 	make wire;
	make build;

# show help
help:
	@echo 'Usage:'
	@echo ' make [target]'
	@echo ''
	@echo 'Targets:'
	@awk '/^[a-zA-Z\-\_0-9]+:/ { \
	helpMessage = match(lastLine, /^# (.*)/); \
		if (helpMessage) { \
			helpCommand = substr($$1, 0, index($$1, ":")-1); \
			helpMessage = substr(lastLine, RSTART + 2, RLENGTH); \
			printf "\033[36m%-22s\033[0m %s\n", helpCommand,helpMessage; \
		} \
	} \
	{ lastLine = $$0 }' $(MAKEFILE_LIST)

.DEFAULT_GOAL := help
