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

## backend: Build and run the backend from source
backend:
	@cd backend && go run .
.PHONY: backend

## frontend: Build and run the frontend from source
frontend:
	@npx vue-cli-service serve
.PHONY: server

## dev: Build and run in docker compose
dev:
	docker compose up --build -d
.PHONY: up

## prod: Run the latest images in docker compose
prod:
	docker compose -f docker-compose.prod.yml up -d
.PHONY: up

## teardown: Teardown docker
teardown:
	docker compose down
.PHONY: teardown
