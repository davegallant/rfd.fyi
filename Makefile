SHELL := bash
.SHELLFLAGS := -eu -o pipefail -c
.DELETE_ON_ERROR:
MAKEFLAGS += --warn-undefined-variables
MAKEFLAGS += --no-builtin-rules

## help: Print this help message
help:
	@echo
	@echo "Usage:"
	@echo
	@sed -n 's/^##//p' ${MAKEFILE_LIST} | column -t -s ':' |  sed -e 's/^/ /' | sort
	@echo
.PHONY: help

## backend-server: Build and run the backend from source
backend-server:
	@cd backend && go run .
.PHONY: server


## frontend-server: Build and run the frontend from source
frontend-server:
	@npx vue-cli-service serve
.PHONY: server
