

.PHONY: build

build: ## Compile the Go code
	@echo ""
	@echo ""
	go mod download
	go get -d -v
	go build -o bin/hack main.go
	@echo "Build Complete"
	@echo ""
	@echo ""

container: ## Build the container image
	@echo ""
	@echo ""
	docker build -t krisnova/hack:latest .
	@echo ""
	@echo ""

push: ## Push to dockerhub
	@echo ""
	@echo ""
	docker push krisnova/hack:latest
	@echo ""
	@echo ""

install: ## Install hack on the local filesystem
	@echo ""
	@echo ""
	cp bin/hack /usr/bin/hack
	@echo ""
	@echo ""

clean: ## Clean the build artifacts
	@echo ""
	@echo ""
	rm -rf bin/*
	@echo ""
	@echo ""

.PHONY: help
help:  ## Show help messages for make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'