SHELL := /bin/bash

help: ## Display this help.
	@awk 'BEGIN {FS = ":.*##"; printf "\nUsage:\n  make \033[36m<target>\033[0m\n"} /^[a-zA-Z_0-9-]+:.*?##/ { printf "  \033[36m%-15s\033[0m %s\n", $$1, $$2 } /^##@/ { printf "\n\033[1m%s\033[0m\n", substr($$0, 5) } ' $(MAKEFILE_LIST)

##@ Development

buf-lint: ## Proto code linter
	@buf lint

buf-gen: ## Generarte protocol buffer code
	@buf generate

fmt: ## Format proto code files
	@find . -name '*.proto' | xargs clang-format -i
