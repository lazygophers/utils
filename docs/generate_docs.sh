#!/bin/bash

# LazyGophers Utils - Documentation Generation Script
# This script automatically generates and updates all project documentation

set -e

# Colors for output
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[1;33m'
BLUE='\033[0;34m'
NC='\033[0m' # No Color

# Configuration
DOCS_DIR="docs"
PROJECT_ROOT="$(cd "$(dirname "${BASH_SOURCE[0]}")/.." && pwd)"
COVERAGE_THRESHOLD=80
BENCHMARK_TIMEOUT=300s

# Utility functions
log_info() {
    echo -e "${BLUE}[INFO]${NC} $1"
}

log_success() {
    echo -e "${GREEN}[SUCCESS]${NC} $1"
}

log_warning() {
    echo -e "${YELLOW}[WARNING]${NC} $1"
}

log_error() {
    echo -e "${RED}[ERROR]${NC} $1"
}

# Check prerequisites
check_prerequisites() {
    log_info "Checking prerequisites..."
    
    # Check Go installation
    if ! command -v go &> /dev/null; then
        log_error "Go is not installed or not in PATH"
        exit 1
    fi
    
    # Check Go version
    GO_VERSION=$(go version | grep -oE '[0-9]+\.[0-9]+\.[0-9]+')
    REQUIRED_VERSION="1.24.0"
    if [[ "$(printf '%s\n' "$REQUIRED_VERSION" "$GO_VERSION" | sort -V | head -n1)" != "$REQUIRED_VERSION" ]]; then
        log_error "Go version $REQUIRED_VERSION or higher is required, found $GO_VERSION"
        exit 1
    fi
    
    # Check if we're in the project root
    if [[ ! -f "go.mod" ]]; then
        log_error "Must be run from project root directory"
        exit 1
    fi
    
    log_success "Prerequisites check passed"
}

# Create docs directory structure
setup_docs_structure() {
    log_info "Setting up documentation structure..."
    
    mkdir -p "${DOCS_DIR}"/{api,architecture,contributing,guides,reports}
    mkdir -p "${DOCS_DIR}/assets"/{images,diagrams}
    mkdir -p "${DOCS_DIR}/examples"
    
    log_success "Documentation structure created"
}

# Generate test coverage report
generate_coverage_report() {
    log_info "Generating test coverage report..."
    
    # Run tests with coverage (excluding problematic packages)
    EXCLUDE_PACKAGES="$(go list ./... | grep -v pgp | tr '\n' ' ')"
    
    if go test -coverprofile="${DOCS_DIR}/reports/coverage.out" -covermode=atomic $EXCLUDE_PACKAGES; then
        # Generate HTML coverage report
        go tool cover -html="${DOCS_DIR}/reports/coverage.out" -o="${DOCS_DIR}/reports/coverage.html"
        
        # Generate function-level coverage summary
        go tool cover -func="${DOCS_DIR}/reports/coverage.out" > "${DOCS_DIR}/reports/coverage_summary.txt"
        
        # Extract overall coverage percentage
        COVERAGE=$(go tool cover -func="${DOCS_DIR}/reports/coverage.out" | grep total | awk '{print $3}' | sed 's/%//')
        
        if (( $(echo "$COVERAGE >= $COVERAGE_THRESHOLD" | bc -l) )); then
            log_success "Test coverage: ${COVERAGE}% (above threshold of ${COVERAGE_THRESHOLD}%)"
        else
            log_warning "Test coverage: ${COVERAGE}% (below threshold of ${COVERAGE_THRESHOLD}%)"
        fi
    else
        log_error "Test coverage generation failed"
        return 1
    fi
}

# Generate benchmark report
generate_benchmark_report() {
    log_info "Generating benchmark report..."
    
    # Run benchmarks for key packages
    BENCHMARK_PACKAGES=("anyx" "candy" "stringx" "wait" "hystrix" "cryptox")
    
    echo "# Benchmark Report - $(date)" > "${DOCS_DIR}/reports/benchmarks.txt"
    echo "Generated on: $(date)" >> "${DOCS_DIR}/reports/benchmarks.txt"
    echo "Platform: $(uname -m) $(uname -s)" >> "${DOCS_DIR}/reports/benchmarks.txt"
    echo "Go Version: $(go version)" >> "${DOCS_DIR}/reports/benchmarks.txt"
    echo "" >> "${DOCS_DIR}/reports/benchmarks.txt"
    
    for package in "${BENCHMARK_PACKAGES[@]}"; do
        if [[ -d "$package" ]]; then
            log_info "Running benchmarks for $package..."
            echo "## Package: $package" >> "${DOCS_DIR}/reports/benchmarks.txt"
            if timeout $BENCHMARK_TIMEOUT go test -bench=. -benchmem -run=^$ "./$package" >> "${DOCS_DIR}/reports/benchmarks.txt" 2>&1; then
                log_success "Benchmarks completed for $package"
            else
                log_warning "Benchmarks failed or timed out for $package"
            fi
            echo "" >> "${DOCS_DIR}/reports/benchmarks.txt"
        fi
    done
}

# Generate API documentation
generate_api_docs() {
    log_info "Generating API documentation..."
    
    # Use godoc to extract documentation
    if command -v godoc &> /dev/null; then
        godoc -templates="${DOCS_DIR}/templates" -html . > "${DOCS_DIR}/api/godoc.html" 2>/dev/null || true
    fi
    
    # Generate package list with descriptions
    echo "# Package Index" > "${DOCS_DIR}/api/packages.md"
    echo "" >> "${DOCS_DIR}/api/packages.md"
    
    for dir in */; do
        if [[ -f "${dir}go.mod" ]] || [[ -f "${dir}*.go" ]]; then
            PACKAGE_NAME=$(basename "$dir")
            if [[ "$PACKAGE_NAME" != "docs" && "$PACKAGE_NAME" != ".git" ]]; then
                echo "## $PACKAGE_NAME" >> "${DOCS_DIR}/api/packages.md"
                
                # Extract package comment if available
                if [[ -f "${dir}doc.go" ]]; then
                    PACKAGE_DESC=$(head -10 "${dir}doc.go" | grep -E "^// " | sed 's|^// ||' | head -1)
                    echo "$PACKAGE_DESC" >> "${DOCS_DIR}/api/packages.md"
                elif [[ -f "${dir}${PACKAGE_NAME}.go" ]]; then
                    PACKAGE_DESC=$(head -10 "${dir}${PACKAGE_NAME}.go" | grep -E "^// " | sed 's|^// ||' | head -1)
                    echo "$PACKAGE_DESC" >> "${DOCS_DIR}/api/packages.md"
                fi
                
                echo "" >> "${DOCS_DIR}/api/packages.md"
            fi
        fi
    done
    
    log_success "API documentation generated"
}

# Generate architecture diagrams
generate_architecture_diagrams() {
    log_info "Generating architecture diagrams..."
    
    # Create a simple dependency graph in DOT format
    cat > "${DOCS_DIR}/assets/diagrams/dependencies.dot" << 'EOF'
digraph Dependencies {
    rankdir=TB;
    node [shape=box, style=rounded];
    
    // Core packages
    utils [label="utils\n(root)", fillcolor=lightblue, style=filled];
    candy [label="candy\n(type conversion)", fillcolor=lightgreen, style=filled];
    json [label="json\n(serialization)", fillcolor=lightgreen, style=filled];
    
    // Infrastructure
    runtime [label="runtime\n(system utils)", fillcolor=lightyellow, style=filled];
    routine [label="routine\n(goroutine mgmt)", fillcolor=lightyellow, style=filled];
    app [label="app\n(build info)", fillcolor=lightyellow, style=filled];
    
    // Utilities
    stringx [label="stringx\n(string ops)", fillcolor=lightcyan, style=filled];
    anyx [label="anyx\n(map ops)", fillcolor=lightcyan, style=filled];
    wait [label="wait\n(concurrency)", fillcolor=lightcyan, style=filled];
    xtime [label="xtime\n(time ext)", fillcolor=lightcyan, style=filled];
    
    // Specialized
    cryptox [label="cryptox\n(crypto)", fillcolor=lightpink, style=filled];
    hystrix [label="hystrix\n(circuit breaker)", fillcolor=lightpink, style=filled];
    config [label="config\n(configuration)", fillcolor=lightpink, style=filled];
    network [label="network\n(networking)", fillcolor=lightpink, style=filled];
    
    // Dependencies
    anyx -> candy;
    anyx -> json;
    wait -> routine;
    wait -> runtime;
    routine -> runtime;
    runtime -> app;
    config -> json;
    config -> runtime;
    hystrix -> randx;
}
EOF

    # Generate SVG if graphviz is available
    if command -v dot &> /dev/null; then
        dot -Tsvg "${DOCS_DIR}/assets/diagrams/dependencies.dot" -o "${DOCS_DIR}/assets/diagrams/dependencies.svg"
        log_success "Architecture diagrams generated"
    else
        log_warning "Graphviz not installed, skipping diagram generation"
    fi
}

# Update README files
update_readme_files() {
    log_info "Updating README files..."
    
    # Update main README with current statistics
    PACKAGE_COUNT=$(find . -maxdepth 1 -type d -name "*" ! -name ".*" ! -name "docs" | wc -l)
    GO_FILES=$(find . -name "*.go" | wc -l)
    TOTAL_LINES=$(find . -name "*.go" -exec wc -l {} + | tail -1 | awk '{print $1}')
    
    # Create/update main README badges section
    if [[ -f "README.md" ]]; then
        # Update existing README with fresh statistics
        sed -i.bak "s/Total Packages: [0-9]*/Total Packages: $PACKAGE_COUNT/" README.md 2>/dev/null || true
        sed -i.bak "s/Go Files: [0-9]*/Go Files: $GO_FILES/" README.md 2>/dev/null || true
        sed -i.bak "s/Lines of Code: [0-9,]*/Lines of Code: $(printf "%'d" $TOTAL_LINES)/" README.md 2>/dev/null || true
        rm -f README.md.bak
    fi
    
    log_success "README files updated"
}

# Generate changelog
generate_changelog() {
    log_info "Generating changelog..."
    
    if command -v git &> /dev/null && [[ -d ".git" ]]; then
        echo "# Changelog" > "${DOCS_DIR}/CHANGELOG.md"
        echo "" >> "${DOCS_DIR}/CHANGELOG.md"
        echo "Generated on: $(date)" >> "${DOCS_DIR}/CHANGELOG.md"
        echo "" >> "${DOCS_DIR}/CHANGELOG.md"
        
        # Get recent commits
        git log --oneline --decorate --graph -n 50 >> "${DOCS_DIR}/CHANGELOG.md" 2>/dev/null || true
        
        log_success "Changelog generated"
    else
        log_warning "Git not available, skipping changelog generation"
    fi
}

# Validate generated documentation
validate_documentation() {
    log_info "Validating generated documentation..."
    
    VALIDATION_ERRORS=0
    
    # Check required files exist
    REQUIRED_FILES=(
        "${DOCS_DIR}/reports/coverage.out"
        "${DOCS_DIR}/reports/coverage.html"
        "${DOCS_DIR}/reports/benchmarks.txt"
        "${DOCS_DIR}/api/packages.md"
    )
    
    for file in "${REQUIRED_FILES[@]}"; do
        if [[ ! -f "$file" ]]; then
            log_error "Required file missing: $file"
            ((VALIDATION_ERRORS++))
        fi
    done
    
    # Check file sizes (should not be empty)
    for file in "${REQUIRED_FILES[@]}"; do
        if [[ -f "$file" && ! -s "$file" ]]; then
            log_warning "File is empty: $file"
        fi
    done
    
    if [[ $VALIDATION_ERRORS -eq 0 ]]; then
        log_success "Documentation validation passed"
    else
        log_error "Documentation validation failed with $VALIDATION_ERRORS errors"
        return 1
    fi
}

# Generate documentation index
generate_index() {
    log_info "Generating documentation index..."
    
    cat > "${DOCS_DIR}/README.md" << EOF
# LazyGophers Utils Documentation

Welcome to the comprehensive documentation for LazyGophers Utils, a high-performance Go utility library.

## 📚 Documentation Sections

### Core Documentation
- [🏗️ Architecture Guide](architecture_en.md) - System design and package overview
- [📖 API Reference](API_REFERENCE.md) - Comprehensive API documentation
- [🤝 Contributing Guide](CONTRIBUTING_en.md) - How to contribute to the project

### Reports and Analysis
- [📊 Test Coverage Report](reports/coverage.html) - Interactive coverage analysis
- [⚡ Performance Report](performance_report.md) - Benchmarks and optimization guide
- [📈 Benchmark Results](reports/benchmarks.txt) - Raw benchmark data

### Multi-Language Support
- [架构文档 (中文)](architecture_zh.md) - Chinese architecture documentation
- [架構文檔 (繁體中文)](architecture_zh-hant.md) - Traditional Chinese documentation
- [贡献指南 (中文)](CONTRIBUTING_zh.md) - Chinese contributing guide

### Development Resources
- [📦 Package Index](api/packages.md) - Overview of all packages
- [🔄 Changelog](CHANGELOG.md) - Recent changes and updates

## 🚀 Quick Start

\`\`\`bash
# Install the library
go get github.com/lazygophers/utils

# Run tests
go test ./...

# Generate fresh documentation
./docs/generate_docs.sh
\`\`\`

## 📊 Project Statistics

- **Total Packages**: $(find . -maxdepth 1 -type d ! -name ".*" ! -name "docs" | wc -l | tr -d ' ')
- **Go Files**: $(find . -name "*.go" | wc -l | tr -d ' ')
- **Lines of Code**: $(find . -name "*.go" -exec wc -l {} + | tail -1 | awk '{printf "%'\''d", $1}')
- **Test Coverage**: $(go tool cover -func="${DOCS_DIR}/reports/coverage.out" 2>/dev/null | grep total | awk '{print $3}' || echo "N/A")
- **Last Updated**: $(date)

## 🔧 Documentation Generation

This documentation is automatically generated using the \`generate_docs.sh\` script. To regenerate:

\`\`\`bash
cd docs
./generate_docs.sh
\`\`\`

The script will:
1. ✅ Generate test coverage reports
2. ⚡ Run performance benchmarks  
3. 📖 Update API documentation
4. 🏗️ Create architecture diagrams
5. 📝 Update README files
6. 🔍 Validate all documentation

## 📞 Support

For questions, issues, or contributions:
- 🐛 [Report Issues](https://github.com/lazygophers/utils/issues)
- 💬 [Discussions](https://github.com/lazygophers/utils/discussions)
- 📧 [Contact](mailto:support@lazygophers.com)
EOF

    log_success "Documentation index generated"
}

# Main execution
main() {
    log_info "Starting documentation generation for LazyGophers Utils..."
    echo "============================================================"
    
    cd "$PROJECT_ROOT"
    
    # Execute all steps
    check_prerequisites
    setup_docs_structure
    generate_coverage_report
    generate_benchmark_report
    generate_api_docs
    generate_architecture_diagrams
    update_readme_files
    generate_changelog
    generate_index
    validate_documentation
    
    echo "============================================================"
    log_success "Documentation generation completed successfully!"
    log_info "Documentation available in: ${DOCS_DIR}/"
    log_info "View coverage report: ${DOCS_DIR}/reports/coverage.html"
    log_info "View documentation index: ${DOCS_DIR}/README.md"
}

# Handle script arguments
case "${1:-}" in
    --help|-h)
        echo "LazyGophers Utils Documentation Generator"
        echo ""
        echo "Usage: $0 [options]"
        echo ""
        echo "Options:"
        echo "  --help, -h     Show this help message"
        echo "  --coverage     Generate only coverage reports"
        echo "  --benchmark    Generate only benchmark reports"
        echo "  --api          Generate only API documentation"
        echo "  --validate     Validate existing documentation"
        echo ""
        echo "Examples:"
        echo "  $0              # Generate all documentation"
        echo "  $0 --coverage   # Generate only coverage reports"
        echo "  $0 --validate   # Validate existing documentation"
        exit 0
        ;;
    --coverage)
        cd "$PROJECT_ROOT"
        check_prerequisites
        setup_docs_structure
        generate_coverage_report
        ;;
    --benchmark)
        cd "$PROJECT_ROOT"
        check_prerequisites
        setup_docs_structure
        generate_benchmark_report
        ;;
    --api)
        cd "$PROJECT_ROOT"
        check_prerequisites
        setup_docs_structure
        generate_api_docs
        ;;
    --validate)
        cd "$PROJECT_ROOT"
        validate_documentation
        ;;
    "")
        main
        ;;
    *)
        log_error "Unknown option: $1"
        log_info "Use --help for usage information"
        exit 1
        ;;
esac