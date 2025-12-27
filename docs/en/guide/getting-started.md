---
title: Getting Started
---

# Getting Started

This guide will help you get started with LazyGophers Utils quickly.

## Installation

Install LazyGophers Utils using Go modules:

```bash
go get github.com/lazygophers/utils
```

## Basic Usage

### Error Handling

LazyGophers Utils provides simplified error handling:

```go
package main

import (
    "fmt"
    "github.com/lazygophers/utils"
)

func main() {
    // Use Must to simplify error handling
    data := utils.Must(loadData())
    fmt.Println(data)
}

func loadData() (string, error) {
    return "Hello, World!", nil
}
```

### Type Conversion

Use the `candy` module for type conversion:

```go
import "github.com/lazygophers/utils/candy"

// String to integer
age := candy.ToInt("25")

// String to boolean
active := candy.ToBool("true")

// String to float
price := candy.ToFloat("99.99")
```

### Time Handling

Use the `xtime` module for time processing:

```go
import "github.com/lazygophers/utils/xtime"

// Get current calendar
cal := xtime.NowCalendar()

// Format date
fmt.Printf("Today: %s\n", cal.String())

// Get lunar date
fmt.Printf("Lunar: %s\n", cal.LunarDate())

// Get zodiac animal
fmt.Printf("Animal: %s\n", cal.Animal())

// Get solar term
fmt.Printf("Solar Term: %s\n", cal.CurrentSolarTerm())
```

### Configuration Management

Use the `config` module to load configuration:

```go
import "github.com/lazygophers/utils/config"

type Config struct {
    Database string `json:"database"`
    Port     int    `json:"port"`
    Debug    bool   `json:"debug"`
}

func main() {
    var cfg Config
    utils.MustSuccess(config.Load(&cfg, "config.json"))
    fmt.Printf("Config: %+v\n", cfg)
}
```

### Data Validation

Use the `validator` module to validate data:

```go
import "github.com/lazygophers/utils/validator"

type User struct {
    Name  string `validate:"required"`
    Email string `validate:"required,email"`
    Age   int    `validate:"min=0,max=150"`
}

func main() {
    user := User{
        Name:  "John Doe",
        Email: "john@example.com",
        Age:   25,
    }

    if err := utils.Validate(&user); err != nil {
        fmt.Printf("Validation failed: %v\n", err)
    } else {
        fmt.Println("Validation successful")
    }
}
```

## Next Steps

- Check out the [Module Overview](/en/modules/overview) to learn about all available modules
- Read the [API Documentation](/en/api/overview) for detailed API information
- Visit the [GitHub Repository](https://github.com/lazygophers/utils) for more examples

## Get Help

- 📖 [Complete API Reference](https://pkg.go.dev/github.com/lazygophers/utils)
- 🐛 [Report Issues](https://github.com/lazygophers/utils/issues)
- 💬 [GitHub Discussions](https://github.com/lazygophers/utils/discussions)
