

.PHONY: build

build: ## Compile the Go code
	@echo "Compiling..."
	@go build -o bin/hack main.go
	@echo "Compiling complete"

container: ## Build the container image
	docker build -t krisnova/hack:latest .

push: ## Push to dockerhub
	docker push krisnova/hack:latest

install: ## Install hack on the local filesystem
	cp bin/hack /usr/bin/hack

clean: ## Clean the build artifacts
	rm -rf bin/*

.PHONY: help
help:  ## Show help messages for make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'
