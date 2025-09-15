# app - Application Lifecycle Management

The `app` package provides utilities for managing application lifecycle, environment detection, and build information. It helps applications understand their runtime environment and build context.

## Features

- **Build Information**: Access compile-time build information (commit, branch, tag, build date)
- **Environment Detection**: Determine application release type (debug, test, alpha, beta, release)
- **Runtime Context**: Access Go version and target platform information
- **Application Metadata**: Manage application name, version, and organization

## Installation

```bash
go get github.com/lazygophers/utils/app
```

## Usage Examples

### Basic Application Info

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils/app"
)

func main() {
    app.Name = "MyApp"
    app.Version = "1.0.0"
    app.PackageType = app.Release

    fmt.Printf("Application: %s v%s\n", app.Name, app.Version)
    fmt.Printf("Release Type: %s\n", app.PackageType)
    fmt.Printf("Organization: %s\n", app.Organization)
}
```

### Environment Detection

```go
// Check if running in development
if app.IsDebug() {
    fmt.Println("Running in debug mode")
}

// Check if running in production
if app.IsRelease() {
    fmt.Println("Running in production")
}

// Check if running tests
if app.IsTest() {
    fmt.Println("Running tests")
}

// Check if alpha/beta release
if app.IsAlpha() {
    fmt.Println("Alpha release")
}

if app.IsBeta() {
    fmt.Println("Beta release")
}
```

### Build Information

```go
// Access build-time information
fmt.Printf("Git Commit: %s\n", app.Commit)
fmt.Printf("Short Commit: %s\n", app.ShortCommit)
fmt.Printf("Branch: %s\n", app.Branch)
fmt.Printf("Tag: %s\n", app.Tag)
fmt.Printf("Build Date: %s\n", app.BuildDate)
fmt.Printf("Go Version: %s\n", app.GoVersion)
fmt.Printf("Target OS: %s\n", app.GoOS)
fmt.Printf("Target Arch: %s\n", app.Goarch)
```

### Release Type Management

```go
// Set different release types
app.PackageType = app.Debug     // Development
app.PackageType = app.Test      // Testing
app.PackageType = app.Alpha     // Alpha release
app.PackageType = app.Beta      // Beta release
app.PackageType = app.Release   // Production release

// Get string representation
fmt.Println(app.PackageType.String()) // "release", "beta", etc.
```

## API Reference

### Variables

#### Application Information
- `Name string` - Application name
- `Version string` - Application version
- `Organization string` - Organization name (default: "lazygophers")
- `PackageType ReleaseType` - Current release type

#### Build Information
- `Commit string` - Full Git commit hash
- `ShortCommit string` - Short Git commit hash
- `Branch string` - Git branch name
- `Tag string` - Git tag name
- `BuildDate string` - Build timestamp
- `Description string` - Application description

#### Runtime Information
- `GoVersion string` - Go version used for build
- `GoOS string` - Target operating system
- `Goarch string` - Target architecture
- `Goarm string` - ARM variant (if applicable)
- `Goamd64 string` - AMD64 variant (if applicable)
- `Gomips string` - MIPS variant (if applicable)

### Types

#### ReleaseType

```go
type ReleaseType uint8

const (
    Debug   ReleaseType = iota // Development/debugging
    Test                       // Testing environment
    Alpha                      // Alpha release
    Beta                       // Beta release
    Release                    // Production release
)
```

#### Methods

- `String() string` - Get string representation of release type
- `Debug() string` - Get debug string representation

### Functions

#### Environment Detection
- `IsDebug() bool` - Check if running in debug mode
- `IsTest() bool` - Check if running in test mode
- `IsAlpha() bool` - Check if running alpha release
- `IsBeta() bool` - Check if running beta release
- `IsRelease() bool` - Check if running production release

## Build Integration

To populate build information at compile time, use Go's `-ldflags`:

```bash
go build -ldflags "
    -X github.com/lazygophers/utils/app.Name=MyApp
    -X github.com/lazygophers/utils/app.Version=1.0.0
    -X github.com/lazygophers/utils/app.Commit=$(git rev-parse HEAD)
    -X github.com/lazygophers/utils/app.ShortCommit=$(git rev-parse --short HEAD)
    -X github.com/lazygophers/utils/app.Branch=$(git branch --show-current)
    -X github.com/lazygophers/utils/app.BuildDate=$(date -u +'%Y-%m-%dT%H:%M:%SZ')
"
```

## Usage Patterns

### Configuration Based on Environment

```go
func setupLogging() {
    switch app.PackageType {
    case app.Debug:
        // Enable debug logging
        log.SetLevel(log.DebugLevel)
    case app.Release:
        // Production logging
        log.SetLevel(log.InfoLevel)
    default:
        // Default logging
        log.SetLevel(log.WarnLevel)
    }
}
```

### Feature Flags

```go
func enableExperimentalFeatures() bool {
    return app.IsDebug() || app.IsAlpha()
}

func enableMetrics() bool {
    return app.IsRelease() || app.IsBeta()
}
```

## Best Practices

1. **Set Application Info Early**: Set `Name`, `Version`, and `PackageType` in your application's init function
2. **Use Environment Detection**: Use environment detection functions for conditional behavior
3. **Build Integration**: Integrate build information injection into your CI/CD pipeline
4. **Version Management**: Use semantic versioning for the `Version` field

## Related Packages

- `config` - Configuration management
- `runtime` - Runtime system information
- `atexit` - Graceful shutdown handling