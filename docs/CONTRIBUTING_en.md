# Contributing to LazyGophers Utils

Thank you for your interest in contributing to LazyGophers Utils! This document provides guidelines and information for contributors.

## 🚀 Getting Started

### Prerequisites

- Go 1.24.0 or later
- Git
- Make (optional, for automation)

### Development Setup

1. **Fork and Clone**
   ```bash
   git clone https://github.com/your-username/utils.git
   cd utils
   ```

2. **Install Dependencies**
   ```bash
   go mod tidy
   ```

3. **Verify Setup**
   ```bash
   go test ./...
   ```

## 📋 Development Guidelines

### Code Style

1. **Follow Go Standards**
   - Use `gofmt` for formatting
   - Follow effective Go practices
   - Use meaningful variable and function names

2. **Package-Specific Guidelines**
   - Each package should be independent and reusable
   - Minimize external dependencies
   - Use generics for type safety where appropriate

3. **Documentation**
   - All public functions must have documentation comments in Chinese
   - Include usage examples for complex functions
   - Document performance characteristics for critical functions

4. **Error Handling**
   - Follow the library's error handling pattern: log then return
   - Use `github.com/lazygophers/log` for consistent logging
   - Provide meaningful error messages

### Performance Guidelines

1. **Memory Optimization**
   - Use object pools for high-frequency operations
   - Prefer zero-copy operations where possible
   - Minimize memory allocations in hot paths

2. **Concurrency**
   - Use atomic operations instead of mutexes when possible
   - Ensure thread safety for concurrent operations
   - Design for lock-free algorithms where appropriate

3. **Benchmarking**
   - Add benchmarks for performance-critical functions
   - Include memory allocation metrics (`-benchmem`)
   - Compare against baseline implementations

## 🧪 Testing Requirements

### Unit Testing

1. **Test Coverage**
   - Aim for 90%+ test coverage for new code
   - Test both success and error paths
   - Include edge cases and boundary conditions

2. **Test Organization**
   ```bash
   # Run tests for specific package
   go test ./candy
   
   # Run with coverage
   go test -cover ./...
   
   # Generate coverage report
   go test -coverprofile=coverage.out ./...
   go tool cover -html=coverage.out
   ```

3. **Test Naming**
   - Use descriptive test names: `TestFunctionName_Condition_ExpectedResult`
   - Group related tests using subtests
   - Use table-driven tests for multiple scenarios

### Benchmark Testing

1. **Performance Tests**
   ```bash
   # Run benchmarks
   go test -bench=. -benchmem ./...
   
   # Specific package benchmarks
   go test -bench=BenchmarkFunctionName ./package
   ```

2. **Benchmark Guidelines**
   - Include memory allocation metrics
   - Test with realistic data sizes
   - Compare against existing implementations

### Integration Testing

1. **Cross-Package Testing**
   - Test interactions between packages
   - Verify compatibility with different Go versions
   - Test concurrent usage patterns

## 📝 Commit Guidelines

### Commit Message Format

```
<type>(<scope>): <description>

<body>

<footer>
```

**Types:**
- `feat`: New feature
- `fix`: Bug fix
- `perf`: Performance improvement
- `refactor`: Code refactoring
- `test`: Adding or updating tests
- `docs`: Documentation changes
- `style`: Code style changes
- `ci`: CI/CD changes

**Examples:**
```bash
feat(candy): add generic slice transformation utilities

- Implement Map, Filter, and Reduce functions using Go generics
- Add comprehensive test coverage for all edge cases
- Include benchmarks showing 30% performance improvement

Closes #123
```

### Branch Naming

- Feature branches: `feature/description`
- Bug fixes: `fix/description`
- Performance improvements: `perf/description`
- Documentation: `docs/description`

## 🔍 Code Review Process

### Pull Request Guidelines

1. **Before Submitting**
   - Ensure all tests pass
   - Run `go fmt ./...`
   - Run `go vet ./...`
   - Update documentation if needed
   - Add or update tests for new functionality

2. **PR Description Template**
   ```markdown
   ## Description
   Brief description of changes

   ## Type of Change
   - [ ] Bug fix
   - [ ] New feature
   - [ ] Performance improvement
   - [ ] Breaking change
   - [ ] Documentation update

   ## Testing
   - [ ] Unit tests pass
   - [ ] Integration tests pass
   - [ ] Benchmarks included/updated
   - [ ] Manual testing completed

   ## Checklist
   - [ ] Code follows style guidelines
   - [ ] Self-review completed
   - [ ] Documentation updated
   - [ ] No breaking changes (or documented)
   ```

3. **Review Criteria**
   - Code quality and readability
   - Test coverage and quality
   - Performance impact
   - Breaking changes
   - Documentation completeness

## 🏗️ Architecture Guidelines

### Package Design

1. **Single Responsibility**
   - Each package should have a clear, focused purpose
   - Avoid mixing unrelated functionality
   - Keep public APIs minimal and clean

2. **Dependencies**
   - Minimize external dependencies
   - Prefer standard library when possible
   - Document dependency rationale

3. **Backwards Compatibility**
   - Avoid breaking changes in minor releases
   - Deprecate functions before removal
   - Provide migration guides for breaking changes

### Performance Considerations

1. **Memory Management**
   - Use sync.Pool for temporary objects
   - Implement proper cleanup in long-running operations
   - Monitor memory usage in benchmarks

2. **CPU Optimization**
   - Profile CPU-intensive operations
   - Use appropriate data structures
   - Consider cache locality

3. **Concurrency**
   - Design for concurrent usage
   - Use channels and goroutines effectively
   - Avoid race conditions

## 📚 Documentation

### API Documentation

1. **Function Documentation**
   ```go
   // ToString 将任意类型转换为字符串
   // 支持基本类型、切片、映射和结构体的转换
   // 对于复杂类型使用JSON序列化
   //
   // 性能特性：
   // - 基本类型转换: O(1)
   // - 复杂类型转换: O(n) where n is serialization complexity
   //
   // 示例：
   //   str := ToString(123)        // "123"
   //   str := ToString([]int{1,2}) // "[1,2]"
   func ToString(v interface{}) string
   ```

2. **Package Documentation**
   - Include package overview in package comment
   - Provide usage examples
   - Document performance characteristics
   - Include migration guides for breaking changes

### README Guidelines

1. **Package README**
   - Clear description of package purpose
   - Installation instructions
   - Basic usage examples
   - Performance benchmarks
   - API reference links

2. **Repository README**
   - Project overview
   - Quick start guide
   - Package directory
   - Contributing guidelines
   - License information

## 🚦 Release Process

### Version Management

1. **Semantic Versioning**
   - MAJOR: Breaking changes
   - MINOR: New features, backwards compatible
   - PATCH: Bug fixes, backwards compatible

2. **Release Preparation**
   - Update CHANGELOG.md
   - Update version in relevant files
   - Ensure all tests pass
   - Update documentation

3. **Release Checklist**
   - [ ] All tests pass
   - [ ] Documentation updated
   - [ ] CHANGELOG.md updated
   - [ ] Version tagged
   - [ ] Release notes prepared

## 🐛 Issue Guidelines

### Bug Reports

```markdown
**Bug Description**
Clear description of the bug

**Reproduction Steps**
1. Step one
2. Step two
3. Step three

**Expected Behavior**
What should happen

**Actual Behavior**
What actually happened

**Environment**
- Go version:
- OS:
- Package version:

**Additional Context**
Any additional information
```

### Feature Requests

```markdown
**Feature Description**
Clear description of the proposed feature

**Use Case**
Why is this feature needed?

**Proposed Implementation**
How should this be implemented?

**Alternatives Considered**
Other approaches considered

**Additional Context**
Any additional information
```

## 🤝 Community Guidelines

### Code of Conduct

1. **Be Respectful**
   - Treat all contributors with respect
   - Be constructive in feedback
   - Welcome newcomers

2. **Be Collaborative**
   - Share knowledge and help others
   - Provide clear, helpful reviews
   - Communicate openly about challenges

3. **Be Professional**
   - Focus on the code, not the person
   - Accept criticism gracefully
   - Give credit where due

### Communication Channels

- **GitHub Issues**: Bug reports, feature requests
- **GitHub Discussions**: General questions, ideas
- **Pull Requests**: Code reviews, discussions

## 📖 Learning Resources

### Go Best Practices
- [Effective Go](https://golang.org/doc/effective_go.html)
- [Go Code Review Comments](https://github.com/golang/go/wiki/CodeReviewComments)
- [Go Proverbs](https://go-proverbs.github.io/)

### Performance Optimization
- [Go Performance Tips](https://github.com/golang/go/wiki/Performance)
- [Profiling Go Programs](https://blog.golang.org/profiling-go-programs)

### Testing
- [Testing in Go](https://golang.org/doc/code.html#Testing)
- [Advanced Testing Patterns](https://golang.org/doc/tutorial/add-a-test)

## 📞 Getting Help

If you need help or have questions:

1. Check existing documentation
2. Search existing issues
3. Create a new issue with a clear description
4. Join our community discussions

Thank you for contributing to LazyGophers Utils! Your contributions help make this library better for everyone.