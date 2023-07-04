.DEFAULT_GOAL := help

help:
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'

build:
	go mod tidy

test: ## Run all tests in packages
	go test -v ./...

test-cover: ## Run tests with coverage and open http report in default browser
	mkdir -p build
	go test --coverprofile=build/cover.out ./...
	go tool cover --html=build/cover.out
