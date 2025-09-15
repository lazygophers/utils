# network - Network Utilities

[![Go Reference](https://pkg.go.dev/badge/github.com/lazygophers/utils/network.svg)](https://pkg.go.dev/github.com/lazygophers/utils/network)

A comprehensive Go package for network operations including IP address utilities, real IP extraction from HTTP headers, network interface management, and HTTP client utilities with advanced proxy and load balancing support.

## Features

### IP Address Utilities
- **Local IP Detection**: Identify private, loopback, and link-local addresses
- **Network Interface Management**: Extract IP addresses from network interfaces
- **IPv4/IPv6 Support**: Comprehensive support for both IP versions
- **Automatic Interface Selection**: Smart selection of eth0, en0, and other common interfaces

### Real IP Extraction
- **Comprehensive Header Support**: Extract real client IP from various proxy headers
- **CDN Support**: Built-in support for CloudFlare, Fastly, and other CDNs
- **Proxy Chain Handling**: Process X-Forwarded-For chains correctly
- **Fallback Mechanisms**: Multiple fallback strategies for reliable IP detection

### HTTP Client Utilities
- **Connection Pooling**: Efficient HTTP connection management
- **Timeout Configuration**: Flexible timeout settings for different scenarios
- **Custom Transport**: Enhanced HTTP transport with better defaults
- **Error Handling**: Robust error handling and retry mechanisms

## Installation

```bash
go get github.com/lazygophers/utils/network
```

## Quick Start

```go
package main

import (
    "fmt"
    "net/http"
    "github.com/lazygophers/utils/network"
)

func main() {
    // Check if IP is local/private
    fmt.Println(network.IsLocalIp("192.168.1.1"))  // true
    fmt.Println(network.IsLocalIp("8.8.8.8"))      // false

    // Get local listening IP
    localIP := network.GetListenIp()
    fmt.Printf("Local IP: %s\n", localIP)

    // Extract real IP from HTTP request
    http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
        realIP := network.RealIpFromHeader(r.Header)
        fmt.Printf("Real client IP: %s\n", realIP)
    })
}
```

## Core API Reference

### IP Address Utilities

```go
// Check if IP address is local/private
isLocal := network.IsLocalIp("192.168.1.100")    // true (private)
isLocal = network.IsLocalIp("127.0.0.1")         // true (loopback)
isLocal = network.IsLocalIp("169.254.1.1")       // true (link-local)
isLocal = network.IsLocalIp("8.8.8.8")           // false (public)

// IPv6 support
isLocal = network.IsLocalIp("::1")                // true (IPv6 loopback)
isLocal = network.IsLocalIp("fe80::1")            // true (IPv6 link-local)
isLocal = network.IsLocalIp("2001:db8::1")        // false (IPv6 global)
```

### Network Interface Management

```go
// Get IP address from specific interface
ip := network.GetInterfaceIpByName("eth0", false)    // IPv4 from eth0
ip = network.GetInterfaceIpByName("eth0", true)      // IPv6 from eth0

// Get best local IP for listening
localIP := network.GetListenIp()         // IPv4 preferred
localIP = network.GetListenIp(true)      // IPv6 preferred

// Example outputs:
// "192.168.1.100" (typical home network)
// "10.0.0.50" (corporate network)
// "172.16.0.10" (Docker/container network)
```

### Interface Selection Logic

The package intelligently selects network interfaces:

1. **eth0** - Common on Linux servers
2. **en0** - Common on macOS
3. **All interfaces** - Fallback to system interface enumeration

```go
// Manual interface IP extraction
interfaces, _ := net.InterfaceAddrs()
ip := network.GetInterfaceIpByAddrs(interfaces, false)  // IPv4
ip = network.GetInterfaceIpByAddrs(interfaces, true)    // IPv6
```

### Real IP Extraction

```go
// Extract real client IP from HTTP headers
func handler(w http.ResponseWriter, r *http.Request) {
    realIP := network.RealIpFromHeader(r.Header)

    if realIP != "" {
        fmt.Printf("Client IP: %s\n", realIP)
    } else {
        // Fallback to remote address
        fmt.Printf("Remote IP: %s\n", r.RemoteAddr)
    }
}
```

### Supported Headers

The package checks headers in priority order:

1. **CloudFlare Headers**:
   - `Cf-Connecting-Ip`
   - `Cf-Pseudo-Ipv4`
   - `Cf-Connecting-Ipv6`
   - `Cf-Pseudo-Ipv6`

2. **CDN Headers**:
   - `Fastly-Client-Ip`
   - `True-Client-Ip`

3. **Standard Proxy Headers**:
   - `X-Real-IP`
   - `X-Client-IP`
   - `X-Original-Forwarded-For`
   - `X-Forwarded-For`
   - `X-Forwarded`
   - `Forwarded-For`
   - `Forwarded`

```go
// Example with multiple proxy headers
headers := http.Header{
    "X-Forwarded-For": []string{"203.0.113.1, 198.51.100.1, 192.168.1.1"},
    "X-Real-IP":       []string{"203.0.113.1"},
}

realIP := network.RealIpFromHeader(headers)
// Returns: "203.0.113.1" (first non-local IP)
```

## Advanced Features

### IPv6 Support

```go
// Prefer IPv6 addresses
ipv6 := network.GetListenIp(true)
fmt.Println(ipv6)  // Output: "2001:db8::1" or similar

// Get IPv6 from specific interface
ipv6 = network.GetInterfaceIpByName("eth0", true)

// Check IPv6 local addresses
isLocal := network.IsLocalIp("::1")         // true (loopback)
isLocal = network.IsLocalIp("fe80::1")      // true (link-local)
isLocal = network.IsLocalIp("fd00::1")      // true (unique local)
```

### Header Chain Processing

```go
// Process X-Forwarded-For chains
headers := http.Header{
    "X-Forwarded-For": []string{
        "8.8.8.8, 192.168.1.1, 10.0.0.1",  // Public, private, private
    },
}

realIP := network.RealIpFromHeader(headers)
// Returns: "8.8.8.8" (first non-local IP in chain)
```

### Error Handling and Fallbacks

```go
// The package handles various error conditions gracefully:

// Non-existent interface
ip := network.GetInterfaceIpByName("nonexistent", false)
// Returns: "" (empty string)

// Invalid IP format in headers
headers := http.Header{
    "X-Real-IP": []string{"invalid-ip"},
}
realIP := network.RealIpFromHeader(headers)
// Returns: "" (empty string, invalid IPs are skipped)

// No valid interfaces
ip := network.GetListenIp()
// Returns: "" if no valid interfaces found (logs error)
```

## Best Practices

### 1. IP Detection in Web Applications

```go
func getRealClientIP(r *http.Request) string {
    // First try header extraction
    if realIP := network.RealIpFromHeader(r.Header); realIP != "" {
        return realIP
    }

    // Fallback to remote address
    if host, _, err := net.SplitHostPort(r.RemoteAddr); err == nil {
        return host
    }

    return r.RemoteAddr
}
```

### 2. Server Binding

```go
func startServer() {
    // Get appropriate local IP for binding
    bindIP := network.GetListenIp()
    if bindIP == "" {
        bindIP = "0.0.0.0"  // Fallback to all interfaces
    }

    addr := fmt.Sprintf("%s:8080", bindIP)
    log.Printf("Starting server on %s", addr)

    http.ListenAndServe(addr, nil)
}
```

### 3. Load Balancer Integration

```go
func setupLoadBalancerHandler() http.HandlerFunc {
    return func(w http.ResponseWriter, r *http.Request) {
        // Extract real client IP for logging/rate limiting
        clientIP := network.RealIpFromHeader(r.Header)

        // Log with real IP
        log.Printf("Request from %s: %s %s", clientIP, r.Method, r.URL.Path)

        // Add to rate limiter with real IP
        if !rateLimiter.Allow(clientIP) {
            http.Error(w, "Rate limit exceeded", http.StatusTooManyRequests)
            return
        }

        // Process request...
    }
}
```

### 4. Docker/Container Environments

```go
func getContainerIP() string {
    // In containers, eth0 is usually the main interface
    if ip := network.GetInterfaceIpByName("eth0", false); ip != "" {
        return ip
    }

    // Fallback to automatic detection
    return network.GetListenIp()
}
```

## Security Considerations

### Header Spoofing Protection

```go
// Always validate that extracted IPs are from trusted sources
func validateTrustedProxy(r *http.Request) bool {
    // Get the immediate upstream IP
    host, _, _ := net.SplitHostPort(r.RemoteAddr)

    // Check if it's a trusted proxy/load balancer
    trustedProxies := []string{
        "192.168.1.100",  // Internal load balancer
        "10.0.0.1",       // Trusted proxy
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
        // Trust headers only from known proxies
        clientIP = network.RealIpFromHeader(r.Header)
    }

    if clientIP == "" {
        // Fallback to direct connection
        host, _, _ := net.SplitHostPort(r.RemoteAddr)
        clientIP = host
    }

    // Use clientIP for security decisions...
}
```

### Rate Limiting by Real IP

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
        limiter = rate.NewLimiter(rate.Every(time.Minute), 60)  // 60 requests per minute
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

    // Handle request...
}
```

## Use Cases

### 1. Microservice Discovery

```go
func registerService() {
    serviceIP := network.GetListenIp()
    servicePort := 8080

    // Register with service discovery
    registry.Register(serviceIP, servicePort)
}
```

### 2. Geo-location Services

```go
func geolocateRequest(r *http.Request) *GeoLocation {
    clientIP := network.RealIpFromHeader(r.Header)
    if clientIP == "" || network.IsLocalIp(clientIP) {
        return nil  // Can't geolocate local IPs
    }

    return geoService.Lookup(clientIP)
}
```

### 3. Analytics and Logging

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

    logger.Info("HTTP request", logEntry)
}
```

### 4. Security Auditing

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

## Error Handling

The package uses Go's standard error handling patterns:

```go
// Functions return empty strings on error, not nil
ip := network.GetInterfaceIpByName("invalid", false)
if ip == "" {
    log.Println("Failed to get IP from interface")
}

// Check logs for detailed error information
// The package logs errors using the lazygophers/log package
```

## Performance Considerations

- **Interface enumeration**: Cached for efficiency
- **Header parsing**: Optimized for common cases
- **IP validation**: Uses efficient netip.ParseAddr
- **Memory allocation**: Minimal allocations in hot paths

## Integration Examples

### Gin Web Framework

```go
import "github.com/gin-gonic/gin"

func setupGin() *gin.Engine {
    r := gin.New()

    // Middleware to extract real IP
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

### Echo Framework

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

## Related Packages

- `net` - Go standard library networking
- `net/netip` - Modern IP address types (dependency)
- `github.com/lazygophers/log` - Logging utilities (dependency)
- `net/http` - HTTP client and server functionality

## Contributing

This package is part of the LazyGophers Utils collection. For contributions:

1. Follow Go networking best practices
2. Add tests for new IP validation logic
3. Ensure compatibility with major proxy/CDN services
4. Document security implications of changes

## License

This package is part of the LazyGophers Utils project. See the main repository for license information.

---

*For production deployments behind load balancers or CDNs, always validate that IP extraction headers are from trusted sources to prevent header spoofing attacks.*