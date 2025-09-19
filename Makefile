# Makefile for LazyGophers Utils
# https://makefiletutorial.com/

# Go parameters
GOCMD=go
GOBUILD=$(GOCMD) build
GOCLEAN=$(GOCMD) clean
GOTEST=$(GOCMD) test
GOGET=$(GOCMD) get
GOMOD=$(GOCMD) mod
GOFMT=gofmt
GOLINT=golangci-lint

# Project info
PROJECT_NAME=lazygophers-utils
VERSION ?= $(shell git describe --tags --always --dirty)
COMMIT ?= $(shell git rev-parse --short HEAD)
BUILD_TIME ?= $(shell date -u +"%Y-%m-%dT%H:%M:%SZ")

# Directories
BUILD_DIR=build
DIST_DIR=dist
DOCS_DIR=docs
COVERAGE_DIR=$(DOCS_DIR)/reports

# Coverage settings
COVERAGE_FILE=$(COVERAGE_DIR)/coverage.out
COVERAGE_HTML=$(COVERAGE_DIR)/coverage.html
COVERAGE_THRESHOLD=70

# Lint settings
LINT_TIMEOUT=10m

# Test settings
TEST_TIMEOUT=5m
TEST_PACKAGES=$(shell $(GOCMD) list ./... | grep -v pgp)

# Colors for output
RED=\033[0;31m
GREEN=\033[0;32m
YELLOW=\033[1;33m
BLUE=\033[0;34m
NC=\033[0m # No Color

.PHONY: all build clean test test-coverage test-verbose test-race lint lint-fix fmt fmt-check mod-tidy mod-verify mod-download help docs docs-serve benchmark security deps update-deps install-tools check prepare release dev

# Default target
all: clean fmt lint test build ## Run all main targets (clean, fmt, lint, test, build)

# Help target
help: ## Display this help message
	@echo "$(GREEN)LazyGophers Utils - Available Commands$(NC)"
	@echo ""
	@awk 'BEGIN {FS = ":.*?## "} /^[a-zA-Z_-]+:.*?## / {printf "$(BLUE)%-20s$(NC) %s\n", $$1, $$2}' $(MAKEFILE_LIST)
	@echo ""
	@echo "$(YELLOW)Examples:$(NC)"
	@echo "  make test              # Run tests"
	@echo "  make test-coverage     # Run tests with coverage"
	@echo "  make lint              # Run linter"
	@echo "  make lint-fix          # Run linter with auto-fix"
	@echo "  make docs              # Generate documentation"
	@echo "  make dev               # Development setup"

# Build targets
build: ## Build the project
	@echo "$(GREEN)Building project...$(NC)"
	$(GOBUILD) -v ./...
	@echo "$(GREEN)✅ Build completed$(NC)"

clean: ## Clean build artifacts
	@echo "$(YELLOW)Cleaning build artifacts...$(NC)"
	$(GOCLEAN)
	rm -rf $(BUILD_DIR) $(DIST_DIR)
	find . -name "*.bak*" -type f -delete
	find . -name "*.test" -type f -delete
	find . -name "cover*.html" -type f -delete
	find . -name "*.out" -type f -delete
	find . -name "*.prof" -type f -delete
	@echo "$(GREEN)✅ Clean completed$(NC)"

# Test targets
test: ## Run tests
	@echo "$(GREEN)Running tests...$(NC)"
	$(GOTEST) -v -timeout $(TEST_TIMEOUT) $(TEST_PACKAGES)
	@echo "$(GREEN)✅ Tests completed$(NC)"

test-verbose: ## Run tests with verbose output
	@echo "$(GREEN)Running tests with verbose output...$(NC)"
	$(GOTEST) -v -timeout $(TEST_TIMEOUT) $(TEST_PACKAGES)

test-race: ## Run tests with race detection
	@echo "$(GREEN)Running tests with race detection...$(NC)"
	$(GOTEST) -race -timeout $(TEST_TIMEOUT) $(TEST_PACKAGES)
	@echo "$(GREEN)✅ Race tests completed$(NC)"

test-coverage: ## Run tests with coverage
	@echo "$(GREEN)Running tests with coverage...$(NC)"
	@mkdir -p $(COVERAGE_DIR)
	$(GOTEST) -race -coverprofile=$(COVERAGE_FILE) -covermode=atomic -timeout $(TEST_TIMEOUT) $(TEST_PACKAGES)
	@$(GOCMD) tool cover -html=$(COVERAGE_FILE) -o $(COVERAGE_HTML)
	@$(GOCMD) tool cover -func=$(COVERAGE_FILE) | tail -1
	@COVERAGE=$$($(GOCMD) tool cover -func=$(COVERAGE_FILE) | grep total | awk '{print $$3}' | sed 's/%//'); \
	echo "Coverage: $$COVERAGE%"; \
	if [ $$(echo "$$COVERAGE < $(COVERAGE_THRESHOLD)" | bc -l) -eq 1 ]; then \
		echo "$(RED)❌ Coverage $$COVERAGE% is below threshold $(COVERAGE_THRESHOLD)%$(NC)"; \
		exit 1; \
	else \
		echo "$(GREEN)✅ Coverage $$COVERAGE% meets threshold$(NC)"; \
	fi

test-short: ## Run tests with short flag
	@echo "$(GREEN)Running short tests...$(NC)"
	$(GOTEST) -short -timeout $(TEST_TIMEOUT) $(TEST_PACKAGES)

# Benchmark targets
benchmark: ## Run benchmarks
	@echo "$(GREEN)Running benchmarks...$(NC)"
	@mkdir -p $(COVERAGE_DIR)
	$(GOTEST) -bench=. -benchmem -run=^$$ $(TEST_PACKAGES) | tee $(COVERAGE_DIR)/benchmarks.txt
	@echo "$(GREEN)✅ Benchmarks completed$(NC)"

benchmark-cpu: ## Run CPU benchmarks
	@echo "$(GREEN)Running CPU benchmarks...$(NC)"
	$(GOTEST) -bench=. -benchtime=10s -cpuprofile=cpu.prof $(TEST_PACKAGES)

benchmark-mem: ## Run memory benchmarks
	@echo "$(GREEN)Running memory benchmarks...$(NC)"
	$(GOTEST) -bench=. -benchtime=10s -memprofile=mem.prof $(TEST_PACKAGES)

# Lint targets
lint: ## Run golangci-lint
	@echo "$(GREEN)Running golangci-lint...$(NC)"
	$(GOLINT) run --timeout $(LINT_TIMEOUT)
	@echo "$(GREEN)✅ Lint completed$(NC)"

lint-fix: ## Run golangci-lint with auto-fix
	@echo "$(GREEN)Running golangci-lint with auto-fix...$(NC)"
	$(GOLINT) run --timeout $(LINT_TIMEOUT) --fix
	@echo "$(GREEN)✅ Lint with fix completed$(NC)"

lint-verbose: ## Run golangci-lint with verbose output
	@echo "$(GREEN)Running golangci-lint with verbose output...$(NC)"
	$(GOLINT) run --timeout $(LINT_TIMEOUT) -v

lint-config: ## Show golangci-lint configuration
	$(GOLINT) config path
	$(GOLINT) linters

# Format targets
fmt: ## Format Go code
	@echo "$(GREEN)Formatting Go code...$(NC)"
	$(GOFMT) -s -w .
	@echo "$(GREEN)✅ Format completed$(NC)"

fmt-check: ## Check if Go code is formatted
	@echo "$(GREEN)Checking Go code format...$(NC)"
	@FILES=$$($(GOFMT) -l .); \
	if [ -n "$$FILES" ]; then \
		echo "$(RED)❌ The following files are not formatted:$(NC)"; \
		echo "$$FILES"; \
		exit 1; \
	else \
		echo "$(GREEN)✅ All files are properly formatted$(NC)"; \
	fi

# Module targets
mod-tidy: ## Tidy Go modules
	@echo "$(GREEN)Tidying Go modules...$(NC)"
	$(GOMOD) tidy
	@echo "$(GREEN)✅ Module tidy completed$(NC)"

mod-verify: ## Verify Go modules
	@echo "$(GREEN)Verifying Go modules...$(NC)"
	$(GOMOD) verify
	@echo "$(GREEN)✅ Module verify completed$(NC)"

mod-download: ## Download Go modules
	@echo "$(GREEN)Downloading Go modules...$(NC)"
	$(GOMOD) download
	@echo "$(GREEN)✅ Module download completed$(NC)"

mod-update: ## Update Go modules
	@echo "$(GREEN)Updating Go modules...$(NC)"
	$(GOMOD) get -u ./...
	$(GOMOD) tidy
	@echo "$(GREEN)✅ Module update completed$(NC)"

# Documentation targets
docs: ## Generate documentation
	@echo "$(GREEN)Generating documentation...$(NC)"
	@chmod +x $(DOCS_DIR)/generate_docs.sh
	@./$(DOCS_DIR)/generate_docs.sh
	@echo "$(GREEN)✅ Documentation generated$(NC)"

docs-serve: ## Serve documentation locally
	@echo "$(GREEN)Serving documentation on http://localhost:8080$(NC)"
	@cd $(DOCS_DIR) && python3 -m http.server 8080 2>/dev/null || python -m SimpleHTTPServer 8080

docs-validate: ## Validate documentation
	@echo "$(GREEN)Validating documentation...$(NC)"
	@./$(DOCS_DIR)/generate_docs.sh --validate
	@echo "$(GREEN)✅ Documentation validation completed$(NC)"

# Security targets
security: ## Run security checks
	@echo "$(GREEN)Running security checks...$(NC)"
	@if command -v gosec >/dev/null 2>&1; then \
		gosec -fmt=text ./...; \
	else \
		echo "$(YELLOW)⚠️  gosec not installed, skipping security check$(NC)"; \
		echo "Install with: go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest"; \
	fi

# Dependency targets
deps: ## Download and verify dependencies
	@echo "$(GREEN)Managing dependencies...$(NC)"
	$(MAKE) mod-download
	$(MAKE) mod-verify
	$(MAKE) mod-tidy

update-deps: ## Update all dependencies
	@echo "$(GREEN)Updating dependencies...$(NC)"
	$(GOGET) -u ./...
	$(MAKE) mod-tidy

# Tool installation
install-tools: ## Install development tools
	@echo "$(GREEN)Installing development tools...$(NC)"
	@echo "Installing golangci-lint..."
	@curl -sSfL https://raw.githubusercontent.com/golangci/golangci-lint/master/install.sh | sh -s -- -b $$(go env GOPATH)/bin
	@echo "Installing gosec..."
	@go install github.com/securecodewarrior/gosec/v2/cmd/gosec@latest
	@echo "Installing other tools..."
	@go install golang.org/x/tools/cmd/godoc@latest
	@go install golang.org/x/tools/cmd/goimports@latest
	@echo "$(GREEN)✅ Tools installation completed$(NC)"

# Development workflow
dev: ## Set up development environment
	@echo "$(GREEN)Setting up development environment...$(NC)"
	$(MAKE) deps
	$(MAKE) install-tools
	$(MAKE) fmt
	$(MAKE) lint
	$(MAKE) test
	@echo "$(GREEN)✅ Development environment ready$(NC)"

# Check targets
check: ## Run all checks (fmt, lint, test, security)
	@echo "$(GREEN)Running all checks...$(NC)"
	$(MAKE) fmt-check
	$(MAKE) lint
	$(MAKE) test-race
	$(MAKE) security
	@echo "$(GREEN)✅ All checks passed$(NC)"

prepare: ## Prepare for commit (fmt, lint, test)
	@echo "$(GREEN)Preparing for commit...$(NC)"
	$(MAKE) fmt
	$(MAKE) mod-tidy
	$(MAKE) lint
	$(MAKE) test-coverage
	@echo "$(GREEN)✅ Ready for commit$(NC)"


# Build for multiple platforms
build-all: ## Build for all platforms
	@echo "$(GREEN)Building for all platforms...$(NC)"
	@mkdir -p $(DIST_DIR)
	@for GOOS in linux darwin windows; do \
		for GOARCH in amd64 arm64; do \
			echo "Building for $$GOOS/$$GOARCH..."; \
			if [ "$$GOOS" = "windows" ]; then \
				BINARY_NAME="$(PROJECT_NAME)-$$GOOS-$$GOARCH.exe"; \
			else \
				BINARY_NAME="$(PROJECT_NAME)-$$GOOS-$$GOARCH"; \
			fi; \
			env GOOS=$$GOOS GOARCH=$$GOARCH $(GOBUILD) -ldflags="-s -w -X main.version=$(VERSION) -X main.commit=$(COMMIT) -X main.buildTime=$(BUILD_TIME)" -o "$(DIST_DIR)/$$BINARY_NAME" ./cmd/...; \
		done; \
	done
	@echo "$(GREEN)✅ Multi-platform build completed$(NC)"

# CI targets
ci: ## Run CI pipeline locally
	@echo "$(GREEN)Running CI pipeline...$(NC)"
	$(MAKE) clean
	$(MAKE) deps
	$(MAKE) fmt-check
	$(MAKE) lint
	$(MAKE) test-coverage
	$(MAKE) security
	$(MAKE) build
	@echo "$(GREEN)✅ CI pipeline completed$(NC)"

# Quick targets for common workflows
quick-test: mod-tidy fmt lint test ## Quick development test cycle

quick-check: fmt-check lint test-short ## Quick check before commit

# Show project information
info: ## Show project information
	@echo "$(BLUE)Project Information$(NC)"
	@echo "Name: $(PROJECT_NAME)"
	@echo "Version: $(VERSION)"
	@echo "Commit: $(COMMIT)"
	@echo "Build Time: $(BUILD_TIME)"
	@echo "Go Version: $$($(GOCMD) version)"
	@echo "GOPATH: $$(go env GOPATH)"
	@echo "GOROOT: $$(go env GOROOT)"

# Clean everything including cache
clean-all: clean ## Clean everything including Go cache
	@echo "$(YELLOW)Cleaning Go cache...$(NC)"
	$(GOCMD) clean -cache -testcache -modcache
	@echo "$(GREEN)✅ Full clean completed$(NC)"

# CI/CD targets
validate-workflows: ## Validate GitHub Actions workflows
	@echo "$(GREEN)Validating GitHub Actions workflows...$(NC)"
	@./scripts/validate-workflows.sh

test-coverage-local: ## Run test coverage locally and generate report
	@echo "$(GREEN)Running local test coverage...$(NC)"
	@./scripts/test-coverage.sh

coverage-badge: test-coverage ## Generate coverage badge information
	@echo "$(GREEN)Generating coverage badge...$(NC)"
	@COVERAGE=$$(go tool cover -func=coverage.out | grep total | awk '{print substr($$3, 1, length($$3)-1)}'); \
	if awk "BEGIN {exit !($$COVERAGE >= 90)}"; then \
		COLOR="brightgreen"; \
	elif awk "BEGIN {exit !($$COVERAGE >= 80)}"; then \
		COLOR="green"; \
	elif awk "BEGIN {exit !($$COVERAGE >= 70)}"; then \
		COLOR="yellow"; \
	elif awk "BEGIN {exit !($$COVERAGE >= 60)}"; then \
		COLOR="orange"; \
	else \
		COLOR="red"; \
	fi; \
	BADGE_URL="https://img.shields.io/badge/coverage-$${COVERAGE}%25-$${COLOR}"; \
	echo "Coverage: $${COVERAGE}%"; \
	echo "Badge URL: $${BADGE_URL}"; \
	echo "$${BADGE_URL}" > coverage-badge-url.txt

pre-commit: ## Run all pre-commit checks
	@echo "$(GREEN)Running pre-commit checks...$(NC)"
	$(MAKE) fmt
	$(MAKE) mod-tidy
	$(MAKE) lint
	$(MAKE) test-coverage
	$(MAKE) validate-workflows
	@echo "$(GREEN)✅ All pre-commit checks passed$(NC)"

ci-local: ## Simulate GitHub Actions CI locally
	@echo "$(GREEN)Simulating CI pipeline locally...$(NC)"
	$(MAKE) clean
	$(MAKE) mod-download
	$(MAKE) fmt-check
	$(MAKE) lint
	$(MAKE) test-coverage
	$(MAKE) security
	$(MAKE) build
	$(MAKE) validate-workflows
	@echo "$(GREEN)✅ Local CI simulation completed$(NC)"
