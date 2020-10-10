

.PHONY: build

build:
	@echo ""
	@echo ""
	go build -o bin/hack main.go
	@echo "Build Complete"
	@echo ""
	@echo ""

container:
	@echo ""
	@echo ""
	docker build -t krisnova/hack:latest image/
	@echo ""
	@echo ""


.PHONY: help
help:  ## Show help messages for make targets
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | sort | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[32m%-30s\033[0m %s\n", $$1, $$2}'