---
title: human
description: Human-friendly data formatting tools
---

# human

The `human` package provides functionality to convert technical data into human-readable formats.

## Use Cases

- **Monitoring displays**: Convert bytes, speeds, and durations to intuitive formats like "KB/s", "2 minutes ago"
- **Log output**: Provide readable time, size, and speed information
- **Multi-language apps**: Display different formats based on locale

## Key Features

### Size and speed formatting

```go
import "github.com/lazygophers/utils/human"

// Byte size formatting
human.ByteSize(1536)        // "1.5 KB"
human.ByteSize(1048576)     // "1.0 MB"

// Speed formatting (bytes/second)
human.Speed(1048576)         // "1.0 MB/s"

// Bit speed formatting (bits/second)
human.BitSpeed(1000000)      // "1.0 Mbps"
```

### Time formatting

```go
import "time"

// Duration formatting
human.Duration(time.Hour)    // "1 hour"
human.Duration(90*time.Minute) // "1 hour 30 minutes"

// Relative time
human.RelativeTime(time.Now().Add(-time.Hour)) // "1 hour ago"
```

### Multi-language support

```go
// Set default language
human.SetLocale("zh")  // Chinese
human.SetLocale("en")  // English

// Get current language
lang := human.GetLocale()
```

### Precision control

```go
// Set default precision (decimal places)
human.SetDefaultPrecision(2)

human.ByteSize(1536)  // "1.50 KB" (2 decimal places)
```

## Usage Recommendations

1. **Monitoring panels**: Use `Speed` and `ByteSize` instead of raw values
2. **User interfaces**: Use `RelativeTime` to display "how long ago"
3. **Internationalization**: Use with `SetLocale` for multi-language switching

## Notes

- Default language is English (`en`), Chinese apps should call `SetLocale("zh")`
- Precision settings are global and affect all subsequent calls
