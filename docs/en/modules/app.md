---
title: app - Application Framework
---

# app - Application Framework

## Overview

The app module provides application framework utilities including organization name, version, and release type management.

## Variables

### Organization

```go
var Organization = "lazygophers"
```

**Description:**
- Organization name for the application
- Used for config and cache directories

---

### Name

```go
var Name string
```

**Description:**
- Application name
- Set at build time

---

### Version

```go
var Version string
```

**Description:**
- Application version
- Set at build time

---

### Build Information

```go
var (
    Commit      string
    ShortCommit string
    Branch      string
    Tag         string
    BuildDate string
    GoVersion string
    GoOS    string
    Goarch  string
    Goarm   string
    Goamd64 string
    Gomips  string
    Description string
)
```

---

## Release Type

### Type Definition

```go
type ReleaseType uint8

const (
    Debug ReleaseType = iota
    Test
    Alpha
    Beta
    Release
)
```

---

### String()

Get release type as string.

```go
func (p ReleaseType) String() string
```

**Returns:**
- "debug", "test", "alpha", "beta", or "release"

**Example:**
```go
fmt.Println(app.Debug.String())    // "debug"
fmt.Println(app.Release.String())  // "release"
```

---

### Debug()

Get release type as debug string.

```go
func (p ReleaseType) Debug() string
```

**Returns:**
- Release type as string

---

## Environment Configuration

### Package Type

```go
var PackageType ReleaseType
```

**Description:**
- Current package type
- Set based on `APP_ENV` environment variable

**Environment Variables:**
- `APP_ENV=dev` or `development` → Debug
- `APP_ENV=test` or `canary` → Test
- `APP_ENV=prod` or `release` or `production` → Release
- `APP_ENV=alpha` → Alpha
- `APP_ENV=beta` → Beta

**Example:**
```bash
# Set environment
export APP_ENV=prod

# In Go code
if app.PackageType == app.Release {
    fmt.Println("Running in production mode")
}
```

---

## Usage Patterns

### Application Information

```go
func printAppInfo() {
    fmt.Printf("Application: %s\n", app.Name)
    fmt.Printf("Version: %s\n", app.Version)
    fmt.Printf("Organization: %s\n", app.Organization)
    fmt.Printf("Release Type: %s\n", app.PackageType.String())
    fmt.Printf("Go Version: %s\n", app.GoVersion)
    fmt.Printf("Build Date: %s\n", app.BuildDate)
}
```

### Environment-Based Behavior

```go
func setupLogging() {
    switch app.PackageType {
    case app.Debug:
        log.SetLevel(log.DebugLevel)
    case app.Test:
        log.SetLevel(log.InfoLevel)
    case app.Release:
        log.SetLevel(log.WarnLevel)
    default:
        log.SetLevel(log.InfoLevel)
    }
}
```

### Feature Flags

```go
func isFeatureEnabled(feature string) bool {
    switch app.PackageType {
    case app.Debug, app.Test:
        return true  // All features enabled in debug/test
    case app.Alpha:
        return isAlphaFeature(feature)
    case app.Beta:
        return isBetaFeature(feature)
    case app.Release:
        return isReleaseFeature(feature)
    default:
        return false
    }
}
```

### Configuration Paths

```go
func getConfigPath() string {
    configDir := runtime.LazyConfigDir()
    return filepath.Join(configDir, app.Name, "config.json")
}

func getCachePath() string {
    cacheDir := runtime.LazyCacheDir()
    return filepath.Join(cacheDir, app.Name, "cache.db")
}
```

---

## Best Practices

### Build Information

```go
// Good: Use build-time variables
func getVersion() string {
    if app.Version != "" {
        return app.Version
    }
    return "dev"
}

// Good: Display build information
func showBuildInfo() {
    fmt.Printf("Version: %s\n", getVersion())
    fmt.Printf("Commit: %s\n", app.Commit)
    fmt.Printf("Build Date: %s\n", app.BuildDate)
}
```

### Environment Handling

```go
// Good: Check package type
func initApplication() {
    switch app.PackageType {
    case app.Debug:
        enableDebugFeatures()
    case app.Release:
        enableProductionFeatures()
    default:
        enableStandardFeatures()
    }
}
```

---

## Related Documentation

- [runtime](/en/modules/runtime) - Runtime information
- [config](/en/modules/config) - Configuration management
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
