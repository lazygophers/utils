# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

`lazygophers/utils` is a comprehensive Go utility library with modular design. Each package provides specific functionality for common development tasks. The library follows Go 1.24+ standards and emphasizes type safety, error handling, and performance optimization.

## Commands

### Testing
- `go test ./...` - Run all tests across the project
- `go test -v ./...` - Run tests with verbose output
- `go test ./[package]` - Run tests for a specific package (e.g., `go test ./candy`)
- `go test -run [TestName] ./[package]` - Run specific test function

### Building and Development
- `go mod tidy` - Clean up module dependencies
- `go build ./...` - Build all packages to verify compilation
- `go fmt ./...` - Format all Go code
- `go vet ./...` - Run Go vet static analysis

## Architecture and Code Organization

### Core Library Structure
- **Root package (`utils`)**: Contains fundamental utilities like error handling (`Must*` functions), database operations (`Scan`, `Value`), and validation (`Validate`)
- **Modular packages**: Each subdirectory is an independent module with specific functionality

### Key Design Patterns

1. **Generic Functions**: Extensive use of Go generics for type-safe utilities (e.g., `Must[T any](value T, err error) T`)

2. **Error Handling Philosophy**: 
   - `Must*` functions for panic-on-error scenarios
   - Consistent error logging using `github.com/lazygophers/log`
   - Functions return errors that are logged before being returned

3. **Database Integration**:
   - `Scan()` function handles JSON deserialization from database fields
   - `Value()` function handles JSON serialization for database storage
   - Automatic default value population using `github.com/mcuadros/go-defaults`

4. **Package Independence**: Each module is self-contained with its own README and can be imported separately

### Important Modules

- **`app`**: Application lifecycle and environment management
- **`candy`**: Type conversion utilities with comprehensive support
- **`config`**: Configuration management with multiple format support
- **`cryptox`**: Cryptographic operations and utilities
- **`event`**: Event-driven programming support
- **`hystrix`**: Circuit breaker pattern implementation
- **`routine`**: Goroutine management and concurrency utilities
- **`wait`**: Synchronization and timeout control utilities
- **`xtime`**: Enhanced time operations with specialized subpackages (007, 955, 996)

### Dependencies
- Uses `github.com/lazygophers/log` for consistent logging across all modules
- Leverages `github.com/go-playground/validator/v10` for struct validation
- Includes `github.com/bytedance/sonic` for high-performance JSON operations
- Custom JSON package wrapper for enhanced serialization

### Testing Conventions
- Test files follow `*_test.go` naming convention
- Tests are co-located with source code in each package
- Uses standard Go testing package with additional test utilities from `github.com/stretchr/testify`

### Code Style
- All public functions include comprehensive documentation comments in Chinese
- Error handling follows the pattern: log error, then return it
- Generic functions use descriptive type constraints
- Consistent use of receiver patterns for methods

## Development Notes

When working with this codebase:
- Each module has its own README with specific usage examples
- Always run `go mod tidy` after adding dependencies
- Test changes across all relevant modules as they may have interdependencies
- Follow the established error handling patterns with logging
- Use the existing JSON package wrapper instead of standard library for consistency