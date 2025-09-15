# Pyroscope - 性能分析集成

`pyroscope` 模块提供与 Grafana Pyroscope 的无缝集成，用于 Go 应用程序的持续性能分析。它提供基于构建条件的分析功能，根据构建标签自动启用/禁用，使其在生产部署中安全，同时在开发期间提供详细的性能洞察。

## 功能特性

- **基于构建条件的分析**: 根据构建标签自动启用/禁用
- **全面的分析类型**: CPU、内存、goroutine、互斥锁和阻塞分析
- **零生产开销**: 禁用时为无操作实现
- **自动应用程序检测**: 使用应用程序上下文中的应用程序名称
- **主机名标记**: 自动为多实例部署的分析添加主机名标记
- **可配置服务器地址**: 灵活的服务器端点配置
- **GC 优化**: 禁用垃圾回收运行以获得更准确的分析

## 安装

```bash
go get github.com/lazygophers/utils
```

## Pyroscope 服务器设置

在使用该模块之前，您需要运行 Pyroscope 服务器。最简单的方法是使用 Docker：

```bash
# 使用 Docker 运行 Pyroscope 服务器
docker run -itd -p 4040:4040 pyroscope/pyroscope:latest server
```

或使用 Docker Compose：

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

## 使用方法

### 基本集成

```go
package main

import (
    "github.com/lazygophers/utils/pyroscope"
    "time"
)

func main() {
    // 使用默认 Pyroscope 服务器初始化分析
    pyroscope.Load("")

    // 您的应用程序逻辑在这里
    doWork()

    // 保持应用程序运行以收集分析数据
    time.Sleep(60 * time.Second)
}

func doWork() {
    // 将被分析的 CPU 密集型工作
    for i := 0; i < 1000000; i++ {
        _ = computeSomething(i)
    }
}

func computeSomething(n int) int {
    return n * n * n
}
```

### 自定义服务器地址

```go
package main

import (
    "github.com/lazygophers/utils/pyroscope"
    "os"
)

func main() {
    // 使用自定义 Pyroscope 服务器地址
    serverURL := os.Getenv("PYROSCOPE_URL")
    if serverURL == "" {
        serverURL = "http://pyroscope.example.com:4040"
    }

    pyroscope.Load(serverURL)

    // 应用程序逻辑...
    runApplication()
}

func runApplication() {
    // 您的应用程序代码在这里
}
```

### Web 应用程序集成

```go
package main

import (
    "github.com/gin-gonic/gin"
    "github.com/lazygophers/utils/pyroscope"
    "net/http"
    "time"
)

func main() {
    // 首先初始化分析
    pyroscope.Load("http://localhost:4040")

    // 设置 Web 服务器
    r := gin.Default()

    r.GET("/api/heavy", heavyComputation)
    r.GET("/api/memory", memoryIntensive)
    r.GET("/api/concurrent", concurrentWork)

    // 启动服务器
    r.Run(":8080")
}

func heavyComputation(c *gin.Context) {
    start := time.Now()

    // CPU 密集型工作
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
    // 内存分配密集型工作
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

    // 生成 goroutine 进行并发工作
    for i := 0; i < 10; i++ {
        go func(id int) {
            time.Sleep(100 * time.Millisecond)
            for j := 0; j < 100000; j++ {
                _ = j * id
            }
            done <- true
        }(i)
    }

    // 等待所有 goroutine
    for i := 0; i < 10; i++ {
        <-done
    }

    c.JSON(http.StatusOK, gin.H{
        "goroutines_completed": 10,
    })
}
```

### 后台服务集成

```go
package main

import (
    "context"
    "github.com/lazygophers/utils/pyroscope"
    "sync"
    "time"
)

func main() {
    // 初始化分析
    pyroscope.Load("")

    ctx, cancel := context.WithTimeout(context.Background(), 5*time.Minute)
    defer cancel()

    var wg sync.WaitGroup

    // 启动后台工作器
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
            // 模拟随工作器变化的工作
            processData(id)
        }
    }
}

func processData(workerID int) {
    // 不同工作器的不同工作负载
    switch workerID % 3 {
    case 0:
        // CPU 密集型
        sum := 0
        for i := 0; i < 1000000; i++ {
            sum += i
        }
    case 1:
        // 内存密集型
        data := make([]int, 100000)
        for i := range data {
            data[i] = i * workerID
        }
    case 2:
        // 使用 goroutine 的 I/O 模拟
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

## 构建配置

### 开发构建（启用分析）

```bash
# 默认构建启用分析
go build -o myapp .

# 或明确启用分析
go build -tags="!release" -o myapp .

# 在发布构建中启用分析
go build -tags="release,pyroscope" -o myapp .
```

### 生产构建（禁用分析）

```bash
# 生产构建禁用分析
go build -tags="release" -o myapp .
```

### 构建标签组合

| 构建命令 | 分析状态 | 使用场景 |
|---|---|---|
| `go build` | 启用 | 开发 |
| `go build -tags="!release"` | 启用 | 开发（明确） |
| `go build -tags="release"` | 禁用 | 生产（默认） |
| `go build -tags="release,pyroscope"` | 启用 | 生产分析 |

## 环境配置

```go
package main

import (
    "github.com/lazygophers/utils/pyroscope"
    "os"
)

func initProfiling() {
    // 检查是否应启用分析
    if os.Getenv("ENABLE_PROFILING") != "true" {
        return
    }

    // 从环境获取服务器地址
    serverAddr := os.Getenv("PYROSCOPE_SERVER")
    if serverAddr == "" {
        serverAddr = "http://localhost:4040"
    }

    pyroscope.Load(serverAddr)
}

func main() {
    initProfiling()

    // 应用程序逻辑...
}
```

使用环境变量：

```bash
# 使用自定义服务器启用分析
ENABLE_PROFILING=true PYROSCOPE_SERVER=http://pyroscope-prod:4040 ./myapp

# 使用主机名标签运行
HOSTNAME=server-01 ./myapp
```

## 分析类型

该模块自动启用全面的分析：

### CPU 分析
- **目的**: 识别 CPU 热点和性能瓶颈
- **使用场景**: 优化计算密集型代码路径

### 内存分析
- **InuseObjects**: 当前分配的对象计数
- **AllocObjects**: 总分配的对象数（包括已释放的）
- **InuseSpace**: 当前分配的内存大小
- **AllocSpace**: 总分配的内存（包括已释放的）
- **使用场景**: 检测内存泄漏和优化分配

### Goroutine 分析
- **目的**: 跟踪 goroutine 创建和生命周期
- **使用场景**: 调试 goroutine 泄漏和并发问题

### 互斥锁分析
- **MutexCount**: 互斥锁争用事件计数
- **MutexDuration**: 等待互斥锁的时间
- **使用场景**: 优化同步并减少锁争用

### 阻塞分析
- **BlockCount**: 阻塞事件计数
- **BlockDuration**: 阻塞操作花费的时间
- **使用场景**: 识别阻塞 I/O 和同步问题

## 高级用法

### 自定义应用程序名称

```go
package main

import (
    "github.com/lazygophers/utils/app"
    "github.com/lazygophers/utils/pyroscope"
)

func init() {
    // 在分析之前设置自定义应用程序名称
    app.Name = "my-custom-service"
}

func main() {
    pyroscope.Load("")
    // 应用程序将以 "my-custom-service" 被分析
}
```

### 条件分析

```go
package main

import (
    "github.com/lazygophers/utils/pyroscope"
    "os"
)

func main() {
    // 仅在特定环境中启用分析
    environment := os.Getenv("ENVIRONMENT")

    switch environment {
    case "development", "staging":
        pyroscope.Load("")
    case "production":
        if os.Getenv("DEBUG_MODE") == "true" {
            pyroscope.Load("http://internal-pyroscope:4040")
        }
    }

    // 应用程序逻辑...
}
```

### Docker 集成

具有条件分析的 Dockerfile：

```dockerfile
FROM golang:1.21-alpine AS builder

WORKDIR /app
COPY . .

# 默认使用 release 标签构建
ARG BUILD_TAGS="release"
RUN go build -tags="${BUILD_TAGS}" -o myapp .

FROM alpine:latest

RUN apk --no-cache add ca-certificates
WORKDIR /root/

COPY --from=builder /app/myapp .

# 设置默认环境
ENV PYROSCOPE_SERVER=http://pyroscope:4040

CMD ["./myapp"]
```

构建命令：

```bash
# 生产构建（无分析）
docker build -t myapp:prod .

# 开发构建（带分析）
docker build --build-arg BUILD_TAGS="" -t myapp:dev .

# 启用分析的生产构建
docker build --build-arg BUILD_TAGS="release,pyroscope" -t myapp:prod-profile .
```

## 性能影响

### 开发构建
- **CPU 开销**: 根据应用程序特征，1-5%
- **内存开销**: 分析数据收集的少量增加
- **网络**: 定期上传到 Pyroscope 服务器

### 生产构建（Release 标签）
- **CPU 开销**: 0%（无操作实现）
- **内存开销**: 0%（无操作实现）
- **二进制大小**: 由于未使用代码消除，增加最小

### 建议
- 在生产中总是使用 release 构建，除非调试
- 在镜像生产的预发布环境中启用分析
- 临时使用启用分析的生产构建进行故障排除

## 故障排除

### 常见问题

#### 分析不工作
```go
// 检查分析是否实际启用
func main() {
    pyroscope.Load("")

    // 添加一些可识别的工作
    for i := 0; i < 1000000; i++ {
        _ = expensiveFunction(i)
    }
}
```

#### 连接问题
- 验证 Pyroscope 服务器正在运行：`curl http://localhost:4040`
- 检查从应用程序到服务器的网络连接
- 验证服务器地址配置

#### 缺少分析
- 确保应用程序运行足够长的时间来收集样本
- 检查应用程序名称配置
- 验证主机名环境变量

### 调试信息

```go
package main

import (
    "github.com/lazygophers/utils/pyroscope"
    "log"
    "os"
)

func main() {
    log.Printf("应用程序: %s", os.Getenv("APP_NAME"))
    log.Printf("主机名: %s", os.Getenv("HOSTNAME"))
    log.Printf("Pyroscope 服务器: %s", os.Getenv("PYROSCOPE_SERVER"))

    pyroscope.Load(os.Getenv("PYROSCOPE_SERVER"))

    // 您的应用程序...
}
```

## 最佳实践

### 1. 构建配置
- 默认在生产中使用 release 构建
- 仅在需要调试时启用分析
- 为您的团队记录分析构建程序

### 2. 服务器管理
- 在稳定环境中运行 Pyroscope 服务器
- 配置适当的保留策略
- 监控 Pyroscope 服务器资源使用

### 3. 应用程序集成
- 在应用程序启动早期初始化分析
- 使用有意义的应用程序名称
- 为多实例部署设置适当的主机名标签

### 4. 性能监控
- 在非生产环境中监控分析的性能影响
- 使用分析数据识别和修复性能瓶颈
- 定期审查和处理分析洞察

## 安全注意事项

- 分析数据可能包含有关应用程序内部的敏感信息
- 限制对 Pyroscope 服务器和分析的访问
- 考虑分析数据传输的网络安全
- 在生产环境中谨慎使用分析

## 相关工具和包

- [Grafana Pyroscope](https://pyroscope.io/): 持续分析平台
- [`runtime/pprof`](https://pkg.go.dev/runtime/pprof): Go 内置分析
- [`net/http/pprof`](https://pkg.go.dev/net/http/pprof): HTTP pprof 端点
- [`github.com/grafana/pyroscope-go`](https://github.com/grafana/pyroscope-go): 底层 Pyroscope Go 客户端

## 示例配置

### Docker Compose 设置
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

### Kubernetes 部署
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