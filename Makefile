SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

BASE_PATH ?= "http://localhost:8080"

## help: Print this help message
help:
	@echo
	@echo "Usage:"
	@echo
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /' | sort
	@echo
.PHONY: help

## build: Build the binary
build:
	@mkdir -p bin
	go build -o bin/rfd-fyi
.PHONY: build

## test: Run tests in colour
test:
	@go install github.com/rakyll/gotest@latest
	gotest -v -count=1
.PHONY: test

## fmt: Format code (with gofumpt)
fmt:
	@go install mvdan.cc/gofumpt@latest
	gofumpt -w .
.PHONY: fmt

## swagger: Generate swagger docs
swagger:
	@go install github.com/swaggo/swag/cmd/swag@latest
	swag init --outputTypes yaml
.PHONY: swagger

## server: Build and run server from source
server:
	@go run .
.PHONY: server

## seed: Generate several issues via the create endpoint
seed:
	BASE_PATH=$(BASE_PATH) ./scripts/generate-issues.sh
.PHONY: seed

## container: Build a container image with Docker
container:
	docker build . -t rfd-fyi
.PHONY: container

## container-run: Build and run a container with Docker
container-run: container
	@docker run \
			--network host \
			-u "$$(id -u)":"$$(id -g)"\
			-v "$$PWD":"/opt/rfd-fyi" \
			rfd-fyi
.PHONY: container-run
