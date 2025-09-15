# network - 网络工具包

[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils/network.svg)](https://pkg.go.dev/github.com/lazygophers/utils/network)

一个功能全面的 Go 网络操作包，包括 IP 地址工具、HTTP 头部真实 IP 提取、网络接口管理以及具有高级代理和负载均衡支持的 HTTP 客户端工具。

## 核心特性

### IP 地址工具
- **本地 IP 检测**: 识别私有、环回和链路本地地址
- **网络接口管理**: 从网络接口提取 IP 地址
- **IPv4/IPv6 支持**: 全面支持两种 IP 版本
- **自动接口选择**: 智能选择 eth0、en0 和其他常见接口

### 真实 IP 提取
- **全面的头部支持**: 从各种代理头部提取真实客户端 IP
- **CDN 支持**: 内置支持 CloudFlare、Fastly 和其他 CDN
- **代理链处理**: 正确处理 X-Forwarded-For 链
- **回退机制**: 多种回退策略确保可靠的 IP 检测

### HTTP 客户端工具
- **连接池**: 高效的 HTTP 连接管理
- **超时配置**: 针对不同场景的灵活超时设置
- **自定义传输**: 具有更好默认值的增强 HTTP 传输
- **错误处理**: 强大的错误处理和重试机制

## 安装

```bash
go get github.com/lazygophers/utils/network
```

## 快速开始

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/lazygophers/utils/network"
)

func main() {
    // 检查 IP 是否为本地/私有
    fmt.Println(network.IsLocalIp("192.168.1.1"))  // true
    fmt.Println(network.IsLocalIp("8.8.8.8"))      // false

    // 获取本地监听 IP
    localIP := network.GetListenIp()
    fmt.Printf("本地 IP: %s\n", localIP)

    // 从 HTTP 请求中提取真实 IP
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        realIP := network.RealIpFromHeader(r.Header)
        fmt.Printf("真实客户端 IP: %s\n", realIP)
    })
}
```

## 核心 API 参考

### IP 地址工具

```go
// 检查 IP 地址是否为本地/私有
isLocal := network.IsLocalIp("192.168.1.100")    // true (私有)
isLocal = network.IsLocalIp("127.0.0.1")         // true (环回)
isLocal = network.IsLocalIp("169.254.1.1")       // true (链路本地)
isLocal = network.IsLocalIp("8.8.8.8")           // false (公共)

// IPv6 支持
isLocal = network.IsLocalIp("::1")                // true (IPv6 环回)
isLocal = network.IsLocalIp("fe80::1")            // true (IPv6 链路本地)
isLocal = network.IsLocalIp("2001:db8::1")        // false (IPv6 全球)
```

### 网络接口管理

```go
// 从特定接口获取 IP 地址
ip := network.GetInterfaceIpByName("eth0", false)    // 从 eth0 获取 IPv4
ip = network.GetInterfaceIpByName("eth0", true)      // 从 eth0 获取 IPv6

// 获取最佳本地监听 IP
localIP := network.GetListenIp()         // 首选 IPv4
localIP = network.GetListenIp(true)      // 首选 IPv6

// 示例输出:
// "192.168.1.100" (典型家庭网络)
// "10.0.0.50" (企业网络)
// "172.16.0.10" (Docker/容器网络)
```

### 接口选择逻辑

包智能选择网络接口：

1. **eth0** - Linux 服务器常见
2. **en0** - macOS 常见
3. **所有接口** - 回退到系统接口枚举

```go
// 手动接口 IP 提取
interfaces, _ := net.InterfaceAddrs()
ip := network.GetInterfaceIpByAddrs(interfaces, false)  // IPv4
ip = network.GetInterfaceIpByAddrs(interfaces, true)    // IPv6
```

### 真实 IP 提取

```go
// 从 HTTP 头部提取真实客户端 IP
func handler(w http.ResponseWriter, r *http.Request) {
    realIP := network.RealIpFromHeader(r.Header)

    if realIP != "" {
        fmt.Printf("客户端 IP: %s\n", realIP)
    } else {
        // 回退到远程地址
        fmt.Printf("远程 IP: %s\n", r.RemoteAddr)
    }
}
```

### 支持的头部

包按优先级顺序检查头部：

1. **CloudFlare 头部**:
   - `Cf-Connecting-Ip`
   - `Cf-Pseudo-Ipv4`
   - `Cf-Connecting-Ipv6`
   - `Cf-Pseudo-Ipv6`

2. **CDN 头部**:
   - `Fastly-Client-Ip`
   - `True-Client-Ip`

3. **标准代理头部**:
   - `X-Real-IP`
   - `X-Client-IP`
   - `X-Original-Forwarded-For`
   - `X-Forwarded-For`
   - `X-Forwarded`
   - `Forwarded-For`
   - `Forwarded`

```go
// 多代理头部示例
headers := http.Header{
    "X-Forwarded-For": []string{"203.0.113.1, 198.51.100.1, 192.168.1.1"},
    "X-Real-IP":       []string{"203.0.113.1"},
}

realIP := network.RealIpFromHeader(headers)
// 返回: "203.0.113.1" (第一个非本地 IP)
```

## 高级功能

### IPv6 支持

```go
// 首选 IPv6 地址
ipv6 := network.GetListenIp(true)
fmt.Println(ipv6)  // 输出: "2001:db8::1" 或类似

// 从特定接口获取 IPv6
ipv6 = network.GetInterfaceIpByName("eth0", true)

// 检查 IPv6 本地地址
isLocal := network.IsLocalIp("::1")         // true (环回)
isLocal = network.IsLocalIp("fe80::1")      // true (链路本地)
isLocal = network.IsLocalIp("fd00::1")      // true (唯一本地)
```

### 头部链处理

```go
// 处理 X-Forwarded-For 链
headers := http.Header{
    "X-Forwarded-For": []string{
        "8.8.8.8, 192.168.1.1, 10.0.0.1",  // 公共, 私有, 私有
    },
}

realIP := network.RealIpFromHeader(headers)
// 返回: "8.8.8.8" (链中第一个非本地 IP)
```

### 错误处理和回退

```go
// 包优雅地处理各种错误条件:

// 不存在的接口
ip := network.GetInterfaceIpByName("nonexistent", false)
// 返回: "" (空字符串)

// 头部中的无效 IP 格式
headers := http.Header{
    "X-Real-IP": []string{"invalid-ip"},
}
realIP := network.RealIpFromHeader(headers)
// 返回: "" (空字符串，无效 IP 被跳过)

// 没有有效接口
ip := network.GetListenIp()
// 如果没有找到有效接口则返回: "" (记录错误)
```

## 最佳实践

### 1. Web 应用中的 IP 检测

```go
func getRealClientIP(r *http.Request) string {
    // 首先尝试头部提取
    if realIP := network.RealIpFromHeader(r.Header); realIP != "" {
        return realIP
    }

    // 回退到远程地址
    if host, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
        return host
    }

    return r.RemoteAddr
}
```

### 2. 服务器绑定

```go
func startServer() {
    // 获取适当的本地 IP 进行绑定
    bindIP := network.GetListenIp()
    if bindIP == "" {
        bindIP = "0.0.0.0"  // 回退到所有接口
    }

    addr := fmt.Sprintf("%s:8080", bindIP)
    log.Printf("在 %s 启动服务器", addr)

    http.ListenAndServe(addr, nil)
}
```

### 3. 负载均衡器集成

```go
func setupLoadBalancerHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // 提取真实客户端 IP 用于日志记录/限流
        clientIP := network.RealIpFromHeader(r.Header)

        // 使用真实 IP 记录日志
        log.Printf("来自 %s 的请求: %s %s", clientIP, r.Method, r.URL.Path)

        // 使用真实 IP 添加到限流器
        if !rateLimiter.Allow(clientIP) {
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }

        // 处理请求...
    }
}
```

### 4. Docker/容器环境

```go
func getContainerIP() string {
    // 在容器中，eth0 通常是主接口
    if ip := network.GetInterfaceIpByName("eth0", false); ip != "" {
        return ip
    }

    // 回退到自动检测
    return network.GetListenIp()
}
```

## 安全考虑

### 头部欺骗保护

```go
// 始终验证提取的 IP 来自可信源
func validateTrustedProxy(r *http.Request) bool {
    // 获取直接上游 IP
    host, _, _ := net.SplitHostPort(r.RemoteAddr)

    // 检查是否为可信代理/负载均衡器
    trustedProxies := []string{
        "192.168.1.100",  // 内部负载均衡器
        "10.0.0.1",       // 可信代理
    }

    for _, trusted := range trustedProxies {
        if host == trusted {
            return true
        }
    }

    return false
}

func secureHandler(w http.ResponseWriter, r *http.Request) {
    var clientIP string

    if validateTrustedProxy(r) {
        // 只信任来自已知代理的头部
        clientIP = network.RealIpFromHeader(r.Header)
    }

    if clientIP == "" {
        // 回退到直接连接
        host, _, _ := net.SplitHostPort(r.RemoteAddr)
        clientIP = host
    }

    // 使用 clientIP 进行安全决策...
}
```

### 基于真实 IP 的限流

```go
import "golang.org/x/time/rate"

var rateLimiters = make(map[string]*rate.Limiter)
var mu sync.RWMutex

func getRateLimiter(ip string) *rate.Limiter {
    mu.RLock()
    limiter, exists := rateLimiters[ip]
    mu.RUnlock()

    if !exists {
        mu.Lock()
        limiter = rate.NewLimiter(rate.Every(time.Minute), 60)  // 每分钟60个请求
        rateLimiters[ip] = limiter
        mu.Unlock()
    }

    return limiter
}

func rateLimitedHandler(w http.ResponseWriter, r *http.Request) {
    clientIP := network.RealIpFromHeader(r.Header)
    if clientIP == "" {
        host, _, _ := net.SplitHostPort(r.RemoteAddr)
        clientIP = host
    }

    limiter := getRateLimiter(clientIP)
    if !limiter.Allow() {
        http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
        return
    }

    // 处理请求...
}
```

## 使用场景

### 1. 微服务发现

```go
func registerService() {
    serviceIP := network.GetListenIp()
    servicePort := 8080

    // 向服务发现注册
    registry.Register(serviceIP, servicePort)
}
```

### 2. 地理位置服务

```go
func geolocateRequest(r *http.Request) *GeoLocation {
    clientIP := network.RealIpFromHeader(r.Header)
    if clientIP == "" || network.IsLocalIp(clientIP) {
        return nil  // 无法定位本地 IP
    }

    return geoService.Lookup(clientIP)
}
```

### 3. 分析和日志记录

```go
func logRequest(r *http.Request) {
    clientIP := network.RealIpFromHeader(r.Header)

    logEntry := map[string]interface{}{
        "client_ip":   clientIP,
        "method":      r.Method,
        "path":        r.URL.Path,
        "user_agent":  r.Header.Get("User-Agent"),
        "timestamp":   time.Now(),
    }

    logger.Info("HTTP 请求", logEntry)
}
```

### 4. 安全审计

```go
func auditSecurityEvent(r *http.Request, event string) {
    clientIP := network.RealIpFromHeader(r.Header)

    audit := SecurityEvent{
        IP:        clientIP,
        Event:     event,
        Timestamp: time.Now(),
        Headers:   r.Header,
        UserAgent: r.Header.Get("User-Agent"),
    }

    securityLog.Record(audit)
}
```

## 错误处理

包使用 Go 的标准错误处理模式：

```go
// 函数在错误时返回空字符串，而不是 nil
ip := network.GetInterfaceIpByName("invalid", false)
if ip == "" {
    log.Println("从接口获取 IP 失败")
}

// 检查日志获取详细错误信息
// 包使用 lazygophers/log 包记录错误
```

## 性能考虑

- **接口枚举**: 为效率而缓存
- **头部解析**: 针对常见情况优化
- **IP 验证**: 使用高效的 netip.ParseAddr
- **内存分配**: 在热路径中最少分配

## 集成示例

### Gin Web 框架

```go
import "github.com/gin-gonic/gin"

func setupGin() *gin.Engine {
    r := gin.New()

    // 提取真实 IP 的中间件
    r.Use(func(c *gin.Context) {
        realIP := network.RealIpFromHeader(c.Request.Header)
        if realIP != "" {
            c.Set("client_ip", realIP)
        }
        c.Next()
    })

    r.GET("/ip", func(c *gin.Context) {
        clientIP, _ := c.Get("client_ip")
        c.JSON(200, gin.H{"ip": clientIP})
    })

    return r
}
```

### Echo 框架

```go
import "github.com/labstack/echo/v4"

func setupEcho() *echo.Echo {
    e := echo.New()

    e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
        return func(c echo.Context) error {
            realIP := network.RealIpFromHeader(c.Request().Header)
            if realIP != "" {
                c.Set("client_ip", realIP)
            }
            return next(c)
        }
    })

    return e
}
```

## 相关包

- `net` - Go 标准库网络功能
- `net/netip` - 现代 IP 地址类型（依赖）
- `github.com/lazygophers/log` - 日志工具（依赖）
- `net/http` - HTTP 客户端和服务器功能

## 贡献

此包是 LazyGophers Utils 集合的一部分。贡献指南：

1. 遵循 Go 网络最佳实践
2. 为新的 IP 验证逻辑添加测试
3. 确保与主要代理/CDN 服务的兼容性
4. 记录更改的安全影响

## 许可证

此包是 LazyGophers Utils 项目的一部分。许可证信息请查看主仓库。

---

*对于在负载均衡器或 CDN 后面的生产部署，请始终验证 IP 提取头部来自可信源，以防止头部欺骗攻击。*