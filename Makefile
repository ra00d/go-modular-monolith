# Go Modular Monolith — common development commands.
#
# Usage:
#   make <target> [NAME=value] [MODULE=value]
#
# Examples:
#   make dev
#   make migration-create NAME=add_users_table MODULE=auth
#   make migrate MODULE=auth

SHELL := /bin/sh
.DEFAULT_GOAL := help

GO             ?= go
BUN            ?= bun
DOCKER_COMPOSE ?= docker compose
MAIN           ?= ./cmd/cli
BINARY         ?= bin/app
MODULE         ?= auth
COVERAGE_FILE  ?= coverage.out

.PHONY: help run dev build clean \
	fmt fmt-check vet test test-race test-module cover \
	lint check tidy \
	migration-create migrate \
	infra-up infra-down infra-restart infra-logs infra-ps db-shell \
	web-dev web-build web-check web-lint web-format web-generate-api

##@ General

help: ## Show this help
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m [NAME=value] [MODULE=value]\n"} \
		/^[a-zA-Z0-9_.-]+:.*?##/ { printf "  \033[36m%-22s\033[0m %s\n", $$1, $$2 } \
		/^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) }' $(MAKEFILE_LIST)

##@ Development

run: ## Run the backend
	$(GO) run $(MAIN) start

dev: ## Run the backend with live reload (requires air)
	air

build: ## Build the backend binary
	@mkdir -p $(dir $(BINARY))
	$(GO) build -trimpath -o $(BINARY) $(MAIN)

clean: ## Remove build and coverage artifacts
	rm -rf $(BINARY) $(COVERAGE_FILE)

##@ Quality

fmt: ## Format Go sources
	$(GO) fmt ./...

fmt-check: ## Fail when Go sources are not formatted
	@unformatted="$$(gofmt -l $$(go list -f '{{.Dir}}' ./... | sort -u))"; \
	if [ -n "$$unformatted" ]; then \
		echo "Unformatted files (run 'make fmt'):"; \
		echo "$$unformatted"; \
		exit 1; \
	fi

vet: ## Run go vet
	$(GO) vet ./...

test: ## Run all Go tests
	$(GO) test ./...

test-race: ## Run Go tests with the race detector
	$(GO) test -race ./...

test-module: ## Run tests for one module (MODULE=auth)
	$(GO) test ./internal/modules/$(MODULE)/...

cover: ## Run tests with a coverage report
	$(GO) test -coverprofile=$(COVERAGE_FILE) ./...
	$(GO) tool cover -func=$(COVERAGE_FILE)

lint: fmt-check vet ## Run formatting and vet checks

check: lint test ## Run lint and tests

tidy: ## Tidy Go module dependencies
	$(GO) mod tidy

##@ Database

migration-create: ## Create a migration (NAME=..., MODULE=auth)
	@test -n "$(NAME)" || { \
		echo "usage: make migration-create NAME=<migration_name> [MODULE=auth]"; \
		exit 1; \
	}
	$(GO) run $(MAIN) migration:create "$(NAME)" --module "$(MODULE)"

migrate: ## Apply migrations and generate go-jet code (MODULE=auth)
	$(GO) run $(MAIN) db:migrate --module "$(MODULE)"

##@ Infrastructure

infra-up: ## Start PostgreSQL and Redis
	$(DOCKER_COMPOSE) up -d

infra-down: ## Stop PostgreSQL and Redis
	$(DOCKER_COMPOSE) down

infra-restart: infra-down infra-up ## Restart PostgreSQL and Redis

infra-logs: ## Follow infrastructure logs
	$(DOCKER_COMPOSE) logs -f

infra-ps: ## Show infrastructure container status
	$(DOCKER_COMPOSE) ps

db-shell: ## Open a psql shell on the development database
	$(DOCKER_COMPOSE) exec postgres psql -U postgres -d erp

##@ Web

web-dev: ## Run the web app
	cd apps/web && $(BUN) --bun run dev

web-build: ## Build the web app
	cd apps/web && $(BUN) --bun run build

web-check: ## Run Biome checks on the web app
	cd apps/web && $(BUN) --bun run check

web-lint: ## Lint the web app
	cd apps/web && $(BUN) --bun run lint

web-format: ## Format the web app
	cd apps/web && $(BUN) --bun run format

web-generate-api: ## Generate the OpenAPI client (backend must be running)
	cd apps/web && $(BUN) --bun run generate-api
