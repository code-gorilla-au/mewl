.PHONY: all

.DEFAULT_GOAL := help

COMMIT := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
REPO := $(shell basename `git rev-parse --show-toplevel`)
DATE := $(shell date +%Y-%m-%d-%H-%M-%S)
APP_NAME := pyrotic

test: ## Run unit tests
	go test --short -cover -failfast ./...

scan: ## run security scan
	gosec ./...
	go vet ./...