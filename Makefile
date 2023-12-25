.PHONY: all

.DEFAULT_GOAL := help

PROJECT_ROOT:=$(shell git rev-parse --show-toplevel)
COMMIT := $(shell git rev-parse --short HEAD)
BRANCH := $(shell git rev-parse --abbrev-ref HEAD)
REPO := $(shell basename `git rev-parse --show-toplevel`)
DATE := $(shell date +%Y-%m-%d-%H-%M-%S)
APP_NAME := $(shell basename `git rev-parse --show-toplevel`)

MAKE_LIB:=$(PROJECT_ROOT)/scripts
-include $(MAKE_LIB)/lints.mk
-include $(MAKE_LIB)/tools.mk

test: lint scan ## Run unit tests
	go test --short -cover -failfast ./...

scan: ## run security scan
	gosec ./...
	govulncheck ./...

lint: ## run linter
	go vet ./...
	golangci-lint run ./...

docgen: test ## generate go doc and append to readme.
	gomarkdoc -o README.md -e .

install: get-tools ## install tools + dependencies

# HELP
# This will output the help for each task
# thanks to https://marmelab.com/blog/2016/02/29/auto-documented-makefile.html
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

