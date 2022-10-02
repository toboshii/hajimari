.DEFAULT_GOAL := help

.PHONY: help
# From: http://disq.us/p/16327nq
help: ## This help.
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}' $(MAKEFILE_LIST)

.PHONY: deps
deps: ## Install required dependencies
	cd frontend/ && npm install

.PHONY: build
build: ## Build the project
	npm run build --prefix frontend/
	go build -ldflags="-s -w" -trimpath -o ./bin/hajimari ./cmd/hajimari/main.go

.PHONY: run
run: build ## Run the program
	./bin/hajimari

.PHONY: dev
dev: ## Run the frontend & backend in dev mode
	make -j 2 dev-backend dev-frontend

.PHONY: dev-backend
dev-backend: ## Run the backend in dev mode
	air &

.PHONY: dev-frontend
dev-frontend: ## Run the frontend in dev mode
	cd frontend/ && sleep 3 && npm run dev -- --open

.PHONY: fmt
fmt: ## Format the project with gofmt
	gofmt -l -w -s .

.PHONY: lint
lint: ## Lint code with golangci-lint
	golangci-lint run

.PHONY: test
test: ## Run the tests
	go test -cover ./...