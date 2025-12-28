---
title: singledo - Singleton Execution
---

# singledo - Singleton Execution

## Overview

The singledo module ensures that operations run only once, preventing duplicate calculations or executions.

## Types

### Single[T]

Generic singleton executor.

```go
type Single[T any] struct {
    mux    sync.Mutex
    last   time.Time
    wait   time.Duration
    call   *call[T]
    result T
}
```

---

## Constructor

### NewSingle[T]()

Create new singleton executor.

```go
func NewSingle[T any](wait time.Duration) *Single[T]
```

**Parameters:**
- `wait` - Wait duration before allowing re-execution

**Returns:**
- Singleton executor instance

**Example:**
```go
single := singledo.NewSingle[string](time.Minute)
```

---

## Methods

### Do()

Execute function with singleton guarantee.

```go
func (s *Single[T]) Do(fn func() (T, error)) (v T, err error)
```

**Parameters:**
- `fn` - Function to execute

**Returns:**
- Result from function
- Error from function

**Behavior:**
- If within wait duration, returns cached result
- Otherwise, executes function and caches result

**Example:**
```go
single := singledo.NewSingle[string](time.Minute)

result, err := single.Do(func() (string, error) {
    log.Info("Executing expensive operation...")
    
    data, err := fetchExpensiveData()
    if err != nil {
        return "", err
    }
    
    return data, nil
})

if err != nil {
    log.Errorf("Operation failed: %v", err)
} else {
    log.Infof("Result: %s", result)
}
```

---

### Reset()

Reset the singleton executor.

```go
func (s *Single[T]) Reset()
```

**Example:**
```go
single.Reset()
```

---

## Usage Patterns

### Expensive Calculations

```go
var calculationSingle = singledo.NewSingle[int64](time.Hour)

func getExpensiveResult() (int64, error) {
    return calculationSingle.Do(func() (int64, error) {
        log.Info("Performing expensive calculation...")
        
        result, err := performExpensiveCalculation()
        if err != nil {
            return 0, err
        }
        
        return result, nil
    })
}
```

### Data Loading

```go
var dataLoaderSingle = singledo.NewSingle[[]byte](5 * time.Minute)

func loadData() ([]byte, error) {
    return dataLoaderSingle.Do(func() ([]byte, error) {
        log.Info("Loading data from remote...")
        
        data, err := http.Get("https://api.example.com/data")
        if err != nil {
            return nil, err
        }
        
        return data, nil
    })
}
```

### Configuration Loading

```go
var configSingle = singledo.NewSingle[Config](time.Hour)

func getConfig() (*Config, error) {
    return configSingle.Do(func() (*Config, error) {
        log.Info("Loading configuration...")
        
        cfg, err := loadConfiguration()
        if err != nil {
            return nil, err
        }
        
        return cfg, nil
    })
}
```

---

## Best Practices

### Wait Duration

```go
// Good: Use appropriate wait duration
var shortCache = singledo.NewSingle[string](time.Minute)      // 1 minute
var mediumCache = singledo.NewSingle[string](time.Hour)      // 1 hour
var longCache = singledo.NewSingle[string](24 * time.Hour) // 24 hours
```

### Error Handling

```go
// Good: Handle errors properly
result, err := single.Do(func() (string, error) {
    data, err := fetchData()
    if err != nil {
        return "", err
    }
    
    return processData(data), nil
})

if err != nil {
    log.Errorf("Operation failed: %v", err)
} else {
    log.Infof("Result: %s", result)
}
```

---

## Related Documentation

- [routine](/en/modules/routine) - Goroutine management
- [wait](/en/modules/wait) - Flow control
- [API Documentation](/en/api/overview)
- [Module Overview](/en/modules/overview)
