# Pyroscope - Performance Profiling Integration

The `pyroscope` module provides seamless integration with Grafana Pyroscope for continuous performance profiling in Go applications. It offers build-conditional profiling that automatically enables/disables based on build tags, making it safe for production deployments while providing detailed performance insights during development.

## Features

- **Build-Conditional Profiling**: Automatically enabled/disabled based on build tags
- **Comprehensive Profile Types**: CPU, memory, goroutines, mutex, and block profiling
- **Zero Production Overhead**: No-op implementation when disabled
- **Automatic Application Detection**: Uses app name from the application context
- **Hostname Tagging**: Automatically tags profiles with hostname for multi-instance deployments
- **Configurable Server Address**: Flexible server endpoint configuration
- **GC Optimization**: Disabled garbage collection runs for more accurate profiling

## Installation

```bash
go get github.com/lazygophers/utils
```

## Pyroscope Server Setup

Before using the module, you need a running Pyroscope server. The easiest way is using Docker:

```bash
# Run Pyroscope server with Docker
docker run -itd -p 4040:4040 pyroscope/pyroscope:latest server
```

Or using Docker Compose:

```yaml
version: '3.8'
services:
  pyroscope:
    image: pyroscope/pyroscope:latest
    command: server
    ports:
      - "4040:4040"
    environment:
      - PYROSCOPE_LOG_LEVEL=debug
    volumes:
      - pyroscope-data:/var/lib/pyroscope

volumes:
  pyroscope-data:
```

## Usage

### Basic Integration

```go
package main

import (
    "github.com/lazygophers/utils/pyroscope"
    "time"
)

func main() {
    // Initialize profiling with default Pyroscope server
    pyroscope.Load("")

    // Your application logic here
    doWork()

    // Keep application running to collect profiles
    time.Sleep(60 * time.Second)
}

func doWork() {
    // CPU-intensive work that will be profiled
    for i := 0; i < 1000000; i++ {
        _ = computeSomething(i)
    }
}

func computeSomething(n int) int {
    return n * n * n
}
```

### Custom Server Address

```go
package main

import (
    "github.com/lazygophers/utils/pyroscope"
    "os"
)

func main() {
    // Use custom Pyroscope server address
    serverURL := os.Getenv("PYROSCOPE_URL")
    if serverURL == "" {
        serverURL = "http://pyroscope.example.com:4040"
    }

    pyroscope.Load(serverURL)

    // Application logic...
    runApplication()
}

func runApplication() {
    // Your application code here
}
```

### Web Application Integration

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/lazygophers/utils/pyroscope"
    "net/http"
    "time"
)

func main() {
    // Initialize profiling first
    pyroscope.Load("http://localhost:4040")

    // Set up web server
    r := gin.Default()

    r.GET("/api/heavy", heavyComputation)
    r.GET("/api/memory", memoryIntensive)
    r.GET("/api/concurrent", concurrentWork)

    // Start server
    r.Run(":8080")
}

func heavyComputation(c *gin.Context) {
    start := time.Now()

    // CPU-intensive work
    result := 0
    for i := 0; i < 10000000; i++ {
        result += i * i
    }

    c.JSON(http.StatusOK, gin.H{
        "result":   result,
        "duration": time.Since(start).String(),
    })
}

func memoryIntensive(c *gin.Context) {
    // Memory allocation intensive work
    data := make([][]int, 1000)
    for i := range data {
        data[i] = make([]int, 1000)
        for j := range data[i] {
            data[i][j] = i * j
        }
    }

    c.JSON(http.StatusOK, gin.H{
        "allocated": len(data) * len(data[0]),
    })
}

func concurrentWork(c *gin.Context) {
    done := make(chan bool, 10)

    // Spawn goroutines for concurrent work
    for i := 0; i < 10; i++ {
        go func(id int) {
            time.Sleep(100 * time.Millisecond)
            for j := 0; j < 100000; j++ {
                _ = j * id
            }
            done <- true
        }(i)
    }

    // Wait for all goroutines
    for i := 0; i < 10; i++ {
        <-done
    }

    c.JSON(http.StatusOK, gin.H{
        "goroutines_completed": 10,
    })
}
```

### Background Service Integration

```go
package main

import (
    "context"
    "github.com/lazygophers/utils/pyroscope"
    "sync"
    "time"
)

func main() {
    // Initialize profiling
    pyroscope.Load("")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
    defer cancel()

    var wg sync.WaitGroup

    // Start background workers
    for i := 0; i < 5; i++ {
        wg.Add(1)
        go backgroundWorker(ctx, &wg, i)
    }

    wg.Wait()
}

func backgroundWorker(ctx context.Context, wg *sync.WaitGroup, id int) {
    defer wg.Done()

    ticker := time.NewTicker(1 * time.Second)
    defer ticker.Stop()

    for {
        select {
        case <-ctx.Done():
            return
        case <-ticker.C:
            // Simulate work that varies by worker
            processData(id)
        }
    }
}

func processData(workerID int) {
    // Different workloads for different workers
    switch workerID % 3 {
    case 0:
        // CPU intensive
        sum := 0
        for i := 0; i < 1000000; i++ {
            sum += i
        }
    case 1:
        // Memory intensive
        data := make([]int, 100000)
        for i := range data {
            data[i] = i * workerID
        }
    case 2:
        // I/O simulation with goroutines
        done := make(chan bool, 10)
        for i := 0; i < 10; i++ {
            go func() {
                time.Sleep(10 * time.Millisecond)
                done <- true
            }()
        }
        for i := 0; i < 10; i++ {
            <-done
        }
    }
}
```

## Build Configuration

### Development Build (Profiling Enabled)

```bash
# Default build enables profiling
go build -o myapp .

# Or explicitly enable profiling
go build -tags="!release" -o myapp .

# Enable profiling in release builds
go build -tags="release,pyroscope" -o myapp .
```

### Production Build (Profiling Disabled)

```bash
# Production build disables profiling
go build -tags="release" -o myapp .
```

### Build Tag Combinations

| Build Command | Profiling Status | Use Case |
|---|---|---|
| `go build` | Enabled | Development |
| `go build -tags="!release"` | Enabled | Development (explicit) |
| `go build -tags="release"` | Disabled | Production (default) |
| `go build -tags="release,pyroscope"` | Enabled | Production with profiling |

## Environment Configuration

```go
package main

import (
    "github.com/lazygophers/utils/pyroscope"
    "os"
)

func initProfiling() {
    // Check if profiling should be enabled
    if os.Getenv("ENABLE_PROFILING") != "true" {
        return
    }

    // Get server address from environment
    serverAddr := os.Getenv("PYROSCOPE_SERVER")
    if serverAddr == "" {
        serverAddr = "http://localhost:4040"
    }

    pyroscope.Load(serverAddr)
}

func main() {
    initProfiling()

    // Application logic...
}
```

With environment variables:

```bash
# Enable profiling with custom server
ENABLE_PROFILING=true PYROSCOPE_SERVER=http://pyroscope-prod:4040 ./myapp

# Run with hostname tag
HOSTNAME=server-01 ./myapp
```

## Profile Types

The module automatically enables comprehensive profiling:

### CPU Profiling
- **Purpose**: Identifies CPU hotspots and performance bottlenecks
- **Use Case**: Optimize computational intensive code paths

### Memory Profiling
- **InuseObjects**: Currently allocated objects count
- **AllocObjects**: Total allocated objects (including freed)
- **InuseSpace**: Currently allocated memory size
- **AllocSpace**: Total allocated memory (including freed)
- **Use Case**: Detect memory leaks and optimize allocations

### Goroutine Profiling
- **Purpose**: Track goroutine creation and lifecycle
- **Use Case**: Debug goroutine leaks and concurrency issues

### Mutex Profiling
- **MutexCount**: Mutex contention events count
- **MutexDuration**: Time spent waiting for mutexes
- **Use Case**: Optimize synchronization and reduce lock contention

### Block Profiling
- **BlockCount**: Blocking events count
- **BlockDuration**: Time spent in blocking operations
- **Use Case**: Identify blocking I/O and synchronization issues

## Advanced Usage

### Custom Application Name

```go
package main

import (
    "github.com/lazygophers/utils/app"
    "github.com/lazygophers/utils/pyroscope"
)

func init() {
    // Set custom application name before profiling
    app.Name = "my-custom-service"
}

func main() {
    pyroscope.Load("")
    // Application will be profiled as "my-custom-service"
}
```

### Conditional Profiling

```go
package main

import (
    "github.com/lazygophers/utils/pyroscope"
    "os"
)

func main() {
    // Only enable profiling in specific environments
    environment := os.Getenv("ENVIRONMENT")

    switch environment {
    case "development", "staging":
        pyroscope.Load("")
    case "production":
        if os.Getenv("DEBUG_MODE") == "true" {
            pyroscope.Load("http://internal-pyroscope:4040")
        }
    }

    // Application logic...
}
```

### Docker Integration

Dockerfile with conditional profiling:

```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .

# Build with release tag by default
ARG BUILD_TAGS="release"
RUN go build -tags="${BUILD_TAGS}" -o myapp .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/myapp .

# Set default environment
ENV PYROSCOPE_SERVER=http://pyroscope:4040

CMD ["./myapp"]
```

Build commands:

```bash
# Production build (no profiling)
docker build -t myapp:prod .

# Development build (with profiling)
docker build --build-arg BUILD_TAGS="" -t myapp:dev .

# Production build with profiling enabled
docker build --build-arg BUILD_TAGS="release,pyroscope" -t myapp:prod-profile .
```

## Performance Impact

### Development Build
- **CPU Overhead**: 1-5% depending on application characteristics
- **Memory Overhead**: Small increase for profile data collection
- **Network**: Periodic uploads to Pyroscope server

### Production Build (Release Tag)
- **CPU Overhead**: 0% (no-op implementation)
- **Memory Overhead**: 0% (no-op implementation)
- **Binary Size**: Minimal increase due to unused code elimination

### Recommendations
- Always use release builds in production unless debugging
- Enable profiling in staging environments that mirror production
- Use profiling-enabled production builds temporarily for troubleshooting

## Troubleshooting

### Common Issues

#### Profiling Not Working
```go
// Check if profiling is actually enabled
func main() {
    pyroscope.Load("")

    // Add some identifiable work
    for i := 0; i < 1000000; i++ {
        _ = expensiveFunction(i)
    }
}
```

#### Connection Issues
- Verify Pyroscope server is running: `curl http://localhost:4040`
- Check network connectivity from application to server
- Verify server address configuration

#### Missing Profiles
- Ensure application runs long enough to collect samples
- Check application name configuration
- Verify hostname environment variable

### Debug Information

```go
package main

import (
    "github.com/lazygophers/utils/pyroscope"
    "log"
    "os"
)

func main() {
    log.Printf("Application: %s", os.Getenv("APP_NAME"))
    log.Printf("Hostname: %s", os.Getenv("HOSTNAME"))
    log.Printf("Pyroscope Server: %s", os.Getenv("PYROSCOPE_SERVER"))

    pyroscope.Load(os.Getenv("PYROSCOPE_SERVER"))

    // Your application...
}
```

## Best Practices

### 1. Build Configuration
- Use release builds in production by default
- Enable profiling only when needed for debugging
- Document profiling build procedures for your team

### 2. Server Management
- Run Pyroscope server in a stable environment
- Configure appropriate retention policies
- Monitor Pyroscope server resource usage

### 3. Application Integration
- Initialize profiling early in application startup
- Use meaningful application names
- Set appropriate hostname tags for multi-instance deployments

### 4. Performance Monitoring
- Monitor the performance impact of profiling in non-production environments
- Use profiling data to identify and fix performance bottlenecks
- Regularly review and act on profiling insights

## Security Considerations

- Profiling data may contain sensitive information about application internals
- Restrict access to Pyroscope server and profiles
- Consider network security for profile data transmission
- Be cautious with profiling in production environments

## Related Tools and Packages

- [Grafana Pyroscope](https://pyroscope.io/): The continuous profiling platform
- [`runtime/pprof`](https://pkg.go.dev/runtime/pprof): Go's built-in profiling
- [`net/http/pprof`](https://pkg.go.dev/net/http/pprof): HTTP pprof endpoints
- [`github.com/grafana/pyroscope-go`](https://github.com/grafana/pyroscope-go): Underlying Pyroscope Go client

## Example Configurations

### Docker Compose Setup
```yaml
version: '3.8'
services:
  app:
    build: .
    environment:
      - PYROSCOPE_SERVER=http://pyroscope:4040
      - HOSTNAME=app-instance-1
    depends_on:
      - pyroscope

  pyroscope:
    image: pyroscope/pyroscope:latest
    command: server
    ports:
      - "4040:4040"
    volumes:
      - pyroscope-data:/var/lib/pyroscope

volumes:
  pyroscope-data:
```

### Kubernetes Deployment
```yaml
apiVersion: apps/v1
kind: Deployment
metadata:
  name: myapp
spec:
  replicas: 3
  selector:
    matchLabels:
      app: myapp
  template:
    metadata:
      labels:
        app: myapp
    spec:
      containers:
      - name: myapp
        image: myapp:latest
        env:
        - name: PYROSCOPE_SERVER
          value: "http://pyroscope-service:4040"
        - name: HOSTNAME
          valueFrom:
            fieldRef:
              fieldPath: spec.nodeName
```