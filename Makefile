.DEFAULT_GOAL := help

.PHONY: help
# From: http://disq.us/p/16327nq
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: build
build: ## Build the Go project.
	go build -ldflags="-s -w" -trimpath -v -o hajimari

.PHONY: run
run: build ## Run the program
	./hajimari

.PHONY: fmt
fmt: ## Format the project with gofmt
	gofmt -l -w -s .

.PHONY: lint
lint: ## Lint code with golangci-lint
	golangci-lint run

.PHONY: test
test: ## Run the tests
	go test -cover ./...
