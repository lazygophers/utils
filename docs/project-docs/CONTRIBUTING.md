# Contributing Guide

Welcome to contribute to the LazyGophers Utils project! We greatly appreciate every contribution from the community.

[![Contributors](https://img.shields.io/badge/Contributors-Welcome-brightgreen.svg)](#how-to-contribute)
[![Code Style](https://img.shields.io/badge/Code%20Style-Go%20Standard-blue.svg)](#code-standards)

## 🤝 How to Contribute

### Types of Contributions

We welcome the following types of contributions:

- 🐛 **Bug Fixes** - Fix known issues
- ✨ **New Features** - Add new utility functions or modules
- 📚 **Documentation Improvements** - Enhance documentation, add examples
- 🎨 **Code Optimization** - Performance optimization, refactoring
- 🧪 **Test Improvements** - Increase test coverage, fix test issues
- 🌐 **Internationalization** - Add multi-language support

### Contribution Process

#### 1. Preparation

**Fork the Project**
```bash
# 1. Fork this project to your GitHub account
# 2. Clone your fork locally
git clone https://github.com/YOUR_USERNAME/utils.git
cd utils

# 3. Add the original project as upstream repository
git remote add upstream https://github.com/lazygophers/utils.git

# 4. Create a new feature branch
git checkout -b feature/your-awesome-feature
```

**Set up Development Environment**
```bash
# Install dependencies
go mod tidy

# Verify environment
go version  # Requires Go 1.24.0+
go test ./... # Ensure all tests pass
```

#### 2. Development Phase

**Write Code**
- Follow [Code Standards](#code-standards)
- Write test cases for new features
- Ensure test coverage doesn't drop below current level
- Add necessary documentation comments

**Commit Standards**
```bash
# Use standardized commit message format
git commit -m "feat(module): add new utility function

- Add FormatDuration function
- Support multiple time format outputs
- Add comprehensive test cases
- Update related documentation

Closes #123"
```

**Commit Message Format**:
```
<type>(<scope>): <subject>

<body>

<footer>
```

**Type Categories**:
- `feat`: New features
- `fix`: Bug fixes  
- `docs`: Documentation updates
- `style`: Code formatting adjustments
- `refactor`: Code refactoring
- `perf`: Performance optimization
- `test`: Test-related
- `chore`: Build tools or dependency updates

**Scope Range** (optional):
- `candy`: candy module
- `xtime`: xtime module
- `config`: config module
- `cryptox`: cryptox module
- etc...

#### 3. Testing and Validation

**Run Tests**
```bash
# Run all tests
go test -v ./...

# Check test coverage
go test -cover -v ./...

# Run benchmark tests
go test -bench=. ./...

# Check code formatting
go fmt ./...

# Static analysis
go vet ./...
```

**Performance Testing**
```bash
# Run performance tests
go test -bench=BenchmarkYourFunction -benchmem ./...

# Ensure no significant performance regression
```

#### 4. Create Pull Request

**Push to Your Fork**
```bash
git push origin feature/your-awesome-feature
```

**Create PR**
1. Visit the project page on GitHub
2. Click "New Pull Request"
3. Select your branch
4. Fill in PR description (refer to [PR Template](#pr-template))
5. Ensure all checks pass

#### 5. Code Review

- Maintainers will review your code
- Make modifications based on feedback
- Maintain communication and cooperative attitude
- Will be merged after tests pass

## 📝 Code Standards

### Go Code Style

**Basic Standards**
```go
// ✅ Good example
package candy

import (
    "context"
    "fmt"
    "time"
    
    "github.com/lazygophers/log"
)

// FormatDuration formats time duration into human-readable string
// Supports multiple precision levels, automatically chooses appropriate units
//
// Parameters:
//   - duration: time duration to format
//   - precision: precision level (1-3)
//
// Returns:
//   - string: formatted string, like "2 hours 30 minutes"
//
// Example:
//   FormatDuration(90*time.Minute, 2) // returns "1 hour 30 minutes"
//   FormatDuration(45*time.Second, 1) // returns "45 seconds"
func FormatDuration(duration time.Duration, precision int) string {
    if duration == 0 {
        return "0 seconds"
    }
    
    // Implementation logic...
    return result
}
```

**Naming Conventions**
- Use CamelCase
- Function names start with verbs: `Get`, `Set`, `Format`, `Parse`
- Constants use ALL_CAPS: `const MaxRetries = 3`
- Private members use lowercase: `internalHelper`
- Package names use lowercase single words: `candy`, `xtime`

**Comment Standards**
- All public functions must have comments
- Comments start with function name
- Include parameter and return value descriptions  
- Provide usage examples
- English comments, concise and clear

**Error Handling**
```go
// ✅ Recommended error handling approach
func ProcessData(data []byte) (*Result, error) {
    if len(data) == 0 {
        log.Warn("Empty data provided")
        return nil, fmt.Errorf("data cannot be empty")
    }
    
    result, err := parseData(data)
    if err != nil {
        log.Error("Failed to parse data", log.Error(err))
        return nil, fmt.Errorf("parse data failed: %w", err)
    }
    
    return result, nil
}
```

### Project Structure Standards

**Module Organization**
```
utils/
├── README.md           # Project overview
├── CONTRIBUTING.md     # Contributing guide  
├── SECURITY.md        # Security policy
├── go.mod             # Go module definition
├── must.go            # Core utility functions
├── candy/             # Data processing tools
│   ├── README.md      # Module documentation
│   ├── to_string.go   # Type conversion
│   └── to_string_test.go
├── xtime/             # Time processing tools  
│   ├── README.md      # Detailed usage documentation
│   ├── TESTING.md     # Test reports
│   ├── PERFORMANCE.md # Performance reports
│   ├── calendar.go    # Calendar functionality
│   └── calendar_test.go
└── ...
```

**File Naming**
- Use lowercase letters and underscores: `to_string.go`
- Test file suffix: `_test.go`
- Benchmark tests: `_benchmark_test.go`
- Documentation files: `README.md`, `TESTING.md`

### Testing Standards

**Test Coverage Requirements**
- New feature test coverage must be ≥ 90%
- Cannot reduce overall test coverage
- Include normal and edge cases
- Error handling paths must be tested

**Test Example**
```go
func TestFormatDuration(t *testing.T) {
    testCases := []struct {
        name      string
        duration  time.Duration
        precision int
        want      string
    }{
        {
            name:      "zero time",
            duration:  0,
            precision: 1,
            want:      "0 seconds",
        },
        {
            name:      "90 minutes high precision",
            duration:  90 * time.Minute,
            precision: 2,
            want:      "1 hour 30 minutes",
        },
        // More test cases...
    }
    
    for _, tc := range testCases {
        t.Run(tc.name, func(t *testing.T) {
            got := FormatDuration(tc.duration, tc.precision)
            assert.Equal(t, tc.want, got)
        })
    }
}

// Benchmark test
func BenchmarkFormatDuration(b *testing.B) {
    duration := 90 * time.Minute
    
    b.ResetTimer()
    for i := 0; i < b.N; i++ {
        _ = FormatDuration(duration, 2)
    }
}
```

## 🎯 Key Development Areas

### High Priority

1. **xtime Module Enhancement**
   - Lunar calendar and solar terms functionality improvement
   - Performance optimization
   - More cultural-specific features

2. **candy Module Extension**  
   - Type conversion functions
   - Data processing tools
   - Performance optimization

3. **Test Coverage Improvement**
   - Target: All modules > 90%
   - Edge case supplementation
   - Performance test improvement

### Medium Priority

4. **New Utility Modules**
   - AI/ML utility functions
   - Cloud service integration tools
   - Microservice tools

5. **Documentation Enhancement**
   - API reference documentation
   - Best practices guide
   - Performance optimization guide

### Welcome Contributions

- 🌏 **Multi-language Support** - English documentation, error message internationalization
- 📊 **More Data Format Support** - XML, YAML, TOML processing
- 🔧 **Development Tools** - Code generation, configuration management
- 🎨 **UI/UX Tools** - Color processing, formatted output
- 🔐 **Security Tools** - Encryption/decryption, signature verification

## 📋 PR Template

Please use the following template when creating a PR:

```markdown
## Change Description

Brief description of the content and purpose of this change.

## Change Type

- [ ] Bug fix
- [ ] New feature
- [ ] Documentation update
- [ ] Performance optimization  
- [ ] Code refactoring
- [ ] Test improvement

## Detailed Changes

### New Features
- Added `FormatDuration` function
- Support multiple precision levels
- Added Chinese time unit display

### Fixed Issues  
- Fixed timezone conversion bug (#123)
- Resolved memory leak issue

### Performance Optimization
- Optimized string concatenation performance
- Reduced memory allocation by 30%

## Testing Description

- [ ] All tests pass
- [ ] Added new test cases
- [ ] Test coverage ≥ 90%
- [ ] Benchmark tests pass

**Test Coverage**: 92.5%

## Documentation Updates

- [ ] Updated README.md
- [ ] Added function comments
- [ ] Updated example code

## Compatibility

- [ ] Backward compatible
- [ ] Requires version upgrade (explain reason)
- [ ] Breaking changes (detailed explanation)

## Checklist

- [ ] Code follows project standards
- [ ] Passed `go fmt` format check
- [ ] Passed `go vet` static check
- [ ] All tests pass
- [ ] Documentation updated
- [ ] Commit messages follow standards

## Related Issues

Closes #123
Refs #456

## Screenshots/Demo

Provide screenshots or demos if necessary.
```

## 🐛 Bug Reports

Found a bug? Please use the following template to create an Issue:

```markdown
## Bug Description

Brief description of the issue encountered.

## Reproduction Steps

1. Execute step 1
2. Execute step 2  
3. Observe result

## Expected Behavior

Describe the correct behavior you expect to see.

## Actual Behavior

Describe the actual erroneous behavior observed.

## Environment Information

- **Operating System**: macOS 12.0
- **Go Version**: 1.24.0
- **Utils Version**: v1.2.0
- **Other relevant information**:

## Error Logs

```
paste error logs here
```

## Minimal Reproducible Example

```go
package main

import (
    "github.com/lazygophers/utils/xtime"
)

func main() {
    // Minimal error reproduction code
}
```
```

## ✨ Feature Requests

Want a new feature? Please use the following template:

```markdown
## Feature Description

Describe the feature you'd like to add.

## Use Cases

Describe when this feature would be used.

## Suggested API Design

```go
// Suggested function signature and usage
func NewAwesomeFunction(param string) (Result, error) {
    // ...
}
```

## Alternative Solutions

Have you considered other solutions?

## Additional Information

Other relevant information or references.
```

## 🏆 Contributor Recognition

### Contribution Type Recognition

We will give different recognition based on contribution types:

- 🥇 **Core Contributors** - Long-term active, important feature contributions
- 🥈 **Active Contributors** - Multiple valuable contributions  
- 🥉 **Community Contributors** - Bug fixes, documentation improvements
- 🌟 **First-time Contributors** - Welcome first-time contributions

### Contribution Statistics

We will showcase contributors in the following places:

- README.md contributor list
- Acknowledgments in release notes
- Project website (if available)
- Annual contributor reports

## 💬 Communication

### Getting Help

- 📖 **Documentation Issues**: Check README.md for each module
- 🐛 **Bug Reports**: [GitHub Issues](https://github.com/lazygophers/utils/issues)
- 💡 **Feature Discussions**: [GitHub Discussions](https://github.com/lazygophers/utils/discussions)
- ❓ **Usage Questions**: [GitHub Discussions Q&A](https://github.com/lazygophers/utils/discussions/categories/q-a)

### Discussion Standards

Please follow these communication standards:

- Use friendly and professional language
- Provide detailed problem descriptions and suggestions
- Provide sufficient context information
- Respect different viewpoints and opinions
- Actively participate in constructive discussions

## 📜 License

This project is licensed under the [GNU Affero General Public License v3.0](LICENSE).

**Contributing means agreeing**:
- You own the copyright to the submitted code
- Agree to release code under AGPL v3.0 license
- Follow the project's contributor code of conduct

## 🙏 Acknowledgments

Thanks to all developers who have contributed to the LazyGophers Utils project!

**Special Thanks**:
- All contributors who submitted Issues and PRs
- Community members who provided suggestions and feedback
- Volunteers who helped improve documentation

---

**Available in other languages:** [简体中文](CONTRIBUTING_zh.md) | [繁體中文](CONTRIBUTING_zh-Hant.md) | [Français](CONTRIBUTING_fr.md) | [Русский](CONTRIBUTING_ru.md) | [Español](CONTRIBUTING_es.md) | [العربية](CONTRIBUTING_ar.md)

**Happy Coding! 🎉**

Feel free to contact the maintainer team anytime if you have questions. We're happy to help you start your contribution journey!