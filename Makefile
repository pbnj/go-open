SHELL := /bin/bash

.PHONY: all
all: fmt lint vet ## Run all rules

.PHONY: fmt
fmt: ## Format
	gofmt -w .

.PHONY: lint
lint: ## Lint
	golint

.PHONY: vet
vet: ## Vet
	go vet

.PHONY: test
test: ## Test
	go test ./...
