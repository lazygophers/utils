# Pull Request

## 📋 Description

<!-- Provide a clear and concise description of what this PR does -->

### Type of Change

<!-- Mark the relevant option with an "x" -->

- [ ] 🐛 Bug fix (non-breaking change which fixes an issue)
- [ ] ✨ New feature (non-breaking change which adds functionality)
- [ ] 💥 Breaking change (fix or feature that would cause existing functionality to not work as expected)
- [ ] 📚 Documentation update
- [ ] 🎨 Code style/formatting changes
- [ ] ♻️ Refactoring (no functional changes)
- [ ] ⚡ Performance improvement
- [ ] 🧪 Test addition or improvement
- [ ] 🔧 Build/CI changes
- [ ] 🗑️ Chore (dependency updates, etc.)

## 🔗 Related Issues

<!-- Link to related issues using keywords like "Fixes", "Closes", "Resolves" -->
<!-- Example: Fixes #123, Closes #456 -->

- Related to #
- Fixes #
- Closes #

## 📝 Changes Made

<!-- List the main changes in this PR -->

- 
- 
- 

## 🧪 Testing

### Test Coverage

- [ ] Unit tests added/updated
- [ ] Integration tests added/updated
- [ ] Benchmark tests added/updated
- [ ] All existing tests pass
- [ ] New code is covered by tests

### Manual Testing

<!-- Describe how you tested your changes -->

```bash
# Commands you ran to test
go test ./...
go run examples/...
```

### Test Results

<!-- Include relevant test output, benchmark results, etc. -->

```
# Test output
```

## 📖 Documentation

- [ ] Code is self-documenting with clear naming
- [ ] Public APIs have godoc comments
- [ ] README updated (if applicable)
- [ ] Documentation updated (if applicable)
- [ ] Examples added/updated (if applicable)
- [ ] CHANGELOG updated

## 🔄 Migration Guide

<!-- If this is a breaking change, provide migration instructions -->

### Before (Old API)

```go
// Old way of doing things
```

### After (New API)

```go
// New way of doing things
```

## ⚡ Performance Impact

<!-- If applicable, describe performance implications -->

- [ ] No performance impact
- [ ] Performance improved
- [ ] Performance regression (explain why it's necessary)

### Benchmark Results

<!-- Include benchmark comparisons if applicable -->

```
# Before
BenchmarkOld-8    1000000    1234 ns/op    567 B/op    8 allocs/op

# After  
BenchmarkNew-8    2000000     617 ns/op    284 B/op    4 allocs/op
```

## 🔒 Security Considerations

- [ ] No security implications
- [ ] Security review required
- [ ] Potential security impact (describe below)

<!-- If there are security implications, describe them -->

## 🌍 Backward Compatibility

- [ ] Fully backward compatible
- [ ] Minor breaking changes (documented above)
- [ ] Major breaking changes (documented above)

## 📋 Checklist

### Code Quality

- [ ] Code follows the project's style guidelines
- [ ] Self-review of code completed
- [ ] Code is properly formatted (`go fmt`)
- [ ] Code passes linting (`golangci-lint run`)
- [ ] No unnecessary dependencies added
- [ ] Error handling is appropriate
- [ ] Code is readable and maintainable

### Testing

- [ ] All tests pass locally
- [ ] Test coverage is maintained or improved
- [ ] Edge cases are covered
- [ ] Performance tests added (if applicable)

### Documentation

- [ ] Code is well-commented
- [ ] Public functions have godoc
- [ ] Complex logic is explained
- [ ] Breaking changes are documented

### Dependencies

- [ ] New dependencies are justified
- [ ] Dependencies are up to date
- [ ] No vulnerable dependencies
- [ ] License compatibility checked

## 🎯 Focus Areas for Review

<!-- Highlight specific areas where you'd like reviewers to focus -->

- [ ] Algorithm correctness
- [ ] Performance optimization
- [ ] Error handling
- [ ] API design
- [ ] Thread safety
- [ ] Memory usage
- [ ] Security implications

## 📷 Screenshots/Examples

<!-- If applicable, add screenshots or code examples showing the changes -->

## 🤝 Reviewer Notes

<!-- Any specific notes for reviewers -->

---

### 📜 Code of Conduct

By submitting this pull request, I confirm that my contribution is made under the terms of the project's license and I agree to follow the [Code of Conduct](https://github.com/lazygophers/utils/blob/master/CODE_OF_CONDUCT.md).

<!-- 
🙏 Thank you for contributing to LazyGophers Utils!
Your contribution helps make Go development more efficient for everyone.
-->