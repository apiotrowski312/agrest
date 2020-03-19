.DEFAULT_GOAL := help
.PHONY: help

RED='\033[0;31m'
NC='\033[0m' # No Color

# -----------------------------------------------------------------------------
# Tests
# -----------------------------------------------------------------------------

unit: ## Run unit test in Docker
	@echo -e ${RED}UNIT TESTS${NC}
	@docker run --rm -v $(PWD):/agrest golang:1.13 sh -c "cd /agrest/src; go mod download 2> /dev/null; go test ./ -cover"
	@echo -e ----------------------------

help: ## Show this help.
	@grep -E '^[a-zA-Z_-]+:.*?## .*$$' $(MAKEFILE_LIST) | awk 'BEGIN {FS = ":.*?## "}; {printf "\033[36m%-30s\033[0m %s\n", $$1, $$2}'
