# Mercado Libre Go SDK - Makefile

.PHONY: test test-verbose test-integration test-unit test-coverage clean build help

# Default target
help:
	@echo "🚀 Mercado Libre Go SDK - Available Commands:"
	@echo "============================================="
	@echo "  test           - Run all tests with custom runner"
	@echo "  test-verbose   - Run all tests with verbose output"  
	@echo "  test-unit      - Run only unit tests (fast)"
	@echo "  test-integration - Run integration tests"
	@echo "  test-coverage  - Run tests with coverage report"
	@echo "  clean          - Clean test cache and temp files"
	@echo "  build          - Build the SDK"
	@echo "  help           - Show this help message"

# Run all tests with custom runner (recommended)
test:
	@go run test_runner.go

# Verbose native Go testing
test-verbose:
	@go test -v ./...

# Unit tests only (excluding integration)
test-unit:
	@go test -short ./...

# Integration tests only
test-integration:
	@go test -run TestIntegration ./test

# Coverage report
test-coverage:
	@go test -coverprofile=coverage.out ./...
	@go tool cover -html=coverage.out -o coverage.html
	@echo "📊 Coverage report generated: coverage.html"

# Clean test artifacts
clean:
	@go clean -testcache
	@rm -f coverage.out coverage.html
	@echo "🧹 Cleaned test artifacts"

# Build validation
build:
	@go build ./...
	@echo "✅ Build successful"

# Quick validation (build + unit tests)
validate: build test-unit
	@echo "🎉 Quick validation completed"